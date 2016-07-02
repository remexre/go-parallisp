package ast

// Node is an interface for an AST node.
type Node interface {
	nodeMarkerFunc()
}

func (*Function) nodeMarkerFunc()     {}
func (*FunctionCall) nodeMarkerFunc() {}
func (ProgN) nodeMarkerFunc()         {}
func (*Value) nodeMarkerFunc()        {}
func (*Variable) nodeMarkerFunc()     {}
