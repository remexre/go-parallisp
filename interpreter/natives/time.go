package natives

import (
	"time"

	"github.com/remexre/go-parallisp/types"
)

// Timestamp returns the current Unix timestamp.
func Timestamp() types.Integer {
	return types.Integer(time.Now().Unix())
}
