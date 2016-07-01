package parser_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp"
	"remexre.xyz/parallisp/parser"
)

var _ = Describe("Vector Parser", func() {
	do(parser.ParseVector, vectorTests)
})

var vectorTests = []test{
	{"[]", parallisp.Vector{}, "", true},
	{"[1]", parallisp.NewVector(
		parallisp.Integer(1),
	), "", true},
	{"[1 2 3]", parallisp.NewVector(
		parallisp.Integer(1),
		parallisp.Integer(2),
		parallisp.Integer(3),
	), "", true},
}
