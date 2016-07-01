package parser

import (
	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp"
)

// ParseVector parses a vector.
func ParseVector(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag("["),
		whitespace,
		parcom.Many0(parcom.Map(parcom.Chain(
			ExprParser,
			whitespace,
		), func(expr parallisp.Expr, ws string) parallisp.Expr {
			return expr
		})),
		parcom.Tag("]"),
	), parseVectorHelper)(in)
}

func parseVectorHelper(open, openWS string, exprs []parallisp.Expr, close string) parallisp.Expr {
	return parallisp.Vector(exprs)
}
