package ast

import "remexre.xyz/go-parallisp/util/stringset"

// Defun represents a function definition.
type Defun struct {
	Name   string
	Params []string
	Doc    string
	Body   Progn
}

// FreeVars returns the free values contained within a node, recursively.
func (d *Defun) FreeVars() stringset.StringSet {
	return d.Body.FreeVars().Difference(stringset.New(d.Params...))
}
