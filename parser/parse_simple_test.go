package parser_test

import "remexre.xyz/parallisp/types"

var simpleTests = []test{
	// Cons
	{"(1 2 3 4 5)", types.NewConsList(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
		types.Integer(4),
		types.Integer(5),
	), "", true},
	// Floating
	{"123.45", types.Floating(123.45), "", true},
	// Integer
	{"123", types.Integer(123), "", true},
	// Nil
	{"()", nil, "", true},
	// String TODO
	// Symbol
	{"cons", types.Symbol("cons"), "", true},
	{"+", types.Symbol("+"), "", true},
	{"-", types.Symbol("-"), "", true},
	{"*", types.Symbol("*"), "", true},
	{"/", types.Symbol("/"), "", true},
	{"a.b", types.Symbol("a.b"), "", true},
	// Vector
	{"[]", types.NewVector(), "", true},
	{"[1 2 3 4 5]", types.NewVector(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
		types.Integer(4),
		types.Integer(5),
	), "", true},
}
