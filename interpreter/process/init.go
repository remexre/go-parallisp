package process

import (
	"os"

	"github.com/remexre/go-parallisp/types"
)

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"args": argsToVector(os.Args),
})
