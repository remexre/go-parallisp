package parser

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parcom"
)

// ParseList parses a list.
func ParseList(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("("),
		optionalWS,
		parcom.Many0(parcom.Map(parcom.Chain(
			ParseExpr,
			optionalWS,
		), func(expr types.Expr, ws string) types.Expr {
			return expr
		})),
		parcom.Opt(parcom.Chain(
			parcom.Tag("."),
			optionalWS,
			ParseExpr,
			optionalWS,
		), nil),
		parcom.Tag(")"),
	), func(_, _ string, exprs []types.Expr, improper []interface{}, _ string) types.Expr {
		if improper != nil {
			exprs = append(exprs, improper[2].(types.Expr))
			return types.NewImproperConsList(exprs...)
		}
		return types.NewConsList(exprs...)
	})(in)
}
