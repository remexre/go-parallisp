package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/interpreter/types"
	"remexre.xyz/go-parallisp/types"
)

// Defmacro defines a macro.
func Defmacro(env types.Env, defmacroArgs ...types.Expr) (types.Expr, error) {
	if len(defmacroArgs) < 2 {
		return nil, errors.New("defmacro: too few arguments")
	}

	name, ok := defmacroArgs[0].(types.Symbol)
	if !ok {
		return nil, errors.New("defmacro: invalid name")
	}

	args, ok := defmacroArgs[1].(types.Vector)
	if !ok {
		return nil, errors.New("defmacro: invalid args")
	}

	f, err := interpreterTypes.NewFunctionLike(true, env, name, args, defmacroArgs[2:]...)
	if err != nil {
		return nil, err
	}
	return f, env.Def(name, f)
}
