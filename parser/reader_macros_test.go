package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parallisp/parser"
	"remexre.xyz/go-parallisp/types"
)

var _ = Describe("Reader Macro Parser", func() {
	doSimple(parser.ParseReaderMacro, readerMacroTests)
})

var readerMacroTests = []simpleTest{
	{"'a", types.NewConsList(
		types.Symbol("quote"),
		types.Symbol("a"),
	)},
	{"`(a b ,c d)", types.NewConsList(
		types.Symbol("quasiquote"),
		types.NewConsList(
			types.Symbol("a"),
			types.Symbol("b"),
			types.NewConsList(
				types.Symbol("unquote"),
				types.Symbol("c"),
			),
			types.Symbol("d"),
		),
	)},
	{"`(e f ,@g h)", types.NewConsList(
		types.Symbol("quasiquote"),
		types.NewConsList(
			types.Symbol("e"),
			types.Symbol("f"),
			types.NewConsList(
				types.Symbol("unquote-splice"),
				types.Symbol("g"),
			),
			types.Symbol("h"),
		),
	)},
}
