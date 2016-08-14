package exprset

import "github.com/remexre/go-parallisp/types"

// Add adds a expr or exprs to a ExprSet, then returns the ExprSet.
func (es ExprSet) Add(exprs ...types.Expr) ExprSet {
	return append(es, exprs...)
}
