package exprset

import (
	"strings"

	"remexre.xyz/go-parallisp/types"
)

// ExprSet is a type for a set of Exprs.
type ExprSet map[types.Expr]struct{}

// New returns a new ExprSet with the given elements.
func New(exprs ...types.Expr) ExprSet {
	if len(exprs) == 0 {
		return make(ExprSet)
	}
	return make(ExprSet, len(exprs)).Add(exprs...)
}

// Copy returns a copy of a ExprSet.
func (es ExprSet) Copy() ExprSet {
	out := make(ExprSet)
	for str := range es {
		out.add(str)
	}
	return out
}

// Contains returns whether ExprSet contains the given types.Expr.
func (es ExprSet) Contains(str types.Expr) bool {
	_, ok := es[str]
	return ok
}

// Equals returns whether a set is equal to another.
func (es ExprSet) Equals(other ExprSet) bool {
	if es.Length() != other.Length() {
		return false
	}
	return es.Difference(other).Length() == 0
}

// Length returns the cardinality of the set.
func (es ExprSet) Length() int {
	return len(es)
}

// String returns the string representation of the set.
func (es ExprSet) String() string {
	exprs := make([]string, 0, len(es))
	for expr := range es {
		exprs = append(exprs, types.ExprToString(expr))
	}

	return "{" + strings.Join(exprs, " ") + "}"
}

// ToSlice returns the elements of the set as a slice.
func (es ExprSet) ToSlice() []types.Expr {
	var out []types.Expr
	for str := range es {
		out = append(out, str)
	}
	return out
}
