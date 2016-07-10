package natives

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Doc returns an expression's documentation.
func Doc(expr types.Expr) types.String {
	if expr == nil {
		return types.String("The nil value.")
	}
	if doccer, ok := expr.(interface {
		Doc() string
	}); ok {
		return types.String(doccer.Doc())
	}
	return types.String(fmt.Sprintf("%v has no documentation.", expr))
}
