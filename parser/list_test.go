package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp"
	"remexre.xyz/parallisp/parser"
)

var _ = Describe("List Parser", func() {
	do(parser.ParseList, listTests)
})

var listTests = []test{
	{"()", nil, "", true},
	{"(())", parallisp.Cons{
		nil,
		nil,
	}, "", true},
	{"(1)", parallisp.Cons{
		parallisp.Integer(1),
		nil,
	}, "", true},
	{"(1 2 3)", parallisp.Cons{
		parallisp.Integer(1),
		parallisp.Cons{
			parallisp.Integer(2),
			parallisp.Cons{
				parallisp.Integer(3),
				nil,
			},
		},
	}, "", true},
	{"(1 2 . 3)", parallisp.Cons{
		parallisp.Integer(1),
		parallisp.Cons{
			parallisp.Integer(2),
			parallisp.Integer(3),
		},
	}, "", true},
}
