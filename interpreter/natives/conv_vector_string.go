package natives

import "github.com/remexre/go-parallisp/types"

// VectorToString converts a vector of integer code points into a Unicode
// string.
func VectorToString(v types.Vector) types.String {
	runes := make([]rune, len(v))
	for i, expr := range v {
		runes[i] = rune(expr.(types.Integer))
	}
	return types.String(runes)
}
