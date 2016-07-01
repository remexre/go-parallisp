package parallisp

import "fmt"

// Integer is a 64-bit two's-complement signed integer.
type Integer int64

// Expr converts an expression to a string.
func (expr Integer) Expr() string {
	return fmt.Sprintf("%d", int64(expr))
}
