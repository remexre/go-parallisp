package ast

// A FunctionCall is the evaluation of an expression involving a function call.
type FunctionCall struct {
	Function Node
	Params   []Node
}

// FreeVars returns the free values contained within a node, recursively.
func (c *FunctionCall) FreeVars() []string {
	freeVars := c.Function.FreeVars()
	for _, param := range c.Params {
		freeVars = append(freeVars, param.FreeVars()...)
	}
	return freeVars
}
