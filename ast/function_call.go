package ast

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/stringset"
)

// A FunctionCall is the evaluation of an expression involving a function call.
type FunctionCall struct {
	Function Node
	Params   []Node
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*FunctionCall) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (c *FunctionCall) FreeVars() stringset.StringSet {
	freeVars := c.Function.FreeVars()
	sets := make([]stringset.StringSet, len(c.Params))
	for i, node := range c.Params {
		sets[i] = node.FreeVars()
	}
	return freeVars.Union(sets...)
}

// ToExpr converts the node to an expr.
func (c *FunctionCall) ToExpr() types.Expr {
	list := []types.Expr{c.Function.ToExpr()}
	for _, param := range c.Params {
		list = append(list, param.ToExpr())
	}
	return types.NewConsList(list...)
}
