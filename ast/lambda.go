package ast

// Lambda represents a lambda expression which accepts multiple parameters.
type Lambda struct {
	Params []string
	Body   Progn
}

// FreeVars returns the free values contained within a node, recursively.
func (l *Lambda) FreeVars() []string {
	var freeVars []string
	for _, sym := range l.Body.FreeVars() {
		free := true
		for _, param := range l.Params {
			if sym == param {
				free = false
				break
			}
		}
		if free {
			freeVars = append(freeVars, sym)
		}
	}
	return freeVars
}
