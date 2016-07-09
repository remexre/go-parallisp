package builtins

import "remexre.xyz/go-parallisp/types"

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"@":        types.MustNewReflectFunction("@", Index),
	"+":        types.MustNewReflectFunction("+", Plus),
	"=":        types.MustNewReflectFunction("=", Eq),
	"apply":    types.SpecialFormFunc(Apply),
	"car":      types.MustNewReflectFunction("car", Car),
	"cdr":      types.MustNewReflectFunction("cdr", Cdr),
	"cond":     types.SpecialFormFunc(Cond),
	"cons":     types.MustNewReflectFunction("cons", types.NewCons),
	"def":      types.SpecialFormFunc(Def),
	"defmacro": types.SpecialFormFunc(Defmacro),
	"defun":    types.SpecialFormFunc(Defun),
	"error":    types.MustNewReflectFunction("error", Error),
	"gensym":   types.MustNewReflectFunction("gensym", Gensym),
	"int->str": types.MustNewReflectFunction("int->str", IntegerToString),
	"len":      types.MustNewReflectFunction("len", Len),
	"let":      types.SpecialFormFunc(Let),
	"lambda":   types.SpecialFormFunc(Lambda),
	"lst->vec": types.MustNewReflectFunction("lst->vec", ListToVector),
	"list":     types.MustNewReflectFunction("list", types.NewConsList),
	"mapvec":   types.SpecialFormFunc(MapVec),
	"nil":      nil,
	"println":  types.MustNewReflectFunction("println", Println),
	"quote":    types.SpecialFormFunc(Quote),
	"set":      types.SpecialFormFunc(Set),
	"type-of":  types.MustNewReflectFunction("type-of", Typeof),
	"vec->lst": types.MustNewReflectFunction("vec->lst", VectorToList),

	"**debug**": types.SpecialFormFunc(Debug),
})
