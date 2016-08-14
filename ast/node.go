package ast

import (
	"github.com/remexre/go-parallisp/types"
	"github.com/remexre/go-parallisp/util/exprset"
	"github.com/remexre/go-parallisp/util/stringset"
)

// Node represents a single AST node.
type Node interface {
	// Defines returns the symbols defined in the parent scope by this node,
	// recursively.
	Defines() stringset.StringSet

	// FreeVars returns the free values contained within a node, recursively.
	FreeVars() stringset.StringSet

	// IsLiteral returns whether the node is a literal. If it is, its Literals()
	// function will always return exactly one expression, which is equivalent to
	// this node.
	IsLiteral() bool

	// Literals returns the constants used in this node and all child nodes.
	Literals() exprset.ExprSet

	// ToExpr converts the node to an expr.
	ToExpr() types.Expr
}
