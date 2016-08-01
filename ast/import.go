package ast

// Import represents an import of zero or more symbols from a module, or
// alternatively a "wildcard" import. An Import with a nil Symbols is a wildcard
// import, as opposed to a zero-length Symbols vector.
type Import struct {
	Module   string
	Symbols  []string
	Wildcard bool
}

// FreeVars returns the free values contained within a node, recursively.
func (*Import) FreeVars() []string { return nil }
