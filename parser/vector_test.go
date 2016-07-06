package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parallisp/parser"
	"remexre.xyz/go-parallisp/types"
)

var _ = Describe("Vector Parser", func() {
	doSimple(parser.ParseVector, vectorTests)
})

var vectorTests = []simpleTest{
	{"[]", types.Vector{}},
	{"[1]", types.NewVector(
		types.Integer(1),
	)},
	{"[1 2 3]", types.NewVector(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
	)},
}
