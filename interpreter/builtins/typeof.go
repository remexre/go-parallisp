package builtins

import "remexre.xyz/go-parallisp/types"

// Typeof returns the type of an expression.
func Typeof(expr types.Expr) types.Symbol {
	return types.Symbol(expr.Type())
}
