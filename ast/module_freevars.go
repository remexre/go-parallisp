package ast

import (
	"fmt"

	"remexre.xyz/go-parallisp/util/stringset"
)

// FreeVars returns a slice of the free variables for the entire parallisp
// module whose AST is passed as input. Imported variables are removed from
// consideration, but standard library functions are not.
func (module *Module) FreeVars() stringset.StringSet {
	imported := stringset.New()
	for _, importNode := range module.Imports {
		if importNode.Wildcard {
			// TODO
			fmt.Println("WARN: Wildcard imports not yet supported")
		} else {
			for _, sym := range importNode.Symbols {
				imported.Add(importNode.Namespace() + ":" + sym)
			}
		}
	}

	return module.Body.FreeVars().Difference(imported)
}
