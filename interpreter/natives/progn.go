package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Progn is the progn special form.
func Progn(env types.Env, exprs ...types.Expr) (types.Expr, error) {
	if len(exprs) == 0 {
		return nil, errors.New("progn: incorrect usage")
	}

	// Create an inner environment.
	env = env.Derive(nil)

	// Evaluate the expressions.
	var out types.Expr
	for _, expr := range exprs {
		var err error
		out, err = types.EvalExpr(env, expr)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
