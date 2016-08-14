package process

import (
	"os"

	"github.com/remexre/go-parallisp/types"
)

func argsToVector(argSlice []string) types.Expr {
	args := make(types.Vector, len(os.Args))
	for i, arg := range os.Args {
		args[i] = types.String(arg)
	}
	return args
}
