package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A FunctionCall is the evaluation of an expression involving a function call.
type FunctionCall struct {
	Function Node
	Params   []Node
}

// FreeVars returns the free values contained within a node, recursively.
func (c *FunctionCall) FreeVars() stringset.StringSet {
	freeVars := c.Function.FreeVars()
	sets := make([]stringset.StringSet, len(c.Params))
	for i, node := range c.Params {
		sets[i] = node.FreeVars()
	}
	return freeVars.Union(sets...)
}
