package ast

// A Floating is an floating-point constant.
type Floating float64

// FreeVars returns the free values contained within a node, recursively.
func (Floating) FreeVars() []string { return nil }
