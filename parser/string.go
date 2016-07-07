package parser

import (
	"bytes"

	"github.com/remexre/go-parcom"

	"remexre.xyz/go-parallisp/types"
)

// ParseString parses a string.
func ParseString(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag(`"`),
		parcom.Many0(parcom.Alt(
			parcom.Map(parcom.Chain(
				parcom.Tag(`\`),
				ParseStringEscape,
			), func(_, s string) string { return s }),
			parcom.OneOfFunc(func(b byte) bool { return b != '"' }),
		)),
		parcom.Tag(`"`),
	), func(_ string, parts []string, _ string) types.Expr {
		buf := new(bytes.Buffer)
		for _, part := range parts {
			buf.WriteString(part)
		}
		return types.String(buf.String())
	})(in)
}
