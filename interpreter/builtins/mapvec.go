package builtins

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// MapVec uses a function to map a vector.
func MapVec(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("mapvec: invalid argn")
	}

	vecArg, err := types.EvalExpr(env, args[1])
	if err != nil {
		return nil, err
	}
	vec, ok := vecArg.(types.Vector)
	if !ok {
		return nil, fmt.Errorf("mapvec: not a vector: %s", vecArg)
	}

	fnArg, err := types.EvalExpr(env, args[0])
	if err != nil {
		return nil, err
	}
	var f func(expr types.Expr) (types.Expr, error)
	switch fn := fnArg.(type) {
	case types.Function:
		f = func(expr types.Expr) (types.Expr, error) { return fn.Call(expr) }
	case types.SpecialForm:
		f = func(expr types.Expr) (types.Expr, error) { return fn.CallSpecialForm(env, expr) }
	default:
		return nil, fmt.Errorf("mapvec: not a function: %s", fn)
	}

	out := make(types.Vector, len(vec))
	for i, expr := range vec {
		var err error
		out[i], err = f(expr)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
