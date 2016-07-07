package builtins

import "remexre.xyz/go-parallisp/types"

// IntegerToString converts an integer code point into a Unicode string.
func IntegerToString(n types.Integer) types.String {
	return types.String(rune(n))
}
