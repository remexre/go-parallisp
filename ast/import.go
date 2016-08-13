package ast

import (
	"fmt"

	"remexre.xyz/go-parallisp/types"
	"remexre.xyz/go-parallisp/util/exprset"
	"remexre.xyz/go-parallisp/util/stringset"
)

// Import represents an import of zero or more symbols from a module, or
// alternatively a "wildcard" import. An Import with a nil Symbols is a wildcard
// import, as opposed to a zero-length Symbols vector.
type Import struct {
	Module   string
	Symbols  []string
	Wildcard bool
}

// NewImport returns a new import from the expressions in its form, excluding
// the initial import symbol.
func NewImport(exprs []types.Expr) (Node, error) {
	module, ok := exprs[0].(types.String)
	if !ok {
		return nil, fmt.Errorf("ast.Convert: invalid import")
	}
	importNode := &Import{
		string(module),
		nil,
		false,
	}
	if exprs[1] == types.Symbol("*") {
		importNode.Wildcard = true
		return importNode, nil
	} else if syms, ok := exprs[1].(types.Vector); ok {
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
}

// Literals returns the constants used in this node and all child nodes.
func (i *Import) Literals() exprset.ExprSet {
	return nil
}

// Defines returns the symbols defined in the parent scope by this node,
// recursively.
func (i *Import) Defines() stringset.StringSet {
	if i.Wildcard {
		// basename := filepath.Base(i.Module)
		// namespace := strings.TrimSuffix(basename, filepath.Ext(basename))
		panic("TODO WILDCARD IMPORT DEFINES")
	}
	return stringset.New(i.Symbols...)
}

// FreeVars returns the free values contained within a node, recursively.
func (*Import) FreeVars() stringset.StringSet { return nil }

// ToExpr converts the node to an expr.
func (i *Import) ToExpr() types.Expr {
	var expr types.Expr
	if i.Wildcard {
		expr = types.Symbol("*")
	} else {
		syms := make(types.Vector, len(i.Symbols))
		for i, sym := range i.Symbols {
			syms[i] = types.Symbol(sym)
		}
		expr = syms
	}
	return types.NewConsList(
		types.Symbol("import"),
		types.String(i.Module),
		expr,
	)
}
