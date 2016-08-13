package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A Floating is an floating-point constant.
type Floating float64

// Literals returns the constants used in this node and all child nodes.
func (f Floating) Literals() exprset.ExprSet {
	return exprset.New(types.Floating(f))
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (Floating) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (Floating) FreeVars() stringset.StringSet { return nil }

// ToExpr converts the node to an expr.
func (f Floating) ToExpr() types.Expr {
	return types.Floating(f)
}
