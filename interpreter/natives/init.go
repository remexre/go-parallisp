package natives

import (
	"remexre.xyz/go-parallisp/interpreter/types"
	"remexre.xyz/go-parallisp/types"
)

// Env is the environment exported by the module.
var Env = types.NewRootEnv(map[types.Symbol]types.Expr{
	"@":        types.MustNewReflectFunction("@", Index),
	"+":        types.MustNewReflectFunction("+", Add),
	"-":        types.MustNewReflectFunction("-", Subtract),
	"*":        types.MustNewReflectFunction("*", Multiply),
	"/":        types.MustNewReflectFunction("/", Divide),
	"%":        types.MustNewReflectFunction("%", Modulo),
	">":        types.MustNewReflectFunction(">", Gt),
	">=":       types.MustNewReflectFunction(">=", Gte),
	"<":        types.MustNewReflectFunction("<", Lt),
	"<=":       types.MustNewReflectFunction("<=", Lte),
	"=":        types.MustNewReflectFunction("=", Eq),
	"apply":    types.NativeFunc(Apply),
	"car":      types.MustNewReflectFunction("car", Car),
	"cdr":      types.MustNewReflectFunction("cdr", Cdr),
	"cond":     types.NativeFunc(Cond),
	"cons":     types.MustNewReflectFunction("cons", types.NewCons),
	"def":      types.NativeFunc(Def),
	"defmacro": types.NativeFunc(Defmacro),
	"defun":    types.NativeFunc(Defun),
	"eval":     types.NativeFunc(Eval),
	"gensym":   types.MustNewReflectFunction("gensym", Gensym),
	"int->str": types.MustNewReflectFunction("int->str", IntegerToString),
	"len":      types.MustNewReflectFunction("len", Len),
	"let":      types.NativeFunc(Let),
	"lambda":   types.NativeFunc(Lambda),
	"list":     types.MustNewReflectFunction("list", types.NewConsList),
	"mapvec":   types.NativeFunc(MapVec),
	"nil":      nil,
	"progn":    types.NativeFunc(interpreterTypes.Progn),
	"quote":    types.NativeFunc(Quote),
	"set":      types.NativeFunc(Set),
	"string":   types.MustNewReflectFunction("string", String),
	"type-of":  types.MustNewReflectFunction("type-of", Typeof),
	"vec->lst": types.MustNewReflectFunction("vec->lst", VectorToList),
	"vector":   types.MustNewReflectFunction("vector", Vector),

	"**error**": types.MustNewReflectFunction("**error**", Error),
	"**print**": types.MustNewReflectFunction("**print**", Print),
})
