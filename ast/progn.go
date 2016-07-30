package ast

// A Progn is a sequence of sequentially evaluated expressions.
type Progn []Node

// FreeVars returns the free values contained within a node, recursively.
func (p *Progn) FreeVars() []string {
	set := make(map[string]struct{})
	for _, node := range *p {
		for _, freeVar := range node.FreeVars() {
			set[freeVar] = struct{}{}
		}
	}

	var freeVars []string
	for freeVar := range set {
		freeVars = append(freeVars, freeVar)
	}
	return freeVars
}
