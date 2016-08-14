package fs

import (
	"fmt"
	"os"

	"github.com/remexre/go-parallisp/types"
)

// File is a wrapper around os.File.
type File struct {
	inner *os.File
}

// Eval evaluates an expression.
func (f File) Eval(env types.Env) (types.Expr, error) {
	return nil, fmt.Errorf("cannot evaluate a file")
}

func (f File) String() string {
	return fmt.Sprintf("<file %s>", f.inner.Name())
}

// Type returns the type of the expression.
func (f File) Type() string { return "file" }
