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

	listArg, err := types.EvalExpr(env, args[1])
	if err != nil {
		return nil, err
	}
	list, ok := listArg.(types.Cons)
	if !ok {
		return nil, fmt.Errorf("apply: not a list: %s", list)
	}
	if !list.IsList() {
		return nil, fmt.Errorf("apply: not a list: %s", list)
	}

	fnArg, err := types.EvalExpr(env, args[0])
	if err != nil {
		return nil, err
	}
	switch fn := fnArg.(type) {
	case types.Function:
		return fn.Call(list.ToSlice()...)
	case types.SpecialForm:
		return fn.CallSpecialForm(env, list.ToSlice()...)
	default:
		return nil, fmt.Errorf("apply: not a function: %s", fn)
	}
}
