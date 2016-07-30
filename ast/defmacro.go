package ast

// Defmacro represents a macro definition.
type Defmacro struct {
	Name   string
	Params []string
	Body   Progn
}

// FreeVars returns the free values contained within a node, recursively.
func (d *Defmacro) FreeVars() []string {
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
