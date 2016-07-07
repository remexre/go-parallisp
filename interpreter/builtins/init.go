package builtins

import "remexre.xyz/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"@":        types.MustNewReflectFunction("@", Index),
	"cond":     types.SpecialFormFunc(Cond),
	"cons":     types.MustNewReflectFunction("cons", types.NewCons),
	"def":      types.SpecialFormFunc(Def),
	"defmacro": types.SpecialFormFunc(Defmacro),
	"error":    types.MustNewReflectFunction("error", Error),
	"int->str": types.MustNewReflectFunction("int->str", IntegerToString),
	"let":      types.SpecialFormFunc(Let),
	"lst->vec": types.MustNewReflectFunction("lst->vec", ListToVector),
	"list":     types.MustNewReflectFunction("list", types.NewConsList),
	"nil":      nil,
	"println":  types.MustNewReflectFunction("println", Println),
	"quote":    types.SpecialFormFunc(Quote),
	"set":      types.SpecialFormFunc(Set),
	"type-of":  types.MustNewReflectFunction("type-of", Typeof),
	"vec->lst": types.MustNewReflectFunction("vec->lst", VectorToList),
})
