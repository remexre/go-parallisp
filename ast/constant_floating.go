package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A Floating is an floating-point constant.
type Floating float64

// FreeVars returns the free values contained within a node, recursively.
func (Floating) FreeVars() stringset.StringSet { return nil }
