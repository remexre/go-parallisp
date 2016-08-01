package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A String is a string constant.
type String string

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (String) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (String) FreeVars() stringset.StringSet { return nil }

// ToExpr converts the node to an expr.
func (s String) ToExpr() types.Expr {
	return types.String(s)
}
