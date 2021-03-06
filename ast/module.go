package ast

import (
	"flag"
	"io/ioutil"

	"github.com/remexre/go-parallisp/parser"
	"github.com/remexre/go-parallisp/types"
	"github.com/remexre/go-parallisp/util/stringset"
)

// Module represents a parallisp module.
type Module struct {
	Name    string
	Imports []*Import
	Body    Progn
}

// LoadModule loads a module from a file.
func LoadModule(filename string) (*Module, error) {
	src, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	exprs, err := parser.Parse(string(src))
	if err != nil {
		panic(err)
	}

	return ConvertModule(filename, exprs)
}

// ConvertModule converts a slice of exprs into a Module and registers it with
// the given name.
func ConvertModule(name string, exprs []types.Expr) (*Module, error) {
	// Convert the body to AST nodes.
	body := make(Progn, len(exprs))
	for i, expr := range exprs {
		var err error
		body[i], err = Convert(expr)
		if err != nil {
			return nil, err
		}
	}

	// Split off imports.
	var imports []*Import
	for len(body) > 0 {
		importNode, ok := body[0].(*Import)
		if !ok {
			break
		}
		imports = append(imports, importNode)
		body = body[1:]
	}

	// Return.
	return &Module{
		name,
		imports,
		body,
	}, nil
}

// FreeVars returns a slice of the free variables for the entire parallisp
// module whose AST is passed as input. Imported variables are removed from
// consideration, but standard library functions are not.
func (module *Module) FreeVars() stringset.StringSet {
	imported := stringset.New()
	for _, importNode := range module.Imports {
		imported.Add(importNode.Defines().ToSlice()...)
	}

	return module.Body.FreeVars().Difference(imported)
}
