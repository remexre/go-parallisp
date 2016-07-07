package builtins

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// Def is the def special form.
func Def(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, errors.New("def: wrong argument count")
	}

	name, ok := args[0].(types.Symbol)
	if !ok {
		return nil, errors.New("def: name not a symbol")
	}

	value, err := types.EvalExpr(env, args[1])
	if err != nil {
		return nil, err
	}

	if err := env.Def(name, value); err != nil {
		return nil, err
	}

	// Return.
	return nil, nil
}
