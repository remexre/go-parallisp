package ast

import (
	"fmt"

	"github.com/remexre/go-parallisp/types"
	"github.com/remexre/go-parallisp/util/exprset"
	"github.com/remexre/go-parallisp/util/stringset"
)

// Quote represents a quoted expression.
type Quote struct {
	Value types.Expr
}

// NewQuote returns a new quote from the expressions in its form, excluding the
// initial quote symbol.
func NewQuote(exprs []types.Expr) (Node, error) {
	if len(exprs) != 1 {
		return nil, fmt.Errorf("ast.Convert: invalid quote")
	}
	return &Quote{exprs[0]}, nil
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (q *Quote) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (*Quote) FreeVars() stringset.StringSet { return nil }

// IsLiteral returns whether the node is a literal.
func (*Quote) IsLiteral() bool { return true }

// Literals returns the constants used in this node and all child nodes.
func (q *Quote) Literals() exprset.ExprSet {
	return exprset.New(q.Value)
}

// ToExpr converts the node to an expr.
func (q *Quote) ToExpr() types.Expr {
	return types.NewConsList(
		types.Symbol("quote"),
		q.Value,
	)
}
