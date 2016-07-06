package types

import (
	"errors"
	"fmt"
)

// SpecialForm is a callable interface that receives its arguments unevaluated.
type SpecialForm interface {
	Expr

	CallSpecialForm(Env, ...Expr) (Expr, error)
}

// SpecialFormFunc is an implementation of SpecialForm based on a function.
type SpecialFormFunc func(Env, ...Expr) (Expr, error)

// CallSpecialForm calls the wrapped function.
func (fn SpecialFormFunc) CallSpecialForm(env Env, exprs ...Expr) (Expr, error) {
	return fn(env, exprs...)
}

// Eval evaluates a special form.
func (fn SpecialFormFunc) Eval(env Env) (Expr, error) {
	return nil, errors.New("parallisp.types: cannot eval a special form")
}

// String converts the special form to a string.
func (fn SpecialFormFunc) String() string {
	var f func(Env, ...Expr) (Expr, error)
	f = fn
	return fmt.Sprintf("special-form-%p", f)
}

// Type returns the type of the special form.
func (fn SpecialFormFunc) Type() string {
	return "special-form"
}
