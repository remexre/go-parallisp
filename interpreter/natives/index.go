package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Index indexes or slices a sequence.
func Index(sequence types.Sequence, from types.Integer, others ...types.Expr) (types.Expr, error) {
	if len(others) > 1 {
		return nil, errors.New("@: too many arguments")
	} else if len(others) == 0 {
		return sequence.Get(from)
	}

	to, ok := others[0].(types.Integer)
	if !ok {
		return nil, errors.New("@: invalid arguments")
	}
	return sequence.Slice(from, to)
}
