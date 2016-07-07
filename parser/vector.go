package parser

import (
	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parcom"
)

// ParseVector parses a vector.
func ParseVector(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("["),
		optionalWS,
		parcom.Many0(parcom.Map(parcom.Chain(
			ParseExpr,
			optionalWS,
		), func(expr types.Expr, _ string) types.Expr {
			return expr
		})),
		parcom.Tag("]"),
	), parseVectorHelper)(in)
}

func parseVectorHelper(_, _ string, exprs []types.Expr, _ string) types.Expr {
	return types.Vector(exprs)
}
