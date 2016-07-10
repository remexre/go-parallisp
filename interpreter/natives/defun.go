package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/interpreter/types"
	"remexre.xyz/go-parallisp/types"
)

// Defun creates a named function.
func Defun(env types.Env, defunArgs ...types.Expr) (types.Expr, error) {
	if len(defunArgs) < 2 {
		return nil, errors.New("defun: too few arguments")
	}

	name, ok := defunArgs[0].(types.Symbol)
	if !ok {
		return nil, errors.New("defun: invalid name")
	}

	args, ok := defunArgs[1].(types.Vector)
	if !ok {
		return nil, errors.New("defun: invalid args")
	}

	f, err := interpreterTypes.NewFunctionLike(false, env, name, args, defunArgs[2:]...)
	if err != nil {
		return nil, err
	}
	return f, env.Def(name, f)
}
