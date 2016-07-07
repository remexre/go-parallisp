package builtins

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Cond is the cond special form.
func Cond(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return nil, errors.New("cond: invalid argn")
	} else if len(args)%2 == 0 {
		args = append(args, nil)
	}
	for i := 0; i < len(args)-1; i += 2 {
		pred, err := types.EvalExpr(env, args[i])
		if err != nil {
			return nil, err
		} else if pred != nil {
			return types.EvalExpr(env, args[i+1])
		}
	}
	return types.EvalExpr(env, args[len(args)-1])
}
