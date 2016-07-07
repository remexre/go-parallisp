package parser

import (
	"bytes"
	"strconv"

	"github.com/remexre/go-parcom"
)

// ParseStringEscape parses a string escape, without the leading backslash.
func ParseStringEscape(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Alt(
		parcom.Tag(`\`),
		parcom.Tag(`"`),
		parcom.Tag("n"),
		parcom.Tag("r"),
		parcom.Tag("t"),
		parcom.Map(parcom.Alt(
			parcom.Chain(
				parcom.Tag("x"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
			),
			parcom.Chain(
				parcom.Tag("u"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
			),
			parcom.Chain(
				parcom.Tag("U"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
				parcom.OneOf("0123456789abcdefABCDEF"),
			),
		), func(strs []string) string {
			buf := new(bytes.Buffer)
			for _, str := range strs {
				buf.WriteString(str)
			}
			return buf.String()
		}),
	), func(sequence string) string {
		switch sequence[0] {
		case '\\', '"':
			return sequence
		case 'n':
			return "\n"
		case 'r':
			return "\r"
		case 't':
			return "\t"
		case 'x', 'u', 'U':
			codepoint, err := strconv.ParseInt(sequence[1:], 16, 32)
			if err != nil {
				panic(err)
			}
			return string(rune(codepoint))
		default:
			panic("Unknown escape " + sequence)
		}
	})(in)
}
