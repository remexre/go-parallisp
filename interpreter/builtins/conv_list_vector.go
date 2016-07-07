package builtins

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// ListToVector converts a cons-list to a vector.
func ListToVector(list types.Cons) (types.Expr, error) {
	if !list.IsList() {
		return nil, fmt.Errorf("lst->vec: not a list: %s", list)
	}
	return types.Vector(list.ToSlice()), nil
}
