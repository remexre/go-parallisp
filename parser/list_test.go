package parser_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/remexre/go-parallisp/parser"
	"github.com/remexre/go-parallisp/types"
)

var _ = Describe("List Parser", func() {
	doSimple(parser.ParseList, listTests)
})

var listTests = []simpleTest{
	{"()", nil},
	{"(())", types.NewCons(
		nil,
		nil,
	)},
	{"(1)", types.NewCons(
		types.Integer(1),
		nil,
	)},
	{"(1 2 3)", types.NewCons(
		types.Integer(1),
		types.NewCons(
			types.Integer(2),
			types.NewCons(
				types.Integer(3),
				nil,
			),
		),
	)},
	{"(1 2 . 3)", types.NewCons(
		types.Integer(1),
		types.NewCons(
			types.Integer(2),
			types.Integer(3),
		),
	)},
}
