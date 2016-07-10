package process

import (
	"os"

	"remexre.xyz/go-parallisp/types"
)

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"args": argsToVector(os.Args),
})
