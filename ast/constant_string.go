package ast

import (
	"github.com/remexre/go-parallisp/types"
	"github.com/remexre/go-parallisp/util/exprset"
	"github.com/remexre/go-parallisp/util/stringset"
)

// A String is a string constant.
type String string

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (String) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (String) FreeVars() stringset.StringSet { return nil }

// IsLiteral returns whether the node is a literal.
func (String) IsLiteral() bool { return true }

// Literals returns the constants used in this node and all child nodes.
func (s String) Literals() exprset.ExprSet {
	return exprset.New(types.String(s))
}

// ToExpr converts the node to an expr.
func (s String) ToExpr() types.Expr {
	return types.String(s)
}
