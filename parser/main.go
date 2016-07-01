package parser

import (
	"errors"

	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp"
	"remexre.xyz/parallisp/parser/number"
)

// ParseExpr parses a single parallisp expression, including optional leading or
// trailing whitespace. It just wraps ExprParser, making it easier to use.
func ParseExpr(in string) (string, parallisp.Expr, bool) {
	remaining, out, ok := parcom.Map(parcom.Chain(
		whitespace,
		ExprParser,
		whitespace,
	), func(_ string, expr parallisp.Expr, _ string) parallisp.Expr {
		return expr
	})(in)
	return remaining, out.(parallisp.Expr), ok
}

// ParseAll parses all available parallisp expressions from the string.
func ParseAll(in string) ([]parallisp.Expr, error) {
	remaining, out, ok := parcom.Map(parcom.Many0(parcom.Map(parcom.Chain(
		whitespace,
		ExprParser,
		whitespace,
	), func(_ string, expr parallisp.Expr, _ string) parallisp.Expr {
		return expr
	})), func(exprs []parallisp.Expr) []parallisp.Expr {
		return exprs
	})(in)

	if !ok {
		return nil, errors.New("parallisp.parser: parsing failed")
	} else if len(remaining) > 0 {
		return nil, errors.New("Unexpected input: " + remaining)
	}
	return out.([]parallisp.Expr), nil
}

// ExprParser parses a single parallisp expression.
func ExprParser(in string) (string, interface{}, bool) {
	return parcom.Alt(
		number.Parse,
		ParseList,
		ParseVector,
		// TODO ParseString,
		ParseSymbol,
	)(in)
}
