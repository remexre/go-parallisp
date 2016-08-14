package parser_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/remexre/go-parallisp/parser"
)

var _ = Describe("Comment Parser", func() {
	do(parser.ParseComment, commentTests)
})

var commentTests = []test{
	{";\n0", "", "0", true},
	{"; asdf\n0", " asdf", "0", true},
}
