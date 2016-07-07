package interpreter

import (
	"errors"
	"fmt"
	"io/ioutil"

	"remexre.xyz/go-parallisp/interpreter/builtins"
	"remexre.xyz/go-parallisp/types"
)

// Import is the import special form.
func Import(env types.Env, exprs ...types.Expr) (types.Expr, error) {
	if len(exprs) != 2 {
		return nil, errors.New("parallisp.types: invalid import")
	}

	path, ok := exprs[0].(types.String)
	if !ok {
		return nil, errors.New("parallisp.types: invalid import")
	}

	symsIn, ok := exprs[1].(types.Vector)
	if !ok {
		return nil, errors.New("parallisp.types: invalid import")
	}

	syms := make([]types.Symbol, len(symsIn))
	for i, symIn := range symsIn {
		syms[i], ok = symIn.(types.Symbol)
		if !ok {
			return nil, errors.New("parallisp.types: invalid import")
		}
	}

	var importEnv types.Env
	if importEnv, ok = builtins.Modules[string(path)]; !ok {
		b, err := ioutil.ReadFile(string(path))
		if err != nil {
			return nil, fmt.Errorf("import: %s", err.Error())
		}

		_, importEnv, err = Interpret(string(b))
		if err != nil {
			return nil, err
		}
	}

	for _, sym := range syms {
		val, ok := importEnv.Get(sym)
		if !ok {
			return nil, fmt.Errorf("import: %s does not contain %s", path, sym)
		}
		if err := env.Def(sym, val); err != nil {
			return nil, err
		}
	}

	return nil, nil
}
