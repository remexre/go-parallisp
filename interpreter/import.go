package interpreter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"remexre.xyz/go-parallisp/debug"
	"remexre.xyz/go-parallisp/interpreter/builtins"
	"remexre.xyz/go-parallisp/types"
)

var cachedEnvs = make(map[string]types.Env)

// Import is the import special form.
func Import(env types.Env, exprs ...types.Expr) (types.Expr, error) {
	if len(exprs) != 2 {
		return nil, errors.New("parallisp.types: invalid import")
	}

	path, ok := exprs[0].(types.String)
	if !ok {
		return nil, errors.New("parallisp.types: invalid import")
	}

	var importEnv types.Env
	if importEnv, ok = builtins.Modules[string(path)]; !ok {
		currentFileExpr, ok := env.Get(types.Symbol("**current-file**"))
		if !ok {
			currentFileExpr = types.String("./dummy.file.ignore.this")
		}
		currentFile, ok := currentFileExpr.(types.String)
		if !ok {
			return nil, fmt.Errorf("current file not a string: %s", currentFileExpr)
		}

		currentDir := filepath.Dir(string(currentFile))
		importPath := filepath.Join(currentDir, string(path))

		if importEnv, ok = cachedEnvs[importPath]; !ok {
			b, err := ioutil.ReadFile(importPath)
			if err != nil {
				return nil, fmt.Errorf("import: %s", err.Error())
			}

			_, importEnv, err = Interpret(string(b), importPath)
			if err != nil {
				return nil, err
			}

			cachedEnvs[importPath] = importEnv
			debug.Log("import", "imported %v from %s", importEnv.List(false), importPath)
		}
	}

	var syms []types.Symbol
	if exprs[1] == types.Symbol("*") {
		syms = importEnv.List(true)
	} else if symsIn, ok := exprs[1].(types.Vector); ok {
		syms = make([]types.Symbol, len(symsIn))
		for i, symIn := range symsIn {
			syms[i], ok = symIn.(types.Symbol)
			if !ok {
				return nil, errors.New("parallisp.types: invalid import")
			}
		}
	} else {
		return nil, errors.New("parallisp.types: invalid import")
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
