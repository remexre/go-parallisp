package natives

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Print prints a string to stdout.
func Print(str types.String) {
	fmt.Print(string(str))
}
