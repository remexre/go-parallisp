package parser

import (
	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp"
)

// ParseList parses a list.
func ParseList(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("("),
		whitespace,
		parcom.Many0(parcom.Map(parcom.Chain(
			ExprParser,
			whitespace,
		), func(expr parallisp.Expr, ws string) parallisp.Expr {
			return expr
		})),
		parcom.Opt(parcom.Chain(
			parcom.Tag("."),
			whitespace,
			ExprParser,
			whitespace,
		), nil),
		parcom.Tag(")"),
	), parseListHelper)(in)
}

func parseListHelper(open, openWS string, exprs []parallisp.Expr, improper []interface{}, close string) parallisp.Expr {
	if improper != nil {
		exprs = append(exprs, improper[2].(parallisp.Expr))
		return parallisp.NewImproperConsList(exprs...)
	}
	return parallisp.NewConsList(exprs...)
}
