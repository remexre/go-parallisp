package natives

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Apply applies a list of arguments to a function.
func Apply(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("apply: invalid argn")
	}

	fnArg, err := types.EvalExpr(env, args[0])
	if err != nil {
		return nil, err
	}
	fn, ok := fnArg.(types.Function)
	if !ok {
		return nil, fmt.Errorf("apply: not a function: %s", fn)
	}

	listArg, err := types.EvalExpr(env, args[1])
	if err != nil {
		return nil, err
	}
	var list []types.Expr
	if listArg != nil {
		listCons, ok := listArg.(types.Cons)
		if !ok {
			return nil, fmt.Errorf("apply: not a list: %s", listCons)
		}
		if !listCons.IsList() {
			return nil, fmt.Errorf("apply: not a list: %s", listCons)
		}
		list = listCons.ToSlice()
	}

	fnArgs := make([]types.Expr, len(list))
	for i, arg := range list {
		fnArgs[i] = types.NewConsList(types.Symbol("quote"), arg)
	}
	return fn.Call(env, fnArgs...)
}
