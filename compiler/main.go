package compiler

import (
	"bytes"
	"fmt"

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
	makeHeader(&buf, "INIT FUNCTION")
	buf.WriteString("\n.section .text\n\n.global _init\n_init:\n")
	buf.WriteString("\t# TODO This is for debugging only...\n")
	for i, lit := range literals {
		buf.WriteString("\tmovq $literal_")
		buf.WriteString(fmt.Sprint(i))
		buf.WriteRune('+')
		buf.WriteString(fmt.Sprint(lit.TypeAsm()))
		buf.WriteString(", %rdi\n\tcall _println\n")
	}
	buf.WriteString("\txorq %rax, %rax\n\tret\n")

	// Add user-defined functions.
	buf.WriteRune('\n')
	makeHeader(&buf, "USER-DEFINED FUNCTIONS")
	buf.WriteRune('\n')
	buf.WriteString("# TODO Other functions\n")

	// Add literals.
	buf.WriteRune('\n')
	makeHeader(&buf, "LITERALS")
	buf.WriteString("\n.section .rodata\n")
	for i, lit := range module.Body.Literals() {
		buf.WriteString("\n.align 16\nliteral_")
		buf.WriteString(fmt.Sprint(i))
		buf.WriteString(": # ")
		buf.WriteString(lit.Type())
		buf.WriteRune('\t')
		buf.WriteString(comment(lit.String()))
		buf.WriteRune('\n')
		buf.WriteString(indent(lit.LiteralAsm()))
		buf.WriteRune('\n')
	}
	return buf.String(), nil
}
