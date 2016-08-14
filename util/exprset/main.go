package exprset

import (
	"strings"

	"github.com/remexre/go-parallisp/types"
)

// ExprSet is a type for a set of Exprs. It is implemented inefficiently due to
// Vector being unhashable.
type ExprSet []types.Expr

// New returns a new ExprSet with the given elements.
func New(exprs ...types.Expr) ExprSet {
	if len(exprs) == 0 {
		return nil
	}
	return ExprSet(nil).Add(exprs...)
}

// Contains returns whether ExprSet contains the given types.Expr.
func (es ExprSet) Contains(needle types.Expr) bool {
	for _, expr := range es {
		if expr == needle {
			return true
		}
	}
	return false
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
	strs := make([]string, 0, len(es))
	for _, expr := range es {
		strs = append(strs, types.ExprToString(expr))
	}

	return "{" + strings.Join(strs, " ") + "}"
}
