package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parallisp/parser"
	"remexre.xyz/go-parallisp/types"
)

var _ = Describe("String Parser", func() {
	doSimple(parser.ParseString, stringTests)
})

var stringTests = []simpleTest{
	{`"hello"`, types.String("hello")},
	{`"Hello, world!"`, types.String("Hello, world!")},
	{`"\""`, types.String(`"`)},
	{`"\"Hello, world!\""`, types.String(`"Hello, world!"`)},
}
