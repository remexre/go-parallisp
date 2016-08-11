package compiler

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/ast"
)

// Compile compiles a module.
func Compile(module *ast.Module) (string, error) {
	fmt.Println("constants", module.Body.Constants())
	fmt.Println("defines", module.Body.Defines())
	fmt.Println("free vars", module.Body.FreeVars())
	return "", errors.New("TODO Actual Compilation")
}
