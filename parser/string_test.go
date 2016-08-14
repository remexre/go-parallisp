package parser_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/remexre/go-parallisp/parser"
	"github.com/remexre/go-parallisp/types"
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
