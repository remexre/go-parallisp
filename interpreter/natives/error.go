package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Error causes an error.
func Error(err types.String) (types.Expr, error) {
	return nil, errors.New(string(err))
}
