package parser

import (
	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp/types"
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
		),
		whitespace,
		ExprParser,
	), func(m, _ string, expr types.Expr) types.Expr {
		var macro string
		switch m {
		case "'":
			macro = "quote"
		case "`":
			macro = "quasiquote"
		case ",":
			macro = "unquote"
		case ",@":
			macro = "unquote-splice"
		}

		return types.Cons{types.Symbol(macro), expr}
	})(in)
}
