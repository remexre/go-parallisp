package natives

import (
	"fmt"

	"remexre.xyz/go-parallisp/ast"
	"remexre.xyz/go-parallisp/interpreter/types"
	"remexre.xyz/go-parallisp/types"
)

// Let is the let special form.
func Let(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("let: insufficient arguments: %v", args)
	}

	argDefsCons, ok := args[0].(types.Cons)
	if !ok {
		return nil, fmt.Errorf("let: definitions not a cons: %v", args[0])
	}
	if !argDefsCons.IsList() {
		return nil, fmt.Errorf("let: definitions not a list: %v", argDefsCons)
	}
	var defs []ast.LetDefinition
	for _, argDef := range argDefsCons.ToSlice() {
		defCons, ok := argDef.(types.Cons)
		if !ok {
			return nil, fmt.Errorf("let: definition not a cons: %v", argDef)
		}
		if !defCons.IsList() {
			return nil, fmt.Errorf("let: definition not a list: %v", defCons)
		}
		def := defCons.ToSlice()

		name, ok := def[0].(types.Symbol)
		if !ok {
			return nil, fmt.Errorf("let: name not a symbol: %v", def[0])
		}

		value, err := interpreterTypes.Progn(env, def[1:]...)
		if err != nil {
			return nil, err
		}

		defs = append(defs, ast.LetDefinition{
			Name:  name,
			Value: value,
		})
	}

	// Populate the environment
	env = env.Derive(nil)
	for _, def := range defs {
		if err := env.Def(def.Name, def.Value); err != nil {
			return nil, err
		}
	}

	// Evaluate the body.
	return interpreterTypes.Progn(env, args[1:]...)
}
