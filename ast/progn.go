package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A Progn is a sequence of sequentially evaluated expressions.
type Progn []Node

// FreeVars returns the free values contained within a node, recursively.
func (p *Progn) FreeVars() stringset.StringSet {
	sets := make([]stringset.StringSet, len(*p))
	for i, node := range *p {
		sets[i] = node.FreeVars()
	}
	return stringset.Union(sets...)
}
