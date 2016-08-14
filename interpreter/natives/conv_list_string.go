package natives

import "github.com/remexre/go-parallisp/types"

// ListToString converts a list of integer code points into a Unicode string.
func ListToString(l types.Expr) types.String {
	var exprs []types.Expr
	if l != nil {
		exprs = l.(types.Cons).ToSlice()
	}

	runes := make([]rune, len(exprs))
	for i, expr := range exprs {
		num := expr.(types.Integer)
		runes[i] = rune(num)
	}
	return types.String(runes)
}
