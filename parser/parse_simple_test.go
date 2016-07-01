package parser_test

import (
	"remexre.xyz/parallisp"
)

var simpleTests = []test{
	// Cons
	{"(1 2 3 4 5)", parallisp.NewConsList(
		parallisp.Integer(1),
		parallisp.Integer(2),
		parallisp.Integer(3),
		parallisp.Integer(4),
		parallisp.Integer(5),
	), "", true},
	// Floating
	{"123.45", parallisp.Floating(123.45), "", true},
	// Integer
	{"123", parallisp.Integer(123), "", true},
	// Nil
	{"()", nil, "", true},
	// String TODO
	// Symbol
	{"cons", parallisp.Symbol("cons"), "", true},
	{"+", parallisp.Symbol("+"), "", true},
	{"-", parallisp.Symbol("-"), "", true},
	{"*", parallisp.Symbol("*"), "", true},
	{"/", parallisp.Symbol("/"), "", true},
	{"a.b", parallisp.Symbol("a.b"), "", true},
	// Vector
	{"[]", parallisp.NewVector(), "", true},
	{"[1 2 3 4 5]", parallisp.NewVector(
		parallisp.Integer(1),
		parallisp.Integer(2),
		parallisp.Integer(3),
		parallisp.Integer(4),
		parallisp.Integer(5),
	), "", true},
}
