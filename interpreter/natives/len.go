package natives

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Len gets the length of a sequence.
func Len(expr types.Expr) (types.Integer, error) {
	if expr == nil {
		return 0, nil
	}
	if sequence, ok := expr.(types.Sequence); ok {
		return sequence.Len()
	}
	return 0, fmt.Errorf("len: cannot take the length of non-sequence %v", expr)
}
