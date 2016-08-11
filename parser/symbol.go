package parser

import (
	"github.com/remexre/go-parcom"
	"remexre.xyz/go-parallisp/types"
)

// ParseSymbol parses a symbol.
func ParseSymbol(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.OneOfFunc(symbolStartByte),
		parcom.Opt(
			parcom.AnyOfFunc(symbolByte),
			"",
		),
	), func(first, rest string) types.Expr {
		return types.Symbol(first + rest)
	})(in)
}

func symbolStartByte(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/' || b == '%' ||
		b == '!' || b == '=' || b == '<' || b == '>' || b == '@' || b == '&' ||
		b == ':' || b == '?' || ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z')
}

func symbolByte(b byte) bool {
	return symbolStartByte(b) || ('0' <= b && b <= '9') || b == '.'
}
