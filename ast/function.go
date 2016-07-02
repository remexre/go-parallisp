package ast

// Function represents a defined function.
type Function struct {
	Name       string
	Parameters []Variable
	Body       Node
}
