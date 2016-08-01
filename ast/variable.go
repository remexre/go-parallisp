package ast

import "remexre.xyz/go-parallisp/util/stringset"

// A Variable is a variable.
type Variable string

// FreeVars returns the free values contained within a node, recursively.
func (v *Variable) FreeVars() stringset.StringSet {
	return stringset.New(string(*v))
}
