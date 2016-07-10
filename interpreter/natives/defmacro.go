package natives

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Defmacro defines a macro.
func Defmacro(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) < 3 {
		return nil, errors.New("defmacro: too few arguments")
	}

	name, ok := args[0].(types.Symbol)
	if !ok {
		return nil, errors.New("defmacro: invalid name")
	}

	macroArgsIn, ok := args[1].(types.Vector)
	if !ok {
		return nil, errors.New("defmacro: invalid args")
	}
	var macroArgs []types.Symbol
	macroVariadic := false
	for _, arg := range macroArgsIn {
		sym, ok := arg.(types.Symbol)
		if !ok {
			return nil, errors.New("defmacro: invalid arg")
		}
		if sym[0] == '&' {
			switch string(sym) {
			case "&rest":
				macroVariadic = true
			default:
				return nil, fmt.Errorf("defmacro: unrecognized metaarg %s", sym)
			}
		} else {
			macroArgs = append(macroArgs, sym)
		}
	}

	m := &macro{
		name:     name,
		args:     macroArgs,
		variadic: macroVariadic,
		body:     args[2:],
	}
	return m, env.Def(name, m)
}
