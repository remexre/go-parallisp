package natives

import "github.com/remexre/go-parallisp/types"

// IntegerToString converts an integer code point into a Unicode string.
func IntegerToString(n types.Integer) types.String {
	return types.String(rune(n))
}
