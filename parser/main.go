package parser

import (
	"errors"

	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp/parser/number"
	"remexre.xyz/parallisp/types"
)

// ParseExpr parses a single parallisp expression, including optional leading or
// trailing whitespace. It just wraps ExprParser, making it easier to use.
func ParseExpr(in string) (string, types.Expr, bool) {
	remaining, out, ok := parcom.Map(parcom.Chain(
		whitespace,
		ExprParser,
		whitespace,
	), func(_ string, expr types.Expr, _ string) types.Expr {
		return expr
	})(in)
	return remaining, out.(types.Expr), ok
}

// ParseAll parses all available parallisp expressions from the string.
func ParseAll(in string) ([]types.Expr, error) {
	remaining, out, ok := parcom.Map(parcom.Many0(parcom.Map(parcom.Chain(
		whitespace,
		ExprParser,
		whitespace,
	), func(_ string, expr types.Expr, _ string) types.Expr {
		return expr
	})), func(exprs []types.Expr) []types.Expr {
		return exprs
	})(in)

	if !ok {
		return nil, errors.New("parallisp.parser: parsing failed")
	} else if len(remaining) > 0 {
		return nil, errors.New("Unexpected input: " + remaining)
	}
	return out.([]types.Expr), nil
}

// ExprParser parses a single parallisp expression.
func ExprParser(in string) (string, interface{}, bool) {
	return parcom.Alt(
		number.Parse,
		ParseList,
		ParseVector,
		ParseString,
		ParseSymbol,
		ParseReaderMacro,
	)(in)
}
