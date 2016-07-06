package parser

import (
	"remexre.xyz/go-parallisp/parser/number"
	"remexre.xyz/go-parcom"
)

// ParseExpr parses a single parallisp expression.
func ParseExpr(in string) (string, interface{}, bool) {
	return parcom.Alt(
		number.Parse,
		ParseList,
		ParseVector,
		ParseString,
		ParseSymbol,
		ParseReaderMacro,
	)(in)
}
