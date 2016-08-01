package ast

import "remexre.xyz/go-parallisp/util/stringset"

// Defmacro represents a macro definition.
type Defmacro struct {
	Name   string
	Params []string
	Doc    string
	Body   Progn
}

// FreeVars returns the free values contained within a node, recursively.
func (d *Defmacro) FreeVars() stringset.StringSet {
	return d.Body.FreeVars().Difference(stringset.New(d.Params...))
}
