package natives

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Lambda creates an anonymous function.
func Lambda(env types.Env, lambdaArgs ...types.Expr) (types.Expr, error) {
	if len(lambdaArgs) < 2 {
		return nil, errors.New("lambda: too few arguments")
	}

	name, err := Gensym("lambda")
	if err != nil {
		return nil, err
	}

	argsIn, ok := lambdaArgs[0].(types.Vector)
	if !ok {
		return nil, errors.New("lambda: invalid args")
	}
	var args []types.Symbol
	variadic := false
	for _, arg := range argsIn {
		sym, ok := arg.(types.Symbol)
		if !ok {
			return nil, errors.New("lambda: invalid arg")
		}
		if sym[0] == '&' {
			switch string(sym) {
			case "&rest":
				variadic = true
			default:
				return nil, fmt.Errorf("lambda: unrecognized metaarg %s", sym)
			}
		} else {
			args = append(args, sym)
		}
	}

	fn := &function{
		name:     name,
		args:     args,
		variadic: variadic,

		env:  env,
		body: lambdaArgs[1:],
	}
	return fn, nil
}
