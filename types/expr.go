package types

// Expr represents an expression.
type Expr interface {
	String() string
	Type() string
}

// ExprToString converts an Expr to its string representation, without panicking
// on nil.
func ExprToString(expr Expr) string {
	if expr == nil {
		return "()"
	}
	return expr.String()
}

// ExprToType converts an Expr's type to its string representation, without
// panicking on nil.
func ExprToType(expr Expr) string {
	if expr == nil {
		return "nil"
	}
	return expr.Type()
}
