package exec

import "github.com/remexre/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"exec": types.MustNewReflectFunction("exec", Exec),
})
