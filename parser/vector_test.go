package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp/parser"
	"remexre.xyz/parallisp/types"
)

var _ = Describe("Vector Parser", func() {
	do(parser.ParseVector, vectorTests)
})

var vectorTests = []test{
	{"[]", types.Vector{}, "", true},
	{"[1]", types.NewVector(
		types.Integer(1),
	), "", true},
	{"[1 2 3]", types.NewVector(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
	), "", true},
}
