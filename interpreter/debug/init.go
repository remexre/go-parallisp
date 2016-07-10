package debug

import "remexre.xyz/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"**debug**":       types.NativeFunc(Debug),
	"**debug-macro**": types.NativeFunc(Macro),
})
