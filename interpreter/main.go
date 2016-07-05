package interpreter

import (
	"remexre.xyz/parallisp/parser"
	"remexre.xyz/parallisp/types"
)

// Interpret interprets source code with a default environment.
func Interpret(src string) (types.Expr, error) {
	exprs, err := parser.Parse(src)
	if err != nil {
		return nil, err
	}

	env := NewEnv()
	var out types.Expr
	for _, expr := range exprs {
		out, err = expr.Eval(env)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
