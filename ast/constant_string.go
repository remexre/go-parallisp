package ast

import "remexre.xyz/go-parallisp/types"

// A String is a string constant.
type String types.String

// FreeVars returns the free values contained within a node, recursively.
func (String) FreeVars() []string { return nil }
