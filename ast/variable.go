package ast

// Variable represents a single bound variable.
type Variable struct {
	Name string
	Free bool
}
