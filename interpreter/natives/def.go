package natives

import (
	"errors"

	"github.com/remexre/go-parallisp/debug"
	"github.com/remexre/go-parallisp/types"
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

	debug.Log("deprecated", "def is deprecated in %v",
		types.NewConsList(types.Symbol("def"), name, args[1]))
	if err := env.Def(name, value); err != nil {
		return nil, err
	}

	// Return.
	return nil, nil
}
