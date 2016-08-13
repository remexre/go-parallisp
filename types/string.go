package types

import (
	"fmt"
	"strconv"
)

// String is a dynamically sized array of bytes.
type String string

// Eval evaluates an expression.
func (expr String) Eval(env Env) (Expr, error) {
	return expr, nil
}

// Get returns the nth element of the string.
func (expr String) Get(n Integer) (Expr, error) {
	if int(n) >= len(expr) || n < 0 {
		return nil, fmt.Errorf("%s does not have an index %d", expr, n)
	}
	return Integer([]rune(string(expr))[int(n)]), nil
}

// Len returns the length of the string.
func (expr String) Len() (Integer, error) {
	return Integer(len(expr)), nil
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (expr String) LiteralAsm() string {
	return fmt.Sprintf(".quad %d\n.ascii %s",
		len(expr),
		strconv.Quote(string(expr)))
}

// Slice slices the string from from to to.
func (expr String) Slice(from, to Integer) (Expr, error) {
	if int(to) > len(expr) || to < 0 {
		return nil, fmt.Errorf("%s does not have an index %d", expr, to)
	} else if to < from || from < 0 {
		return nil, fmt.Errorf("%d is not a valid start index", from)
	}

	var out []rune
	for _, r := range []rune(string(expr))[int(from):int(to)] {
		out = append(out, r)
	}
	return String(out), nil
}

// String converts an expression to a string.
func (expr String) String() string {
	return strconv.Quote(string(expr))
}

// Type converts the type of an expression to a string.
func (String) Type() string {
	return "string"
}
