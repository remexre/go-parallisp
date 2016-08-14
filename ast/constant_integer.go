package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Integer is an integer constant.
type Integer int64

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (Integer) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (Integer) FreeVars() stringset.StringSet { return nil }

// IsLiteral returns whether the node is a literal.
func (Integer) IsLiteral() bool { return true }

// Literals returns the constants used in this node and all child nodes.
func (i Integer) Literals() exprset.ExprSet {
	return exprset.New(types.Integer(i))
}

// ToExpr converts the node to an expr.
func (i Integer) ToExpr() types.Expr {
	return types.Integer(i)
}
