package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Vector is a vector literal.
type Vector []Node

// Constants returns the constants used in this node and all child nodes.
func (v *Vector) Constants() exprset.ExprSet {
	sets := make([]exprset.ExprSet, len(*v))
	for i, node := range *v {
		sets[i] = node.Constants()
	}
	return exprset.Union(sets...)
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Vector) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (v *Vector) FreeVars() stringset.StringSet {
	var freeVars stringset.StringSet
	for _, node := range *v {
		freeVars = freeVars.Union(node.FreeVars())
	}
	return freeVars
}

// ToExpr converts the node to an expr.
func (v *Vector) ToExpr() types.Expr {
	exprs := make(types.Vector, len(*v))
	for i, node := range *v {
		exprs[i] = node.ToExpr()
	}
	return exprs
}
