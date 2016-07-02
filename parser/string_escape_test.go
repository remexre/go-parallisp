package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp/parser"
)

var _ = Describe("String Escape Parser", func() {
	doSimple(parser.ParseStringEscape, stringEscapeTests)
})

var stringEscapeTests = []simpleTest{
	{`\`, `\`},
	{`"`, `"`},
	{"n", "\n"},
	{"r", "\r"},
	{"t", "\t"},

	{"x0a", "\n"},
	{"u2713", "✓"},
	{"U0001F604", "😄"},

	{"xc0", "À"},
	{"u00c0", "À"},
	{"U000000c0", "À"},
}
