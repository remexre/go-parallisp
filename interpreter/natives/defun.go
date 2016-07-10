package natives

import (
	"errors"
	"fmt"

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

	argsIn, ok := defunArgs[1].(types.Vector)
	if !ok {
		return nil, errors.New("defun: invalid args")
	}
	var args []types.Symbol
	variadic := false
	for _, arg := range argsIn {
		sym, ok := arg.(types.Symbol)
		if !ok {
			return nil, errors.New("defun: invalid arg")
		}
		if sym[0] == '&' {
			switch string(sym) {
			case "&rest":
				variadic = true
			default:
				return nil, fmt.Errorf("defun: unrecognized metaarg %s", sym)
			}
		} else {
			args = append(args, sym)
		}
	}

	return nil, env.Def(name, &function{
		name:     name,
		args:     args,
		variadic: variadic,

		env:  env,
		body: defunArgs[2:],
	})
}
