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

	args := expr.ToSlice()[1:]
	fn, err := expr.Car().Eval(env)
	if err != nil {
		return nil, err
	}

	if fn, ok := fn.(Function); ok {
		for i, arg := range args {
			args[i], err = arg.Eval(env)
			if err != nil {
				return nil, err
			}
		}
		return fn.Call(env, args...)
	} else if specialForm, ok := fn.(SpecialForm); ok {
		return specialForm.CallSpecialForm(env, args...)
	}

	return nil, fmt.Errorf("parallisp.types: cannot call non-function %s in %s", fn, expr)
}

// IsList returns true if the expression is a proper cons-list.
func (expr Cons) IsList() bool {
	if expr[1] == nil {
		return true
	} else if next, ok := expr[1].(Cons); ok {
		return next.IsList()
	}
	return false
}

// String converts an expression to a string.
func (expr Cons) String() string {
	out := bytes.NewBufferString("(")
	expr.stringNoParen(out)
	out.WriteRune(')')
	return out.String()
}

func (expr Cons) stringNoParen(w io.Writer) {
	io.WriteString(w, ExprToString(expr[0]))
	if expr[1] == nil {
		return
	}
	io.WriteString(w, " ")
	if next, ok := expr[1].(Cons); ok {
		next.stringNoParen(w)
	} else {
		io.WriteString(w, ExprToString(expr[1]))
	}
}

// Type converts the type of an expression to a string.
func (Cons) Type() string {
	return "cons"
}

// ToSlice converts a cons-list into a slice.
func (expr Cons) ToSlice() []Expr {
	var out []Expr
	for {
		out = append(out, expr[0])
		if expr[1] == nil {
			break
		}
		expr = expr[1].(Cons)
	}
	return out
}
