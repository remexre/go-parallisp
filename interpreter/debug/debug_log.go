package debug

import (
	"fmt"

	"github.com/remexre/go-parallisp/debug"
	"github.com/remexre/go-parallisp/types"
)

// Log is the **debug-log** function.
func Log(namespace types.Symbol, args ...types.Expr) {
	debug.Log(string(namespace), fmt.Sprint(args))
}
