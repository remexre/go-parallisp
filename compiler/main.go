package compiler

import (
	"bytes"
	"fmt"
	"strings"

	"remexre.xyz/go-parallisp/ast"
)

// Compile compiles a module.
func Compile(module *ast.Module) (string, error) {
	fmt.Println("defines", module.Body.Defines())
	fmt.Println("free vars", module.Body.FreeVars())
	fmt.Println("literals", module.Body.Literals())

	var out bytes.Buffer
	out.WriteString("# TODO _init\n")
	out.WriteString("# TODO Other functions\n")

	// Add literals
	for i, lit := range module.Body.Literals() {
		out.WriteString("\n# ")
		out.WriteString(strings.Replace(lit.String(), "\n", "\n# ", -1))
		out.WriteString("\nliteral_")
		out.WriteString(fmt.Sprint(i))
		out.WriteString(":\n")
		out.WriteString(lit.LiteralAsm())
		out.WriteRune('\n')
	}
	return out.String(), nil
}
