package types

import (
	"fmt"

	"github.com/remexre/go-parallisp/debug"
)

// Expr represents an expression.
type Expr interface {
	String() string
	Type() string

	// Interpreter support
	Eval(env Env) (Expr, error)

	// Compiler support
	LiteralAsm() string
	TypeAsm() byte
}

// EvalExpr evaluates an expression, without panicking on nil.
func EvalExpr(env Env, expr Expr) (Expr, error) {
	debug.Log("eval-expr", "evaluating %v", expr)
	if expr == nil {
		return nil, nil
	}
	// return expr.Eval(env)
	out, err := expr.Eval(env)
	if err != nil {
		return nil, fmt.Errorf("%v\n\t%s", expr, err.Error())
	}
	return out, nil
}

// ExprToLiteralAsm converts an Expr to its representation in AT&T syntax x86-64
// assembly, without panicking on nil.
func ExprToLiteralAsm(expr Expr) string {
	if expr == nil {
		return ".quad 0"
	}
	return expr.LiteralAsm()
}

// ExprToTypeAsm converts an Expr to its type code, without panicking on nil.
func ExprToTypeAsm(expr Expr) byte {
	if expr == nil {
		return 0
	}
	return expr.TypeAsm()
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
