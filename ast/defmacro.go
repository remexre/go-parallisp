package ast

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// Defmacro represents a macro definition.
type Defmacro struct {
	Name   string
	Params []string
	Doc    string
	Body   Progn
}

// NewDefmacro returns a new defmacro from the expressions in its form,
// excluding the initial defmacro symbol.
func NewDefmacro(exprs []types.Expr) (Node, error) {
	argVector, ok := exprs[1].(types.Vector)
	if !ok {
		return nil, fmt.Errorf("ast.Convert: invalid defmacro")
	}
	name, ok := exprs[0].(types.Symbol)
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
	defmacro.Doc, defmacro.Body, err = ConvertDocProgn(exprs[2:])
	return defmacro, err
}

// Constants returns the constants used in this node and all child nodes.
func (d *Defmacro) Constants() exprset.ExprSet {
	return d.Body.Constants().Add(types.String(d.Doc))
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (d *Defmacro) Defines() stringset.StringSet {
	return stringset.New(d.Name)
}

// FreeVars returns the free values contained within a node, recursively.
func (d *Defmacro) FreeVars() stringset.StringSet {
	return d.Body.FreeVars().Difference(stringset.New(d.Params...))
}

// ToExpr converts the node to an expr.
func (d *Defmacro) ToExpr() types.Expr {
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
		types.Symbol("defmacro"),
		types.Symbol(d.Name),
		types.Vector(params),
		types.String(d.Doc),
	}, body...)...)
}
