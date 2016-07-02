package ast

// FunctionCall represents a call to a function.
type FunctionCall struct {
	Function   Function
	Parameters []Node
}
