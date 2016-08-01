package analysis

import "remexre.xyz/go-parallisp/ast"

// FreeVars returns a slice of the free variables for the entire parallisp
// module whose AST is passed as input. Imported variables are removed from
// consideration, but standard library functions are not.
func FreeVars(module ast.Module) []string {
	var imported map[string]struct{}
	for _, importNode := range module.Imports {
		if importNode.Wildcard {
			// TODO
		} else {
			for _, sym := range importNode.Symbols {
				imported[importNode.Namespace()+":"+sym] = struct{}{}
			}
		}
	}

	var out []string
	for _, freeVar := range module.Body.FreeVars() {
		if _, ok := imported[freeVar]; ok {
			continue
		}
		out = append(out, freeVar)
	}
	return out
}
