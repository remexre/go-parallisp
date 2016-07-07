package parser

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parcom"
)

// Parse parses all available parallisp expressions from the string.
func Parse(in string) ([]types.Expr, error) {
	remaining, out, ok := parcom.Map(parcom.Many0(parcom.Map(parcom.Chain(
		trash,
		ParseExpr,
		trash,
	), func(_ interface{}, expr types.Expr, _ interface{}) types.Expr {
		return expr
	})), func(exprs []types.Expr) []types.Expr {
		return exprs
	})(in)

	if !ok {
		return nil, fmt.Errorf("parallisp.parser: parsing failed")
	} else if len(remaining) > 0 {
		return nil, fmt.Errorf("parallisp.parser: unexpected input at byte %d: %s",
			len(in)-len(remaining), remaining)
	}
	return out.([]types.Expr), nil
}

func trash(in string) (string, interface{}, bool) {
	return parcom.Many0(parcom.Alt(
		whitespace,
		ParseComment,
	))(in)
}
