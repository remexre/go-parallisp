package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A Integer is an integer constant.
type Integer int64

// FreeVars returns the free values contained within a node, recursively.
func (Integer) FreeVars() stringset.StringSet { return nil }
