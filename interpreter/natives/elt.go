package natives

import (
	"errors"

	"github.com/remexre/go-parallisp/types"
)

// Elt indexes or slices a sequence.
func Elt(sequence types.Sequence, from types.Integer, others ...types.Expr) (types.Expr, error) {
	if len(others) > 1 {
		return nil, errors.New("elt: too many arguments")
	} else if len(others) == 0 {
		return sequence.Get(from)
	}

	to, ok := others[0].(types.Integer)
	if !ok {
		return nil, errors.New("elt: invalid arguments")
	}
	return sequence.Slice(from, to)
}
