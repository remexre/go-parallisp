package builtins

import "remexre.xyz/go-parallisp/types"

// VectorToList converts a vector to a cons-list.
func VectorToList(vec types.Vector) types.Expr {
	return types.NewConsList(vec...)
}
