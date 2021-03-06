package types

import "fmt"

// Floating is a 64-bit IEEE 754 floating-point number.
type Floating float64

// Eval evaluates an expression.
func (expr Floating) Eval(env Env) (Expr, error) {
	return expr, nil
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (expr Floating) LiteralAsm() string {
	f := float64(expr)
	return fmt.Sprintf(".double %f", f)
}

// String converts an expression to a string.
func (expr Floating) String() string {
	return fmt.Sprintf("%f", float64(expr))
}

// Type converts the type of an expression to a string.
func (Floating) Type() string {
	return "floating"
}

// TypeAsm converts an Expr to its type code, without panicking on nil.
func (Floating) TypeAsm() byte {
	return 2
}
