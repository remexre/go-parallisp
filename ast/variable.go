package ast

// A Variable is a variable.
type Variable string

// FreeVars returns the free values contained within a node, recursively.
func (v *Variable) FreeVars() []string {
	return []string{string(*v)}
}
