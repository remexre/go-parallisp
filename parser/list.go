package parser

import (
	"github.com/remexre/go-parcom"
	"remexre.xyz/go-parallisp/types"
)

// ParseList parses a list.
func ParseList(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("("),
		trash,
		parcom.Many0(parcom.Map(parcom.Chain(
			ParseExpr,
			trash,
		), func(expr types.Expr, _ interface{}) types.Expr {
			return expr
		})),
		parcom.Opt(parcom.Chain(
			parcom.Tag("."),
			trash,
			ParseExpr,
			trash,
		), nil),
		parcom.Tag(")"),
	), func(_ string, _ interface{}, exprs []types.Expr, improper []interface{}, _ string) types.Expr {
		if improper != nil {
			exprs = append(exprs, improper[2].(types.Expr))
			return types.NewImproperConsList(exprs...)
		}
		return types.NewConsList(exprs...)
	})(in)
}
