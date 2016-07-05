package types

import "fmt"

// Floating is a 64-bit IEEE 754 floating-point number.
type Floating float64

// String converts an expression to a string.
func (expr Floating) String() string {
	return fmt.Sprintf("%f", float64(expr))
}

// Type converts the type of an expression to a string.
func (Floating) Type() string {
	return "floating"
}
