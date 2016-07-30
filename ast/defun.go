package ast

// Defun represents a function definition.
type Defun struct {
	Name   string
	Params []string
	Body   Progn
}

// FreeVars returns the free values contained within a node, recursively.
func (d *Defun) FreeVars() []string {
	var freeVars []string
	for _, sym := range d.Body.FreeVars() {
		free := true
		for _, param := range d.Params {
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
