package natives

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/remexre/go-parallisp/types"
)

// Gensym generates a symbol.
func Gensym(names ...string) types.Symbol {
	if names == nil {
		names = []string{"gensym"}
	}
	names = append(names, strconv.FormatInt(int64(rand.Int31()), 36))
	return types.Symbol(strings.Join(names, "-"))
}
