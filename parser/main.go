package parser

import (
	"fmt"

	"remexre.xyz/go-parcom"
	"remexre.xyz/go-parallisp/types"
)

// Parse parses all available parallisp expressions from the string.
func Parse(in string) ([]types.Expr, error) {
	remaining, out, ok := parcom.Map(parcom.Many0(parcom.Map(parcom.Chain(
		whitespace,
		ParseExpr,
		whitespace,
	), func(_ string, expr types.Expr, _ string) types.Expr {
		return expr
	})), func(exprs []types.Expr) []types.Expr {
		return exprs
	})(in)

	if !ok {
		return nil, fmt.Errorf("parallisp.parser: parsing failed")
	} else if len(remaining) > 0 {
		return nil, fmt.Errorf("parallisp.parser: unexpected input at byte %d", len(remaining))
	}
	return out.([]types.Expr), nil
}
