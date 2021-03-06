package interpreter

import (
	"github.com/remexre/go-parallisp/parser"
	"github.com/remexre/go-parallisp/types"
)

// Interpret interprets source code with a default environment from a given
// file.
func Interpret(src, file string) (types.Expr, types.Env, error) {
	exprs, err := parser.Parse(src)
	if err != nil {
		return nil, nil, err
	}

	fileSpecificVars := map[types.Symbol]types.Expr{
		"**current-file**": nil,
	}
	if file != "" {
		fileSpecificVars[types.Symbol("**current-file**")] = types.String(file)
	}
	env := NewEnv(fileSpecificVars).Derive(nil)

	var out types.Expr
	for _, expr := range exprs {
		out, err = types.EvalExpr(env, expr)
		if err != nil {
			return nil, nil, err
		}
	}
	return out, env, nil
}
