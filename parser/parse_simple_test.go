package parser_test

import "remexre.xyz/go-parallisp/types"

var simpleTests = []simpleTest{
	// Comments
	{"; a comment\n0", types.Integer(0)},
	{"(1; asdf\n2 3)", types.NewConsList(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
	)},
	// Cons
	{"(1 2 3 4 5)", types.NewConsList(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
		types.Integer(4),
		types.Integer(5),
	)},
	// Floating
	{"123.45", types.Floating(123.45)},
	// Integer
	{"123", types.Integer(123)},
	// Nil
	{"()", nil},
	// String
	{`"hello"`, types.String("hello")},
	// Symbol
	{"cons", types.Symbol("cons")},
	{"+", types.Symbol("+")},
	{"-", types.Symbol("-")},
	{"*", types.Symbol("*")},
	{"/", types.Symbol("/")},
	{"a.b", types.Symbol("a.b")},
	// Vector
	{"[]", types.NewVector()},
	{"[1 2 3 4 5]", types.NewVector(
		types.Integer(1),
		types.Integer(2),
		types.Integer(3),
		types.Integer(4),
		types.Integer(5),
	)},
}
