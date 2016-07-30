package ast

import "remexre.xyz/go-parallisp/types"

// A Integer is an integer constant.
type Integer types.Integer

// FreeVars returns the free values contained within a node, recursively.
func (Integer) FreeVars() []string { return nil }
