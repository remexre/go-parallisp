package natives

import (
	"fmt"

	"github.com/remexre/go-parallisp/types"
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
	fn, ok := fnArg.(types.Function)
	if !ok {
		return nil, fmt.Errorf("apply: not a function: %s", fn)
	}

	out := make(types.Vector, len(vec))
	for i, expr := range vec {
		var err error
		out[i], err = fn.Call(env, expr)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
