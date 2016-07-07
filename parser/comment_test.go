package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parallisp/parser"
)

var _ = Describe("Comment Parser", func() {
	do(parser.ParseComment, commentTests)
})

var commentTests = []test{
	{";\n0", "", "0", true},
	{"; asdf\n0", " asdf", "0", true},
}
