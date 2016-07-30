package ast

// A Vector is a vector literal.
type Vector []Node

// FreeVars returns the free values contained within a node, recursively.
func (vec *Vector) FreeVars() []string {
	var freeVars []string
	for _, node := range *vec {
		freeVars = append(freeVars, node.FreeVars()...)
	}
	return freeVars
}
