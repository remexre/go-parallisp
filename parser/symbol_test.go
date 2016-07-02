package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp/parser"
	"remexre.xyz/parallisp/types"
)

var _ = Describe("Symbol Parser", func() {
	do(parser.ParseSymbol, symbolTests)
})

var symbolTests = []test{
	{"+", types.Symbol("+"), "", true},
	{"cons", types.Symbol("cons"), "", true},
	{"car", types.Symbol("car"), "", true},
	{"cdr", types.Symbol("cdr"), "", true},
}
