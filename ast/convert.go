package ast

import (
	"errors"
	"fmt"
	"strings"

	"remexre.xyz/go-parallisp/types"
)

// Convert converts an expression to an AST node.
func Convert(exprIn types.Expr) (Node, error) {
	switch expr := exprIn.(type) {
	case types.Cons:
		if !expr.IsList() {
			return nil, fmt.Errorf("ast.Convert: cannot call improper list %v", expr)
		}
		return ConvertCall(expr.ToSlice())
	case types.Floating:
		return Floating(expr), nil
	case types.Integer:
		return Integer(expr), nil
	case types.String:
		return String(expr), nil
	case types.Symbol:
		v := Variable(expr)
		return &v, nil
	case types.Vector:
		var vec Vector
		for _, expr := range expr {
			node, err := Convert(expr)
			if err != nil {
				return nil, err
			}
			vec = append(vec, node)
		}
		return &vec, nil
	default:
		return nil, fmt.Errorf("ast.Convert: unknown expr %v of type %T", expr, expr)
	}
}

// ConvertCall converts a function call to an AST Node.
func ConvertCall(exprs []types.Expr) (Node, error) {
	switch exprs[0] {
	case types.Symbol("defun"):
		return NewDefun(exprs[1:])
	case types.Symbol("defmacro"):
		return NewDefmacro(exprs[1:])
	case types.Symbol("import"):
		return NewImport(exprs[1:])
	case types.Symbol("lambda"):
		return NewLambda(exprs[1:])
	case types.Symbol("let"):
		return NewLet(exprs[1:])
	case types.Symbol("let*"):
		return NewSequentialLet(exprs[1:])
	case types.Symbol("quote"):
		// TODO Quote
		return nil, errors.New("TODO Quote")
	}

	nodes := make([]Node, len(exprs))
	for i, expr := range exprs {
		var err error
		nodes[i], err = Convert(expr)
		if err != nil {
			return nil, err
		}
	}
	return &FunctionCall{nodes[0], nodes[1:]}, nil
}

// ConvertProgn converts a slice of exprs to a Progn.
func ConvertProgn(exprs []types.Expr) (Progn, error) {
	progn := make([]Node, len(exprs))
	for i, expr := range exprs {
		var err error
		progn[i], err = Convert(expr)
		if err != nil {
			return nil, err
		}
	}
	return progn, nil
}

// ConvertDocProgn converts a slice of exprs to an optional doc-string and a
// Progn.
func ConvertDocProgn(exprs []types.Expr) (string, Progn, error) {
	var docStrs []string
	for len(exprs) > 1 {
		if str, ok := exprs[0].(types.String); ok {
			docStrs = append(docStrs, string(str))
			exprs = exprs[1:]
		} else {
			break
		}
	}
	progn := make([]Node, len(exprs))
	for i, expr := range exprs {
		var err error
		progn[i], err = Convert(expr)
		if err != nil {
			return "", nil, err
		}
	}
	return strings.Join(docStrs, " "), progn, nil
}
