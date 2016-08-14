package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Vector is a vector literal.
type Vector []Node

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

// IsLiteral returns whether the node is a literal.
func (v *Vector) IsLiteral() bool {
	for _, child := range *v {
		if !child.IsLiteral() {
			return false
		}
	}
	return true
}

// Literals returns the constants used in this node and all child nodes.
func (v *Vector) Literals() exprset.ExprSet {
	if v.IsLiteral() {
		return exprset.New(v.ToExpr())
	}
	sets := make([]exprset.ExprSet, len(*v))
	for i, node := range *v {
		sets[i] = node.Literals()
	}
	return exprset.Union(sets...)
}

// ToExpr converts the node to an expr.
func (v *Vector) ToExpr() types.Expr {
	exprs := make(types.Vector, len(*v))
	for i, node := range *v {
		exprs[i] = node.ToExpr()
	}
	return exprs
}
