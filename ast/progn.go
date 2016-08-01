package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Progn is a sequence of sequentially evaluated expressions.
type Progn []Node

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Progn) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (p *Progn) FreeVars() stringset.StringSet {
	sets := make([]stringset.StringSet, len(*p))
	for i, node := range *p {
		sets[i] = node.FreeVars()
	}
	return stringset.Union(sets...)
}

// ToExpr converts the node to an expr.
func (p *Progn) ToExpr() types.Expr {
	exprs := make([]types.Expr, len(*p)+1)
	for i, node := range *p {
		exprs[i+1] = node.ToExpr()
	}
	exprs[0] = types.Symbol("progn")
	return types.NewConsList(exprs...)
}
