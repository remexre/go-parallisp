package ast

// Module represents a parallisp module.
type Module struct {
	Imports []Import
	Body    Progn
}
