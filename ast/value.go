package ast

import "remexre.xyz/parallisp/types"

// Value is an AST node that represents a value.
type Value struct {
	Expr types.Expr
}
