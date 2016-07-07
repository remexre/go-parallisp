package builtins

import "remexre.xyz/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"@":        types.MustNewReflectFunction("@", Index),
	"def":      types.SpecialFormFunc(Def),
	"int->str": types.MustNewReflectFunction("int->str", IntegerToString),
	"let":      types.SpecialFormFunc(Let),
	"nil":      nil,
	"set":      types.SpecialFormFunc(Set),
	"println":  types.MustNewReflectFunction("println", Println),
})
