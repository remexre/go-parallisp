package natives

import "github.com/remexre/go-parallisp/types"

// StringToSymbol converts a string into a symbol.
func StringToSymbol(str types.String) types.Symbol {
	return types.Symbol(str)
}
