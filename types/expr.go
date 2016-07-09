package types

import "remexre.xyz/go-parallisp/debug"

// Expr represents an expression.
type Expr interface {
	Eval(env Env) (Expr, error)
	String() string
	Type() string
}

// EvalExpr evaluates an expression, without panicking on nil.
func EvalExpr(env Env, expr Expr) (Expr, error) {
	debug.Log("eval-expr", "evaluating %v", expr)
	if expr == nil {
		return nil, nil
	}
	return expr.Eval(env)
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
