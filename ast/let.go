package ast

import "remexre.xyz/go-parallisp/util/stringset"

// Let represents a let-expression.
type Let struct {
	Definitions []LetDefinition
	Child       Progn
	Sequential  bool
}

// LetDefinition is one let definition.
type LetDefinition struct {
	Name string
	Node Node
}

// FreeVars returns the free values contained within a node, recursively.
func (l *Let) FreeVars() stringset.StringSet {
	freeVars := l.Child.FreeVars()
	if l.Sequential {
		for i := len(l.Definitions) - 1; i >= 0; i-- {
			def := l.Definitions[i]
			freeVars = freeVars.
				Difference(stringset.New(def.Name)).
				Union(def.Node.FreeVars())
		}
	} else {
		length := len(l.Definitions)
		names := make([]string, 0, length)
		nodes := make([]Node, 0, length)
		for _, def := range l.Definitions {
			names = append(names, def.Name)
			nodes = append(nodes, def.Node)
		}
		nodeFreeVars := make([]stringset.StringSet, length)
		for i, node := range nodes {
			nodeFreeVars[i] = node.FreeVars()
		}
		freeVars = freeVars.
			Difference(stringset.New(names...)).
			Union(nodeFreeVars...)
	}
	return freeVars
}
