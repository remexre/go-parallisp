package parser

import (
	"github.com/remexre/go-parcom"
	"github.com/remexre/go-parallisp/parser/number"
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
