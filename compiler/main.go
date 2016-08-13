package compiler

import (
	"bytes"
	"fmt"
	"strings"

	"remexre.xyz/go-parallisp/ast"
)

// Compile compiles a module.
func Compile(module *ast.Module) (string, error) {
	var buf bytes.Buffer
	defines := module.Body.Defines()
	freeVars := module.Body.FreeVars()
	literals := module.Body.Literals()

	// For debugging purposes.
	fmt.Println("defines", defines)
	fmt.Println("freeVars", freeVars)
	fmt.Println("literals", literals)

	// Add _init function.
	buf.WriteString("_init: # TODO This is for debugging only...\n")
	buf.WriteString("\tmovq $literal_0+")
	buf.WriteString(fmt.Sprint(literals[0].TypeAsm()))
	buf.WriteString(", %rax\n")
	buf.WriteString("\tret\n")

	// Add user-created functions.
	buf.WriteString("# TODO Other functions\n")

	// Add literals.
	for i, lit := range module.Body.Literals() {
		buf.WriteString("\n# ")
		buf.WriteString(lit.Type())
		buf.WriteRune('\t')
		buf.WriteString(strings.Replace(lit.String(), "\n", "\n# ", -1))
		buf.WriteString("\nliteral_")
		buf.WriteString(fmt.Sprint(i))
		buf.WriteString(":\n")
		buf.WriteString(lit.LiteralAsm())
		buf.WriteRune('\n')
	}
	return buf.String(), nil
}
