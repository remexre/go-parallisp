package fs

import (
	"io/ioutil"

	"github.com/remexre/go-parallisp/types"
)

// ReadAll reads a file's contents as a string.
func ReadAll(file File) (types.Expr, error) {
	b, err := ioutil.ReadAll(file.inner)
	if err != nil {
		return nil, err
	}
	return types.String(b), nil
}
