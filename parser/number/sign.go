package number

import (
	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp"
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

func applySign(sign int64, expr parallisp.Expr) parallisp.Expr {
	if !(sign == 1 || sign == -1) {
		return expr
	}
	switch e := expr.(type) {
	case parallisp.Floating:
		return parallisp.Floating(sign) * e
	case parallisp.Integer:
		return parallisp.Integer(sign) * e
	default:
		return expr
	}
}
