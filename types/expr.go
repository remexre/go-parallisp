package types

// Expr represents an expression.
type Expr interface {
	Expr() string
}

// ExprToString converts an Expr to its string representation, without panicking
// on nil.
func ExprToString(expr Expr) string {
	if expr == nil {
		return "()"
	}
	return expr.Expr()
}
