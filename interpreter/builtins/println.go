package builtins

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Println prints expressions.
func Println(exprs ...types.Expr) {
	args := make([]interface{}, len(exprs))
	for i, expr := range exprs {
		args[i] = expr
	}
	fmt.Println(args...)
}
