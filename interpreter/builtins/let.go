package builtins

import (
	"errors"

	"remexre.xyz/go-parallisp/ast"
	"remexre.xyz/go-parallisp/types"
)

// Let is the let special form.
func Let(env types.Env, args ...types.Expr) (types.Expr, error) {
	if len(args) < 2 {
		return nil, errors.New("let: insufficient arguments")
	}

	argDefs, ok := args[0].(types.Vector)
	if !ok {
		return nil, errors.New("let: definitions not a vector")
	}
	defs := make([]ast.LetDefinition, len(argDefs))
	for i, argDef := range argDefs {
		def, ok := argDef.(types.Vector)
		if !ok {
			return nil, errors.New("let: definition not a vector")
		}
		if len(def) != 2 {
			return nil, errors.New("let: definition vector has wrong length")
		}

		name, ok := def[0].(types.Symbol)
		if !ok {
			return nil, errors.New("let: name not a symbol")
		}

		value, err := types.EvalExpr(env, def[1])
		if err != nil {
			return nil, err
		}

		defs[i] = ast.LetDefinition{
			Name:  name,
			Value: value,
		}
	}

	// Populate the environment
	env = env.Derive(nil)
	for _, def := range defs {
		if err := env.Def(def.Name, def.Value); err != nil {
			return nil, err
		}
	}

	// Evaluate the body.
	return Progn(env, args[1:]...)
}
