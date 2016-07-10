package natives

import "remexre.xyz/go-parallisp/types"

// String converts an expression to its string representation.
func String(expr types.Expr) types.String {
	return types.String(types.ExprToString(expr))
}
