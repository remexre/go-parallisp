package natives

import (
	"errors"

	"github.com/remexre/go-parallisp/types"
)

// Eval evaluates a form.
func Eval(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 1 {
		return nil, errors.New("eval: wrong argn")
	}

	expr, err := types.EvalExpr(env, args[0])
	if err != nil {
		return nil, err
	}

	return types.EvalExpr(env, expr)
}
