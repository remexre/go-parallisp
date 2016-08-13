package types

import "fmt"

// Integer is a 64-bit two's-complement signed integer.
type Integer int64

// Eval evaluates an expression.
func (expr Integer) Eval(env Env) (Expr, error) {
	return expr, nil
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (expr Integer) LiteralAsm() string {
	return fmt.Sprintf(".quad %d", int64(expr))
}

// String converts an expression to a string.
func (expr Integer) String() string {
	return fmt.Sprintf("%d", int64(expr))
}

// Type converts the type of an expression to a string.
func (Integer) Type() string {
	return "integer"
}
