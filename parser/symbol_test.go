package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp/parser"
	"remexre.xyz/parallisp/types"
)

var _ = Describe("Symbol Parser", func() {
	doSimple(parser.ParseSymbol, symbolTests)
})

var symbolTests = []simpleTest{
	{"+", types.Symbol("+")},
	{"cons", types.Symbol("cons")},
	{"car", types.Symbol("car")},
	{"cdr", types.Symbol("cdr")},
}
