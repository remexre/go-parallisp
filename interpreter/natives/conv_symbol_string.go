package natives

import "github.com/remexre/go-parallisp/types"

// SymbolToString converts a symbol into a string.
func SymbolToString(sym types.Symbol) types.String {
	return types.String(sym)
}
