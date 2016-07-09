package builtins

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Println prints expressions.
func Println(exprs ...types.Expr) {
	args := make([]interface{}, len(exprs))
	for i, expr := range exprs {
		if str, ok := expr.(types.String); ok {
			args[i] = string(str)
		} else {
			args[i] = expr
		}
	}
	fmt.Println(args...)
}
