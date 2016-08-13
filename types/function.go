package types

import (
	"errors"
	"fmt"
)

// Function is a callable interface.
type Function interface {
	Expr

	Call(Env, ...Expr) (Expr, error)
}

// NativeFunc is an implementation of Function based on a Go function.
type NativeFunc func(Env, ...Expr) (Expr, error)

// Call calls the wrapped function.
func (fn NativeFunc) Call(env Env, exprs ...Expr) (Expr, error) {
	return fn(env, exprs...)
}

// Eval evaluates a special form.
func (fn NativeFunc) Eval(env Env) (Expr, error) {
	return nil, errors.New("parallisp.types: cannot eval a special form")
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (NativeFunc) LiteralAsm() string {
	panic("types.Function.LiteralAsm: cannot make a function into a literal")
}

// String converts the special form to a string.
func (fn NativeFunc) String() string {
	var f func(Env, ...Expr) (Expr, error)
	f = fn
	return fmt.Sprintf("native-func-%p", f)
}

// Type returns the type of the special form.
func (fn NativeFunc) Type() string {
	return "special-form"
}

// TypeAsm converts an Expr to its type code, without panicking on nil.
func (NativeFunc) TypeAsm() byte {
	return 6
}
