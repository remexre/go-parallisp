package debug

import (
	"errors"

	"remexre.xyz/go-parallisp/interpreter/types"
	"remexre.xyz/go-parallisp/types"
)

// Macro is the **debug-macro** special form.
func Macro(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) < 1 {
		return nil, errors.New("**debug-macro**: wrong argn")
	}

	mArg, err := types.EvalExpr(env, args[0])
	if err != nil {
		return nil, err
	}
	m, ok := mArg.(*interpreterTypes.FunctionLike)
	if !ok {
		return nil, errors.New("**debug-macro**: wrong arg type")
	}

	return m.CallRaw(env, args[1:]...)
}
