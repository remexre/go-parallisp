package types

import "fmt"

// Symbol is a fixed symbol, which may be hashed for brevity.
type Symbol string

// Eval evaluates an expression.
func (expr Symbol) Eval(env Env) (Expr, error) {
	if expr, ok := env.Get(expr); ok {
		return expr, nil
	}
	return nil, fmt.Errorf("parallisp.types: unknown variable %s", expr)
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (expr Symbol) LiteralAsm() string {
	panic("types.Function.LiteralAsm: TODO literal symbols")
}

// String converts an expression to a string.
func (expr Symbol) String() string {
	return string(expr)
}

// Type converts the type of an expression to a string.
func (Symbol) Type() string {
	return "symbol"
}

// TypeAsm converts an Expr to its type code, without panicking on nil.
func (Symbol) TypeAsm() byte {
	return 1
}
