package parser

import (
	"github.com/remexre/go-parcom"
	"remexre.xyz/go-parallisp/types"
)

// ParseReaderMacro parses a quoted, quasiquoted, unquoted, or unquote-spliced
// expression.
func ParseReaderMacro(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Alt(
			parcom.Tag("'"),
			parcom.Tag("`"),
			parcom.Tag(",@"),
			parcom.Tag(","),
			parcom.Tag("#"),
		),
		trash,
		ParseExpr,
	), func(m string, _ interface{}, expr types.Expr) types.Expr {
		macro := "unknown-reader-macro"
		switch m {
		case "'":
			macro = "quote"
		case "`":
			macro = "quasiquote"
		case ",":
			macro = "unquote"
		case ",@":
			macro = "unquote-splice"
		case "#":
			macro = "basic-lambda"
		}

		return types.Cons{types.Symbol(macro), types.Cons{expr, nil}}
	})(in)
}
