package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp/parser"
	"remexre.xyz/parallisp/types"
)

var _ = Describe("List Parser", func() {
	do(parser.ParseList, listTests)
})

var listTests = []test{
	{"()", nil, "", true},
	{"(())", types.NewCons(
		nil,
		nil,
	), "", true},
	{"(1)", types.NewCons(
		types.Integer(1),
		nil,
	), "", true},
	{"(1 2 3)", types.NewCons(
		types.Integer(1),
		types.NewCons(
			types.Integer(2),
			types.NewCons(
				types.Integer(3),
				nil,
			),
		),
	), "", true},
	{"(1 2 . 3)", types.NewCons(
		types.Integer(1),
		types.NewCons(
			types.Integer(2),
			types.Integer(3),
		),
	), "", true},
}
