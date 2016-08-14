package natives

import (
	"errors"

	"github.com/remexre/go-parallisp/types"
)

// Quote returns its argument literally.
func Quote(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) != 1 {
		return nil, errors.New("quote: incorrect argn")
	}
	return args[0], nil
}
