package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A Vector is a vector literal.
type Vector []Node

// FreeVars returns the free values contained within a node, recursively.
func (vec *Vector) FreeVars() stringset.StringSet {
	var freeVars stringset.StringSet
	for _, node := range *vec {
		freeVars = freeVars.Union(node.FreeVars())
	}
	return freeVars
}
