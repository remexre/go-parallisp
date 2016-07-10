package natives

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

// DebugMacro is the **debug-macro** special form.
func DebugMacro(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) < 1 {
		return nil, errors.New("**debug-macro**: wrong argn")
	}

	mArg, err := types.EvalExpr(env, args[0])
	if err != nil {
		return nil, err
	}
	m, ok := mArg.(*macro)
	if !ok {
		return nil, errors.New("**debug-macro**: wrong arg type")
	}

	return m.CallMacro(env, args[1:]...)
}
