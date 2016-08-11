package exprset

import "remexre.xyz/go-parallisp/types"

// Add adds a expr or exprs to a ExprSet, then returns the ExprSet.
func (es ExprSet) Add(exprs ...types.Expr) ExprSet {
	for _, str := range exprs {
		es.add(str)
	}
	return es
}

func (es ExprSet) add(str types.Expr) {
	es[str] = struct{}{}
}
