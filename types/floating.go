package types

import "fmt"

// Floating is a 64-bit IEEE 754 floating-point number.
type Floating float64

// Expr converts an expression to a string.
func (expr Floating) Expr() string {
	return fmt.Sprintf("%f", float64(expr))
}
