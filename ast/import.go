package ast

import (
	"path/filepath"
	"strings"

	"remexre.xyz/go-parallisp/util/stringset"
)

// Import represents an import of zero or more symbols from a module, or
// alternatively a "wildcard" import. An Import with a nil Symbols is a wildcard
// import, as opposed to a zero-length Symbols vector.
type Import struct {
	Module   string
	Symbols  []string
	Wildcard bool
}

// FreeVars returns the free values contained within a node, recursively.
func (*Import) FreeVars() stringset.StringSet { return nil }

// Namespace returns the namespace imported variables will be stored in.
func (i *Import) Namespace() string {
	basename := filepath.Base(i.Module)
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}
