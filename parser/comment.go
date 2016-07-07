package parser

import "remexre.xyz/go-parcom"

// ParseComment parses a comment.
func ParseComment(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		parcom.Tag(";"),
		parcom.Opt(parcom.AnyOfFunc(func(b byte) bool { return b != '\n' }), ""),
		parcom.Tag("\n"),
	), func(_, text, _ string) string {
		return text
	})(in)
}
