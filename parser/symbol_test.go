package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp"
	"remexre.xyz/parallisp/parser"
)

var _ = Describe("Symbol Parser", func() {
	do(parser.ParseSymbol, symbolTests)
})

var symbolTests = []test{
	{"+", parallisp.Symbol("+"), "", true},
	{"cons", parallisp.Symbol("cons"), "", true},
	{"car", parallisp.Symbol("car"), "", true},
	{"cdr", parallisp.Symbol("cdr"), "", true},
}
