package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parallisp/parser"
)

var _ = Describe("ParseExpr", func() {
	Describe("Simple tests", func() {
		doSimple(parser.ParseExpr, simpleTests)
	})
})
