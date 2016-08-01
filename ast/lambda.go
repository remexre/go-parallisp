package ast

import "remexre.xyz/go-parallisp/util/stringset"

// Lambda represents a lambda expression which accepts multiple parameters.
type Lambda struct {
	Params []string
	Doc    string
	Body   Progn
}

// FreeVars returns the free values contained within a node, recursively.
func (l *Lambda) FreeVars() stringset.StringSet {
	return l.Body.FreeVars().Difference(stringset.New(l.Params...))
}
