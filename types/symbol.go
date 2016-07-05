package types

// Symbol is a fixed symbol, which may be hashed for brevity.
type Symbol string

// String converts an expression to a string.
func (expr Symbol) String() string {
	return string(expr)
}

// Type converts the type of an expression to a string.
func (Symbol) Type() string {
	return "symbol"
}
