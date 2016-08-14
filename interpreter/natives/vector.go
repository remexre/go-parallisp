package natives

import "github.com/remexre/go-parallisp/types"

// Vector creates a vector.
func Vector(exprs ...types.Expr) types.Expr {
	return types.Vector(exprs)
}
