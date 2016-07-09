package builtins

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Plus adds numbers.
func Plus(args ...types.Expr) (types.Expr, error) {
	if args, ok := allIntegers(args); ok {
		var i types.Integer
		for _, arg := range args {
			i += arg
		}
		return i, nil
	}
	if args, ok := allNumbers(args); ok {
		var i types.Floating
		for _, arg := range args {
			i += arg
		}
		return i, nil
	}
	return nil, fmt.Errorf("+: cannot add non-numbers %s", args)
}
