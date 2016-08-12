package ast

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// Let represents a let-expression.
type Let struct {
	Definitions []LetDefinition
	Body        Progn
	Sequential  bool
}

// LetDefinition is one let definition.
type LetDefinition struct {
	Label string
	Value Node
}

// NewLet returns a new let from the expressions in its form, excluding the
// initial let symbol.
func NewLet(exprs []types.Expr) (Node, error) {
	defExprCons, ok := exprs[0].(types.Cons)
	if !ok {
		return nil, fmt.Errorf("ast.Convert: invalid let")
	} else if !defExprCons.IsList() {
		return nil, fmt.Errorf("ast.Convert: invalid let")
	}
	defExprs := defExprCons.ToSlice()
	defs := make([]LetDefinition, len(defExprs))
	for i, expr := range defExprs {
		exprCons, ok := expr.(types.Cons)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid let definition")
		} else if !exprCons.IsList() {
			return nil, fmt.Errorf("ast.Convert: invalid let definition")
		}
		def := exprCons.ToSlice()
		if len(def) != 2 {
			return nil, fmt.Errorf("ast.Convert: invalid let definition")
		}

		label, ok := def[0].(types.String)
		if !ok {
			return nil, fmt.Errorf("ast.Convert: invalid let definition label")
		}
		value, err := Convert(def[1])
		if err != nil {
			return nil, err
		}

		defs[i] = LetDefinition{string(label), value}
	}

	body, err := ConvertProgn(exprs)
	if err != nil {
		return nil, err
	}

	return &Let{
		defs,
		body,
		false,
	}, nil
}

// NewSequentialLet returns a new sequential let from the expressions in its
// form, excluding the initial let* symbol.
func NewSequentialLet(exprs []types.Expr) (Node, error) {
	let, err := NewLet(exprs)
	if err != nil {
		return nil, err
	}
	let.(*Let).Sequential = true
	return let, nil
}

// Constants returns the constants used in this node and all child nodes.
func (l *Let) Constants() exprset.ExprSet {
	sets := make([]exprset.ExprSet, len(l.Definitions))
	for i, def := range l.Definitions {
		sets[i] = def.Value.Constants()
	}
	return l.Body.Constants().Union(sets...)
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (*Let) Defines() stringset.StringSet { return nil }

// FreeVars returns the free values contained within a node, recursively.
func (l *Let) FreeVars() stringset.StringSet {
	freeVars := l.Body.FreeVars()
	if l.Sequential {
		for i := len(l.Definitions) - 1; i >= 0; i-- {
			def := l.Definitions[i]
			freeVars = freeVars.
				Difference(stringset.New(def.Label)).
				Union(def.Value.FreeVars())
		}
	} else {
		length := len(l.Definitions)
		names := make([]string, 0, length)
		nodes := make([]Node, 0, length)
		for _, def := range l.Definitions {
			names = append(names, def.Label)
			nodes = append(nodes, def.Value)
		}
		nodeFreeVars := make([]stringset.StringSet, length)
		for i, node := range nodes {
			nodeFreeVars[i] = node.FreeVars()
		}
		freeVars = freeVars.
			Difference(stringset.New(names...)).
			Union(nodeFreeVars...)
	}
	return freeVars
}

// ToExpr converts the node to an expr.
func (l *Let) ToExpr() types.Expr {
	// Choose the symbol.
	symbol := types.Symbol("let")
	if l.Sequential {
		symbol += "*"
	}

	// Build the definition list.
	defs := make([]types.Expr, len(l.Definitions))
	for i, def := range l.Definitions {
		defs[i] = types.NewConsList(types.Symbol(def.Label), def.Value.ToExpr())
	}

	// Build the body.
	body := make([]types.Expr, len(l.Body))
	for i, node := range l.Body {
		body[i] = node.ToExpr()
	}

	// Build the defmacro and return.
	return types.NewConsList(append([]types.Expr{
		symbol,
		types.NewConsList(defs...),
	}, body...)...)
}
