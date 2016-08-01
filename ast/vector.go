package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Vector is a vector literal.
type Vector []Node

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Vector) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (vec *Vector) FreeVars() stringset.StringSet {
	var freeVars stringset.StringSet
	for _, node := range *vec {
		freeVars = freeVars.Union(node.FreeVars())
	}
	return freeVars
}

// ToExpr converts the node to an expr.
func (vec *Vector) ToExpr() types.Expr {
	exprs := make(types.Vector, len(*vec))
	for i, node := range *vec {
		exprs[i] = node.ToExpr()
	}
	return exprs
}
