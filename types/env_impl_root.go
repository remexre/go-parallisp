package types

import (
	"errors"
)

type rootEnvImpl map[Symbol]Expr

// NewRootEnv creates a new environment with no parent, containing the given
// values.
func NewRootEnv(m map[Symbol]Expr) Env {
	if m == nil {
		return make(rootEnvImpl)
	}
	return rootEnvImpl(m)
}

func (env rootEnvImpl) Derive(other Env) Env {
	return newDerivedEnv(env, other)
}

func (env rootEnvImpl) Get(sym Symbol) (Expr, bool) {
	val, ok := env[sym]
	return val, ok
}

func (env rootEnvImpl) Def(sym Symbol, val Expr) error {
	return errors.New("parallisp.interpreter: root env is immutable")
}

func (env rootEnvImpl) List(recursive bool) []Symbol {
	var out []Symbol
	for sym := range env {
		out = append(out, sym)
	}
	return out
}

func (env rootEnvImpl) Set(sym Symbol, val Expr) error {
	return errors.New("parallisp.interpreter: root env is immutable")
}
