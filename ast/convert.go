package ast

import (
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
			"",
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
		defun.Doc, defun.Body, err = ConvertDocProgn(exprs[3:])
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
			"",
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
		defmacro.Doc, defmacro.Body, err = ConvertDocProgn(exprs[3:])
		return defmacro, err
	case types.Symbol("import"):
		module, ok := exprs[1].(types.String)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid import")
		}
		importNode := &Import{
			Module: string(module),
		}
		if exprs[2] == types.Symbol("*") {
			importNode.Wildcard = true
			return importNode, nil
		} else if syms, ok := exprs[2].(types.Vector); ok {
			for _, symExpr := range syms {
				sym, ok := symExpr.(types.Symbol)
				if !ok {
					return nil, fmt.Errorf("ast.Convert: invalid import")
				}
				importNode.Symbols = append(importNode.Symbols, string(sym))
			}
			return importNode, nil
		}
		return nil, fmt.Errorf("ast.Convert: invalid import")
	case types.Symbol("lambda"):
		argVector, ok := exprs[1].(types.Vector)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid defun")
		}
		lambda := &Lambda{
			make([]string, len(argVector)),
			"",
			nil,
		}
		for i, arg := range argVector {
			param, ok := arg.(types.Symbol)
			if !ok {
				return nil, fmt.Errorf("ast.Convert: invalid lambda")
			}
			lambda.Params[i] = string(param)
		}
		var err error
		lambda.Doc, lambda.Body, err = ConvertDocProgn(exprs[2:])
		return lambda, err
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
