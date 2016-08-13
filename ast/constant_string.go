package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A String is a string constant.
type String string

// Literals returns the constants used in this node and all child nodes.
func (s String) Literals() exprset.ExprSet {
	return exprset.New(types.String(s))
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (String) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (String) FreeVars() stringset.StringSet { return nil }

// ToExpr converts the node to an expr.
func (s String) ToExpr() types.Expr {
	return types.String(s)
}
