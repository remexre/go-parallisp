package interpreter

import (
	"remexre.xyz/go-parallisp/parser"
	"remexre.xyz/go-parallisp/types"
)

// Interpret interprets source code with a default environment from a given
// file.
func Interpret(src, file string) (types.Expr, types.Env, error) {
	exprs, err := parser.Parse(src)
	if err != nil {
		return nil, nil, err
	}

	env := NewEnv()
	if file != "" {
		env.Def(types.Symbol("**current-file**"), types.String(file))
	}
	env = env.Derive(nil)

	var out types.Expr
	for _, expr := range exprs {
		out, err = types.EvalExpr(env, expr)
		if err != nil {
			return nil, nil, err
		}
	}
	return out, env, nil
}
