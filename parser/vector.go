package parser

import (
	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp/types"
)

// ParseVector parses a vector.
func ParseVector(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("["),
		whitespace,
		parcom.Many0(parcom.Map(parcom.Chain(
			ExprParser,
			whitespace,
		), func(expr types.Expr, ws string) types.Expr {
			return expr
		})),
		parcom.Tag("]"),
	), parseVectorHelper)(in)
}

func parseVectorHelper(open, openWS string, exprs []types.Expr, close string) types.Expr {
	return types.Vector(exprs)
}
