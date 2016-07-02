package types

import "strconv"

// String is a dynamically sized array of bytes.
type String string

// Expr converts an expression to a string.
func (expr String) Expr() string {
	return strconv.Quote(string(expr))
}
