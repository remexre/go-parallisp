package ast

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// Defun represents a function definition.
type Defun struct {
	Name   string
	Params []string
	Doc    string
	Body   Progn
}

// NewDefun returns a new defun from the expressions in its form, excluding the
// initial defun symbol.
func NewDefun(exprs []types.Expr) (Node, error) {
	argVector, ok := exprs[1].(types.Vector)
	if !ok {
		return nil, fmt.Errorf("ast.Convert: invalid defun")
	}
	name, ok := exprs[0].(types.Symbol)
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
	defun.Doc, defun.Body, err = ConvertDocProgn(exprs[2:])
	return defun, err
}

// Literals returns the constants used in this node and all child nodes.
func (d *Defun) Literals() exprset.ExprSet {
	return d.Body.Literals().Add(types.String(d.Doc))
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (d *Defun) Defines() stringset.StringSet {
	return stringset.New(d.Name)
}

// FreeVars returns the free values contained within a node, recursively.
func (d *Defun) FreeVars() stringset.StringSet {
	return d.Body.FreeVars().Difference(stringset.New(d.Params...))
}

// ToExpr converts the node to an expr.
func (d *Defun) ToExpr() types.Expr {
	// Build the parameter list.
	params := make([]types.Expr, len(d.Params))
	for i, param := range d.Params {
		params[i] = types.Symbol(param)
	}

	// Build the body.
	body := make([]types.Expr, len(d.Body))
	for i, node := range d.Body {
		body[i] = node.ToExpr()
	}

	// Build the defmacro and return.
	return types.NewConsList(append([]types.Expr{
		types.Symbol("defun"),
		types.Symbol(d.Name),
		types.Vector(params),
		types.String(d.Doc),
	}, body...)...)
}
