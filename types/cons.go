package types

import (
	"bytes"
	"io"
)

// Cons is a pair of pointers, possibly to nil.
type Cons [2]Expr

// NewCons is a helper to avoid https://github.com/golang/go/issues/9171.
func NewCons(car, cdr Expr) Cons {
	return Cons{car, cdr}
}

// NewConsList creates a list of cons cells from Exprs.
func NewConsList(exprs ...Expr) Expr {
	var out Expr
	for i := len(exprs) - 1; i >= 0; i-- {
		out = Cons{exprs[i], out}
	}
	return out
}

// NewImproperConsList creates an improper list of cons cells from Exprs.
func NewImproperConsList(exprs ...Expr) Expr {
	out := exprs[len(exprs)-1]
	for i := len(exprs) - 2; i >= 0; i-- {
		out = Cons{exprs[i], out}
	}
	return out
}

// Expr converts an expression to a string.
func (expr Cons) Expr() string {
	out := bytes.NewBufferString("(")
	expr.exprNoParen(out)
	out.WriteRune(')')
	return out.String()
}

func (expr Cons) exprNoParen(w io.Writer) {
	io.WriteString(w, ExprToString(expr[0]))
	if expr[1] == nil {
		return
	}
	io.WriteString(w, " ")
	if next, ok := expr[1].(Cons); ok {
		next.exprNoParen(w)
	} else {
		io.WriteString(w, ExprToString(expr[1]))
	}
}
