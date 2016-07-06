package interpreter

import (
	"errors"

	"remexre.xyz/go-parallisp/types"
)

type rootEnvImpl map[types.Symbol]types.Expr

func (env rootEnvImpl) Derive(vars map[types.Symbol]types.Expr) types.Env {
	return &rwEnvImpl{
		Parent:    env,
		Variables: vars,
	}
}

func (env rootEnvImpl) Get(sym types.Symbol) (types.Expr, bool) {
	val, ok := env[sym]
	return val, ok
}

func (env rootEnvImpl) Def(sym types.Symbol, val types.Expr) error {
	return errors.New("parallisp.interpreter: root env is immutable")
}

func (env rootEnvImpl) List(recursive bool) []types.Symbol {
	var out []types.Symbol
	for sym := range env {
		out = append(out, sym)
	}
	return out
}

func (env rootEnvImpl) Set(sym types.Symbol, val types.Expr) error {
	return errors.New("parallisp.interpreter: root env is immutable")
}
