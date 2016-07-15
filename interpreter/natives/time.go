package natives

import (
	"time"

	"remexre.xyz/go-parallisp/types"
)

// Timestamp returns the current Unix timestamp.
func Timestamp() types.Integer {
	return types.Integer(time.Now().Unix())
}
