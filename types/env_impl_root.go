package types

import (
	"errors"
)

type rootEnvImpl map[Symbol]Expr

// NewRootEnv creates a new environment with no parent, containing the given
// values.
func NewRootEnv(ms ...map[Symbol]Expr) Env {
	if ms == nil {
		return make(rootEnvImpl)
	}
	env := make(map[Symbol]Expr)
	for _, m := range ms {
		for k, v := range m {
			env[k] = v
		}
	}
	return rootEnvImpl(env)
}

func (env rootEnvImpl) All(bool) map[Symbol]Expr {
	m := make(map[Symbol]Expr)
	for k, v := range env {
		m[k] = v
	}
	return m
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
