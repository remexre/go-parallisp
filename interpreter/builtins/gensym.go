package builtins

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"remexre.xyz/go-parallisp/types"
)

// Gensym generates a symbol.
func Gensym(name ...string) (types.Symbol, error) {
	prefix := strings.Join(name, "")
	if prefix == "" {
		prefix = "gensym"
	}
	num := strconv.FormatInt(int64(rand.Int31()), 36)
	return types.Symbol(fmt.Sprintf("%s-%s", prefix, num)), nil
}
