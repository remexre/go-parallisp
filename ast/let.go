package ast

// Let represents a let-expression.
type Let struct {
	Definitions []LetDefinition
	Child       Progn
	Sequential  bool
}

// FreeVars returns the free values contained within a node, recursively.
func (l *Let) FreeVars() []string {
	// TODO Add/remove definitions' freevars.
	return l.Child.FreeVars()
}

// LetDefinition is a definition in a let-expression.
type LetDefinition struct {
	Name  string
	Value Node
}
