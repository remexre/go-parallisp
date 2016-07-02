package types

// Symbol is a fixed symbol, which may be hashed for brevity.
type Symbol string

// Expr converts an expression to a string.
func (expr Symbol) Expr() string {
	return string(expr)
}
