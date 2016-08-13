package compiler

import (
	"bytes"
	"fmt"

	"remexre.xyz/go-parallisp/ast"
)

// Compile compiles a module.
func Compile(module *ast.Module) (string, error) {
	fmt.Println("defines", module.Body.Defines())
	fmt.Println("free vars", module.Body.FreeVars())
	fmt.Println("literals", module.Body.Literals())

	var out bytes.Buffer
	out.WriteString("// TODO _init\n")
	out.WriteString("// TODO Other functions\n")

	// Add literals
	for i, lit := range module.Body.Literals() {
		fmt.Fprintf(out, "literal_%d:\n", i)
		out.WriteString(lit.LiteralAsm())
	}
	return out, nil
}
