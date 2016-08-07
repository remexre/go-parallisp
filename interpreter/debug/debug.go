package debug

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Debug is the **debug** special form.
func Debug(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 1 {
		return nil, errors.New("**debug**: wrong argn")
	}
	what, ok := args[0].(types.Symbol)
	if !ok {
		return nil, errors.New("**debug**: wrong arg type")
	}
	switch what {
	case "env":
		return types.String(fmt.Sprint(env)), nil
	case "env-list":
		var vector types.Vector
		for _, sym := range env.List(false) {
			vector = append(vector, sym)
		}
		return vector, nil
	}
	return nil, fmt.Errorf("**debug**: unknown arg %s", what)
}
