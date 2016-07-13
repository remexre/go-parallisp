package natives

import "remexre.xyz/go-parallisp/types"

// StringToList converts a string into a list of Unicode code points.
func StringToList(str types.String) types.Expr {
	bytes := []byte(str)
	exprs := make([]types.Expr, len(bytes))
	for i, b := range bytes {
		exprs[i] = types.Integer(b)
	}
	return types.NewConsList(exprs...)
}
