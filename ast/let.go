package ast

import "remexre.xyz/go-parallisp/types"

// Let represents a let-expression.
type Let struct {
	Definitions []LetDefinition
	Child       Progn
}

// LetDefinition is a definition in a let-expression.
type LetDefinition struct {
	Name  types.Symbol
	Value types.Expr
}
