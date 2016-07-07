package builtins

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Index indexes a string, vector, or any other type that meets the types.Slicer
// interface.
func Index(slicer types.Slicer, from types.Integer, others ...types.Expr) (types.Expr, error) {
	if len(others) > 1 {
		return nil, errors.New("@: too many arguments")
	} else if len(others) == 0 {
		return slicer.Get(from)
	}

	to, ok := others[1].(types.Integer)
	if !ok {
		return nil, errors.New("@: invalid arguments")
	}
	return slicer.Slice(from, to)
}
