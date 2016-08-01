package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A String is a string constant.
type String string

// FreeVars returns the free values contained within a node, recursively.
func (String) FreeVars() stringset.StringSet { return nil }
