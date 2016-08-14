package fs

import "github.com/remexre/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"open":  types.MustNewReflectFunction("open", Open),
	"close": types.MustNewReflectFunction("close", Close),

	"read-all": types.MustNewReflectFunction("read-all", ReadAll),
})
