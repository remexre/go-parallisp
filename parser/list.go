package parser

import (
	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp/types"
)

// ParseList parses a list.
func ParseList(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("("),
		whitespace,
		parcom.Many0(parcom.Map(parcom.Chain(
			ExprParser,
			whitespace,
		), func(expr types.Expr, ws string) types.Expr {
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

func parseListHelper(open, openWS string, exprs []types.Expr, improper []interface{}, close string) types.Expr {
	if improper != nil {
		exprs = append(exprs, improper[2].(types.Expr))
		return types.NewImproperConsList(exprs...)
	}
	return types.NewConsList(exprs...)
}
