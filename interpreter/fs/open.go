package fs

import (
	"os"

	"github.com/remexre/go-parallisp/types"
)

// Open opens a file.
func Open(path types.String) (types.Expr, error) {
	file, err := os.Open(string(path))
	if err != nil {
		return nil, err
	}
	return File{file}, nil
}
