package types

import (
	"bytes"
	"fmt"
)

// Vector is a dynamically sized array of expressions.
type Vector []Expr

// NewVector creates a vector from Exprs.
func NewVector(exprs ...Expr) Expr {
	if exprs == nil {
		return Vector{}
	}
	return Vector(exprs)
}

// Eval evaluates an expression.
func (expr Vector) Eval(env Env) (Expr, error) {
	return expr, nil
}

// Get returns the nth element of the vector.
func (expr Vector) Get(n Integer) (Expr, error) {
	if int(n) >= len(expr) || n < 0 {
		return nil, fmt.Errorf("%s does not have an index %d", expr, n)
	}
	return expr[int(n)], nil
}

// Slice slices the vector from from to to.
func (expr Vector) Slice(from, to Integer) (Expr, error) {
	if int(to) >= len(expr) || to < 0 {
		return nil, fmt.Errorf("%s does not have an index %d", expr, to)
	} else if to < from || from < 0 {
		return nil, fmt.Errorf("%d is not a valid start index", from)
	}
	return expr[int(from):int(to)], nil
}

// String converts an expression to a string.
func (expr Vector) String() string {
	buf := new(bytes.Buffer)
	buf.WriteRune('[')
	for i, expr := range expr {
		if i > 0 {
			buf.WriteRune(' ')
		}
		buf.WriteString(ExprToString(expr))
	}
	buf.WriteRune(']')
	return buf.String()
}

// Type converts the type of an expression to a string.
func (Vector) Type() string {
	return "vector"
}

// ToSlice converts a vector into a slice, which it already is. This method
// exists mainly for interface purposes.
func (expr Vector) ToSlice() []Expr {
	return expr
}
