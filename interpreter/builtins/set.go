package builtins

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Set is the set special form.
func Set(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, errors.New("set: wrong argument count")
	}

	name, ok := args[0].(types.Symbol)
	if !ok {
		return nil, errors.New("set: name not a symbol")
	}

	value, err := types.EvalExpr(env, args[1])
	if err != nil {
		return nil, err
	}

	if err := env.Set(name, value); err != nil {
		return nil, err
	}

	// Return.
	return nil, nil
}
