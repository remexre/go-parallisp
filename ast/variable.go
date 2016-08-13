package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Variable is a variable.
type Variable string

// Literals returns the constants used in this node and all child nodes.
func (v *Variable) Literals() exprset.ExprSet {
	return nil
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Variable) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (v *Variable) FreeVars() stringset.StringSet {
	return stringset.New(string(*v))
}

// ToExpr converts the node to an expr.
func (v *Variable) ToExpr() types.Expr {
	return types.Symbol(*v)
}
