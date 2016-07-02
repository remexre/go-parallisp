package ast

import "remexre.xyz/parallisp/types"

// ToAST converts an expression or expressions to an AST node.
func ToAST(exprs ...types.Expr) Node {
	env := map[string]*Variable{
	// TODO Default environment
	}

	// If there's only one expression, convert and return it.
	if len(exprs) == 1 {
		return oneToAST(exprs[0], env)
	}

	// Otherwise, return a progn of all the expressions.
	progn := make(ProgN, len(exprs))
	for i, expr := range exprs {
		progn[i] = oneToAST(expr, env)
	}
	return progn
}

func oneToAST(expr types.Expr, env map[string]*Variable) Node {
	// TODO
	panic(expr)
}
