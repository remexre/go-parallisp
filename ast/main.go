package ast

import (
	"fmt"

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
		argVector, ok := exprs[2].(types.Vector)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid defun")
		}
		name, ok := exprs[1].(types.Symbol)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid defun")
		}
		defun := &Defun{
			string(name),
			make([]string, len(argVector)),
			nil,
		}
		for i, arg := range argVector {
			param, ok := arg.(types.Symbol)
			if !ok {
				return nil, fmt.Errorf("ast.Convert: invalid defun")
			}
			defun.Params[i] = string(param)
		}
		var err error
		defun.Body, err = ConvertProgn(exprs[3:])
		return defun, err
	case types.Symbol("defmacro"):
		argVector, ok := exprs[2].(types.Vector)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid defmacro")
		}
		name, ok := exprs[1].(types.Symbol)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid defmacro")
		}
		defmacro := &Defmacro{
			string(name),
			make([]string, len(argVector)),
			nil,
		}
		for i, arg := range argVector {
			param, ok := arg.(types.Symbol)
			if !ok {
				return nil, fmt.Errorf("ast.Convert: invalid defmacro")
			}
			defmacro.Params[i] = string(param)
		}
		var err error
		defmacro.Body, err = ConvertProgn(exprs[3:])
		return defmacro, err
	case types.Symbol("lambda"):
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
