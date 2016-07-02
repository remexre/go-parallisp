package ast

// Binding represents a binding in a let- or letrec-expression.
type Binding struct {
	Variable Variable
	Value    Node
}

// Let represents a let-expression.
type Let struct {
	Bindings []Binding
	Body     Node
}

// LetRec represents a letrec-expression.
type LetRec struct {
	Bindings []Binding
	Body     Node
}
