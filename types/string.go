package types

import "strconv"

// String is a dynamically sized array of bytes.
type String string

// String converts an expression to a string.
func (expr String) String() string {
	return strconv.Quote(string(expr))
}

// Type converts the type of an expression to a string.
func (String) Type() string {
	return "string"
}
