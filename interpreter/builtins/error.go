package builtins

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Error causes an error.
func Error(exprs ...types.Expr) (types.Expr, error) {
	args := make([]interface{}, len(exprs))
	for i, expr := range exprs {
		args[i] = expr
	}
	return nil, errors.New(fmt.Sprint(args...))
}
