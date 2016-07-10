package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/interpreter/types"
	"remexre.xyz/go-parallisp/types"
)

// Lambda creates an anonymous function.
func Lambda(env types.Env, lambdaArgs ...types.Expr) (types.Expr, error) {
	if len(lambdaArgs) < 2 {
		return nil, errors.New("lambda: too few arguments")
	}

	name := Gensym("lambda")
	args, ok := lambdaArgs[0].(types.Vector)
	if !ok {
		return nil, errors.New("lambda: invalid args")
	}

	f, err := interpreterTypes.NewFunctionLike(false, env, name, args, lambdaArgs[1:]...)
	if err != nil {
		return nil, err
	}
	return f, env.Def(name, f)
}
