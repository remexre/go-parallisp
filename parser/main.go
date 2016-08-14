package parser

import (
	"fmt"

	"github.com/remexre/go-parcom"

	"github.com/remexre/go-parallisp/types"
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
