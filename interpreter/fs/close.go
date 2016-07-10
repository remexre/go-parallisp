package fs

import "remexre.xyz/go-parallisp/types"

// Close closes a file.
func Close(file File) (types.Expr, error) {
	return nil, file.inner.Close()
}
