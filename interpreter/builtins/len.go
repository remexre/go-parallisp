package builtins

import "remexre.xyz/go-parallisp/types"

// Len gets the length of a sequence.
func Len(sequence types.Sequence) (types.Integer, error) {
	return sequence.Len()
}
