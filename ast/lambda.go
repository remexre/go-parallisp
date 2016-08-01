package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/stringset"
)

// Lambda represents a lambda expression which accepts multiple parameters.
type Lambda struct {
	Params []string
	Doc    string
	Body   Progn
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Lambda) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (l *Lambda) FreeVars() stringset.StringSet {
	return l.Body.FreeVars().Difference(stringset.New(l.Params...))
}

// ToExpr converts the node to an expr.
func (l *Lambda) ToExpr() types.Expr {
	// Build the parameter list.
	params := make([]types.Expr, len(l.Params))
	for i, param := range l.Params {
		params[i] = types.Symbol(param)
	}

	// Build the body.
	body := make([]types.Expr, len(l.Body))
	for i, node := range l.Body {
		body[i] = node.ToExpr()
	}

	// Build the defmacro and return.
	return types.NewConsList(append([]types.Expr{
		types.Symbol("lambda"),
		types.Vector(params),
		types.String(l.Doc),
	}, body...)...)
}