package ast

import (
	"fmt"

	"github.com/remexre/go-parallisp/types"
	"github.com/remexre/go-parallisp/util/exprset"
	"github.com/remexre/go-parallisp/util/stringset"
)

// Lambda represents a lambda expression which accepts multiple parameters.
type Lambda struct {
	Params []string
	Doc    string
	Body   Progn
}

// NewLambda returns a new lambda from the expressions in its form, excluding
// the initial lambda symbol.
func NewLambda(exprs []types.Expr) (Node, error) {
	argVector, ok := exprs[0].(types.Vector)
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
	lambda.Doc, lambda.Body, err = ConvertDocProgn(exprs[1:])
	return lambda, err
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Lambda) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (l *Lambda) FreeVars() stringset.StringSet {
	return l.Body.FreeVars().Difference(stringset.New(l.Params...))
}

// IsLiteral returns whether the node is a literal.
func (*Lambda) IsLiteral() bool { return false }

// Literals returns the constants used in this node and all child nodes.
func (l *Lambda) Literals() exprset.ExprSet {
	return l.Body.Literals().Add(types.String(l.Doc))
}

// ToExpr converts the node to an expr.
func (l *Lambda) ToExpr() types.Expr {
	// Build the parameter list.
	params := make([]types.Expr, len(l.Params))
	for i, param := range l.Params {
		params[i] = types.Symbol(param)
	}

	// Build the body.
	body := make([]types.Expr, len(l.Body))
	for i, node := range l.Body {
		body[i] = node.ToExpr()
	}

	// Build the defmacro and return.
	return types.NewConsList(append([]types.Expr{
		types.Symbol("lambda"),
		types.Vector(params),
		types.String(l.Doc),
	}, body...)...)
}
