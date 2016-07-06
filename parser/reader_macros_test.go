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
	{"'a", types.NewCons(
		types.Symbol("quote"),
		types.Symbol("a"),
	)},
	{"`(a b ,c d)", types.NewCons(
		types.Symbol("quasiquote"),
		types.NewConsList(
			types.Symbol("a"),
			types.Symbol("b"),
			types.NewCons(
				types.Symbol("unquote"),
				types.Symbol("c"),
			),
			types.Symbol("d"),
		),
	)},
	{"`(e f ,@g h)", types.NewCons(
		types.Symbol("quasiquote"),
		types.NewConsList(
			types.Symbol("e"),
			types.Symbol("f"),
			types.NewCons(
				types.Symbol("unquote-splice"),
				types.Symbol("g"),
			),
			types.Symbol("h"),
		),
	)},
}
