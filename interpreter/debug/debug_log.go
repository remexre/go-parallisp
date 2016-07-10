package debug

import (
	"fmt"

	"remexre.xyz/go-parallisp/debug"
	"remexre.xyz/go-parallisp/types"
)

// Log is the **debug-log** function.
func Log(namespace types.Symbol, args ...types.Expr) {
	debug.Log(string(namespace), fmt.Sprint(args))
}
