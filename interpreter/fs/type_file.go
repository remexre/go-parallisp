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

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (File) LiteralAsm() string {
	panic("interpreter.fs.File.LiteralAsm: cannot make a file into a literal")
}

func (f File) String() string {
	return fmt.Sprintf("<file %s>", f.inner.Name())
}

// TypeAsm converts an Expr to its type code, without panicking on nil.
func (File) TypeAsm() byte {
	panic("interpreter.fs.File.TypeAsm: cannot make a file into a literal")
}

// Type returns the type of the expression.
func (f File) Type() string { return "file" }
