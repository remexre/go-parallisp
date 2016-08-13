package types

import (
	"bytes"
	"fmt"
	"io"
)

// Cons is a pair of pointers, possibly to nil.
type Cons [2]Expr

// NewCons is a helper to avoid https://github.com/golang/go/issues/9171.
func NewCons(car, cdr Expr) Cons {
	return Cons{car, cdr}
}

// NewConsList creates a list of cons cells from Exprs.
func NewConsList(exprs ...Expr) Expr {
	var out Expr
	for i := len(exprs) - 1; i >= 0; i-- {
		out = Cons{exprs[i], out}
	}
	return out
}

// NewImproperConsList creates an improper list of cons cells from Exprs.
func NewImproperConsList(exprs ...Expr) Expr {
	out := exprs[len(exprs)-1]
	for i := len(exprs) - 2; i >= 0; i-- {
		out = Cons{exprs[i], out}
	}
	return out
}

// Car returns the car of the cons cell it is called on.
func (expr Cons) Car() Expr { return expr[0] }

// Cdr returns the cdr of the cons cell it is called on.
func (expr Cons) Cdr() Expr { return expr[1] }

// Eval evaluates an expression.
func (expr Cons) Eval(env Env) (Expr, error) {
	if !expr.IsList() {
		return nil, fmt.Errorf("parallisp.types: cannot evaluate non-list %s", expr)
	}

	fnExpr, err := EvalExpr(env, expr.Car())
	if err != nil {
		return nil, err
	}
	if fn, ok := fnExpr.(Function); ok {
		return fn.Call(env, expr.ToSlice()[1:]...)
	}

	return nil, fmt.Errorf("parallisp.types: cannot call non-function %s in %s", fnExpr, expr)
}

// Get returns the nth element of the string.
func (expr Cons) Get(n Integer) (Expr, error) {
	if !expr.IsList() {
		return nil, fmt.Errorf("not a list: %s", expr)
	}
	list := expr.ToSlice()
	if int(n) >= len(list) || n < 0 {
		return nil, fmt.Errorf("%s does not have an index %d", expr, n)
	}
	return list[n], nil
}

// IsList returns true if the expression is a proper cons-list.
func (expr Cons) IsList() bool {
	if expr.Cdr() == nil {
		return true
	} else if next, ok := expr.Cdr().(Cons); ok {
		return next.IsList()
	}
	return false
}

// Len returns the length of the string.
func (expr Cons) Len() (Integer, error) {
	if !expr.IsList() {
		return 0, fmt.Errorf("not a list: %s", expr)
	}
	return Integer(len(expr.ToSlice())), nil
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (expr Cons) LiteralAsm() string {
	return fmt.Sprintf(".quad $1f+%d, $2f+%d\n1:%s\n2:%s",
		ExprToTypeAsm(expr[0]),
		ExprToTypeAsm(expr[1]),
		ExprToLiteralAsm(expr[0]),
		ExprToLiteralAsm(expr[1]))
}

// Slice slices the list from the from index to the to index.
func (expr Cons) Slice(from, to Integer) (Expr, error) {
	if !expr.IsList() {
		return nil, fmt.Errorf("not a list: %s", expr)
	}
	list := expr.ToSlice()
	if int(to) > len(list) || to < 0 {
		return nil, fmt.Errorf("%s does not have an index %d", list, to)
	} else if to < from || from < 0 {
		return nil, fmt.Errorf("%d is not a valid start index", from)
	}

	var out []Expr
	for _, r := range list[int(from):int(to)] {
		out = append(out, r)
	}
	return NewConsList(out...), nil
}

// String converts an expression to a string.
func (expr Cons) String() string {
	out := bytes.NewBufferString("(")
	expr.stringNoParen(out)
	out.WriteRune(')')
	return out.String()
}

func (expr Cons) stringNoParen(w io.Writer) {
	io.WriteString(w, ExprToString(expr.Car()))
	if expr.Cdr() == nil {
		return
	}
	io.WriteString(w, " ")
	if next, ok := expr.Cdr().(Cons); ok {
		next.stringNoParen(w)
	} else {
		io.WriteString(w, ". "+expr.Cdr().String())
	}
}

// Type converts the type of an expression to a string.
func (Cons) Type() string {
	return "cons"
}

// TypeAsm converts an Expr to its type code, without panicking on nil.
func (Cons) TypeAsm() byte {
	return 7
}

// ToSlice converts a cons-list into a slice.
func (expr Cons) ToSlice() []Expr {
	var out []Expr
	for {
		out = append(out, expr.Car())
		if expr.Cdr() == nil {
			break
		}
		expr = expr.Cdr().(Cons)
	}
	return out
}
