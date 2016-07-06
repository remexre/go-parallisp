package fs

import "remexre.xyz/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"open":  types.MustNewReflectFunction(Open),
	"close": types.MustNewReflectFunction(Open), // TODO
})
