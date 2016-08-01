package ast

import "remexre.xyz/go-parallisp/util/stringset"

// Node represents a single AST node.
type Node interface {
	// FreeVars returns the free values contained within a node, recursively.
	FreeVars() stringset.StringSet
}
