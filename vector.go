package parallisp

import "bytes"

// Vector is a dynamically sized array of expressions.
type Vector []Expr

// NewVector creates a vector from Exprs.
func NewVector(exprs ...Expr) Expr {
	if exprs == nil {
		return Vector{}
	}
	return Vector(exprs)
}

// Expr converts an expression to a string.
func (expr Vector) Expr() string {
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
