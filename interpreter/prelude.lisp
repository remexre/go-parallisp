;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;; MACROS ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defmacro != [a b] `(not (= ,a ,b)))

(defmacro ->> [start &rest forms]
	(defun helper [start forms]
		(if forms
			(let* ((form (car forms))
						(type (type-of form)))
				(switch type
					'symbol	(helper (list form start) (cdr forms))
					'cons		(helper (append form (list start)) (cdr forms))
									(error "Invalid type for ->>: " type ": " form)))
			start))
	(helper start forms))

(defmacro for [var start pred next &rest code]
	(let ((sym (gensym)))
		`(progn
			(defun ,sym ,[var]
				(if ,pred
					(progn
						,@code
						(,sym ,next))
					nil))
			(,sym ,start))))

(defmacro if [pred then else]
	(list 'cond pred then else))

(defmacro if-log [pred then else]
	`(if ,pred
		(progn (println ,then) 't)
		(progn (println ,else) nil)))

(defmacro let* [defs &rest code]
	(defun helper [defs]
		(if (= (len defs) 0) (cons 'progn code)
			`(let (,(car defs)) ,(helper (cdr defs)))))
	(helper defs))

(defmacro quasiquote [expr]
	(defun helper [expr]
		(if (= (type-of expr) 'cons)
			(switch (car expr)
				'unquote				(cons 'list (cdr expr))
				'unquote-splice	(car (cdr expr))
												(list 'list (cons 'append (mapcar helper expr))))
			(list 'list (list 'quote expr))))
	(switch (type-of expr)
		'cons		(switch (car expr)
							'unquote				(car (cdr expr))
							'unquote-splice	(error "cannot splice into root of quasiquote")
															(cons 'append (mapcar helper expr)))
		'vector	(list 'lst->vec (list 'apply 'append (list 'vec->lst (mapvec helper expr))))
						(list 'quote expr)))

(defmacro switch [expr &rest cases]
	(defun helper [expr cases out]
		(cond
			(nil? cases)				out
			(nil? (cdr cases))	(cons (car cases) out)
													(helper expr
														(cdr (cdr cases))
														(cons
															(car (cdr cases))
															(cons (list '= expr (car cases)) out)))))
	(let ((sym (gensym)))
		(list 'let (list (list sym expr))
			(cons 'cond (reverse (helper sym cases nil))))))

(defmacro test-suite [parser &rest args]
	(defun helper [args out]
		(if (nil? args) out
			(let ((expected  (car args))
						(input     (car (cdr args)))
						(next-args (cdr (cdr args))))
				(helper next-args (cons [
					input
					(eval (list parser input))
					expected
				] out)))))
	(cond
		(= (len args) 0)			't
		(!= (% (len args) 2)	0)
			(error "test-suite: needs odd number of arguments")
		`(if (not (run-tests ',parser ',(helper (reverse args) nil)))
			(error "Test suite failed")
			't)))

(defmacro unless [condition &rest code]
	(list 'if condition 'nil (cons 'progn code)))

(defmacro when [condition &rest code]
	(list 'if condition (cons 'progn code) 'nil))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;; FUNCTIONS ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defun and [&rest conds]
	(defun helper [conds]
		(if (nil? conds)
			't
			(if (car conds) (helper (cdr conds)) nil)))
	(helper conds))

(defun append [a &rest b]
	(defun helper [a b]
		(cons (car a)
			(if (nil? (cdr a))
				b
				(helper (cdr a) b))))
	(if (nil? b) a
		(apply append (cons (helper a (car b)) (cdr b)))))

(defun contains? [item lst]
	(->> lst
		(filter (lambda [x] (= x item)))
		len
		(!= 0)))

(defun color [str &rest colors]
	(defun helper [color]
		(string (switch color
			'none				0				'bold				1				'underline	4
			'black			30			'red				31			'green			32			'yellow			33
			'blue				34			'magenta		35			'cyan				36			'white			37
			'bg-black		40			'bg-red			41			'bg-green		42			'bg-yellow	43
			'bg-blue		44			'bg-magenta	45			'bg-cyan		46			'bg-white		47
			(error "Unknown color: " color))))
	(format nil (+ "\x1b[" (join (mapcar helper colors) ";") "m") str "\x1b[0m"))

(defun error [&rest exprs]
	(**error** (join (mapcar string-bare exprs) "")))

(defun filter [pred lst]
	(defun helper [in out]
		(if in
			(let ((x (car in)))
				(let ((next (if (pred x) (cons x out) out)))
					(helper (cdr in) next)))
			out))
	(reverse (helper lst nil)))

(defun format [format &rest exprs]
	(defun nil-format [exprs]
		(join (mapcar string-bare exprs) ""))
	(if format
		(error "NYI: (format " format " " (nil-format exprs) ")")
		(nil-format exprs)))

(defun join [strs sep]
	(defun helper [strs out]
		(cond
			(nil? strs)				out
			(nil? (cdr strs))	(cons (car strs) out)
												(helper (cdr strs) (cons sep (cons (car strs) out)))))
	(if strs
		(apply + (reverse (helper strs nil)))
		""))

(defun lst->vec [lst] (apply vector lst))

(defun mapcar [fn &rest lists]
	(defun mapcar-one [fn lst]
		(defun helper [in out]
			(if in
				(helper (cdr in) (cons (fn (car in)) out))
				(reverse out)))
			(helper lst nil))
	(defun helper [lists out]
		(if (= (len (filter nil? lists)) (len lists))
			out
			(helper
				(mapcar-one cdr lists)
				(cons (apply fn (mapcar-one car lists)) out))))
	(reverse (helper lists nil)))

(defun nil? [expr] (= expr nil))

(defun not [expr] (if expr nil 't))

(defun or [&rest conds]
	(defun helper [conds]
		(if (nil? conds)
			nil
			(if (car conds) 't (helper (cdr conds)))))
	(helper conds))

(defun print [&rest exprs]
	(**print** (join (mapcar string-bare exprs) "")))

(defun println [&rest exprs]
	(apply print exprs)
	(**print** "\n"))

(defun range [&rest args]
	(defun helper [start stop step]
		(defun helper [i out]
			(if (< i stop)
				(helper (+ i step) (cons i out))
				out))
		(reverse (helper start nil)))
	(switch (len args)
		1	(helper 0          (car args)  1)
		2	(helper (car args) (cadr args) 1)
		3	(apply helper args)
			(error "range: incorrect usage")))

(defun reverse [lst]
	(defun helper [out in]
		(if (nil? in) out
			(helper (cons (car in) out) (cdr in))))
	(helper nil lst))

(defun string-bare [expr]
	(if (= (type-of expr) 'string) expr (string expr)))

(defun run-test [input got expected]
	(let ((check-mark (color "\u2713"      'bold 'green))
				(x-mark     (color "\u2717"      'bold 'red))
				(arrow-blue (color "===>"        'bold 'blue))
				(arrow-ok   (color "===>"        'bold 'green))
				(arrow-fail (color "=/=>"        'bold 'red))
				(msg-fail   (color "TEST FAILED" 'bold 'red))
				(inp-str    (string input))
				(exp-str    (string expected))
				(got-str    (string got)))
		(if-log (= got expected)
			(join (list inp-str arrow-ok exp-str check-mark) " ")
			(join (list
				(join (list inp-str arrow-fail exp-str x-mark) " ")
				(+ "instead, got " (string got))
				msg-fail) "\n"))))

(defun run-tests [parser tests]
	(defun header [text]
		(color
			(if (< (len text) 70)
				(let* ((n (- (/ (- 80 (len text)) 2) 1))
							(bar (* "=" n)))
					(join (list bar text bar) " "))
				text) 'bold))
	(defun helper [tests]
		(cond
			(nil? tests)												't
			(nil? (let ((row (car tests)))
							(let ((inp (@ row 0))
										(got (@ row 1))
										(exp (@ row 2)))
								(run-test inp got exp))))	nil
																					(helper (cdr tests))))
	(println (header (string parser)))
	(helper tests))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;; CADR FUNCTIONS ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defun caar [x] (car (car x)))
(defun cdar [x] (cdr (car x)))
(defun cadr [x] (car (cdr x)))
(defun cddr [x] (cdr (cdr x)))

(defun caaar [x] (car (car (car x))))
(defun cdaar [x] (cdr (car (car x))))
(defun cadar [x] (car (cdr (car x))))
(defun cddar [x] (cdr (cdr (car x))))
(defun caadr [x] (car (car (cdr x))))
(defun cdadr [x] (cdr (car (cdr x))))
(defun caddr [x] (car (cdr (cdr x))))
(defun cdddr [x] (cdr (cdr (cdr x))))

(defun caaaar [x] (car (car (car (car x)))))
(defun cdaaar [x] (cdr (car (car (car x)))))
(defun cadaar [x] (car (cdr (car (car x)))))
(defun cddaar [x] (cdr (cdr (car (car x)))))
(defun caadar [x] (car (car (cdr (car x)))))
(defun cdadar [x] (cdr (car (cdr (car x)))))
(defun caddar [x] (car (cdr (cdr (car x)))))
(defun cdddar [x] (cdr (cdr (cdr (car x)))))
(defun caaadr [x] (car (car (car (cdr x)))))
(defun cdaadr [x] (cdr (car (car (cdr x)))))
(defun cadadr [x] (car (cdr (car (cdr x)))))
(defun cddadr [x] (cdr (cdr (car (cdr x)))))
(defun caaddr [x] (car (car (cdr (cdr x)))))
(defun cdaddr [x] (cdr (car (cdr (cdr x)))))
(defun cadddr [x] (car (cdr (cdr (cdr x)))))
(defun cddddr [x] (cdr (cdr (cdr (cdr x)))))
