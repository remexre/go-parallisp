package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parallisp/parser"
	"remexre.xyz/go-parallisp/types"
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
