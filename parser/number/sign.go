package number

import (
	"github.com/remexre/go-parcom"
	"github.com/remexre/go-parallisp/types"
)

var sign = parcom.Opt(parcom.Map(parcom.Alt(
	parcom.Tag("+"),
	parcom.Tag("-"),
), func(i string) int64 {
	if i == "-" {
		return -1
	}
	return 1
}), int64(1))

func applySign(sign int64, expr types.Expr) types.Expr {
	if !(sign == 1 || sign == -1) {
		return expr
	}
	switch e := expr.(type) {
	case types.Floating:
		return types.Floating(sign) * e
	case types.Integer:
		return types.Integer(sign) * e
	default:
		return expr
	}
}
