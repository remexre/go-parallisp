package types

import "remexre.xyz/go-parallisp/debug"

type rwEnvImpl struct {
	Parent    Env
	Variables map[Symbol]Expr
}

func newDerivedEnv(parent, other Env) Env {
	new := &rwEnvImpl{
		Parent:    parent,
		Variables: make(map[Symbol]Expr),
	}
	if other != nil {
		for _, sym := range other.List(true) {
			if val, ok := other.Get(sym); ok {
				new.Variables[sym] = val
			}
		}
	}
	return new
}

func (env rwEnvImpl) All(recursive bool) map[Symbol]Expr {
	m := make(map[Symbol]Expr)
	if recursive && env.Parent != nil {
		for k, v := range env.Parent.All(true) {
			m[k] = v
		}
	}
	for k, v := range env.Variables {
		m[k] = v
	}
	return m
}

func (env *rwEnvImpl) Derive(other Env) Env {
	return newDerivedEnv(env, other)
}

func (env *rwEnvImpl) Get(sym Symbol) (Expr, bool) {
	if val, ok := env.Variables[sym]; ok {
		return val, true
	} else if env.Parent != nil {
		return env.Parent.Get(sym)
	}
	return nil, false
}

func (env *rwEnvImpl) Def(sym Symbol, val Expr) error {
	debug.Log("env-def", "defining %s as %v", sym, val)
	env.Variables[sym] = val
	return nil
}

func (env *rwEnvImpl) List(recursive bool) []Symbol {
	var out []Symbol
	for sym := range env.Variables {
		out = append(out, sym)
	}
	if recursive && env.Parent != nil {
		out = append(out, env.Parent.List(true)...)
	}
	return out
}

func (env *rwEnvImpl) Set(sym Symbol, val Expr) error {
	debug.Log("env-set", "setting %s to %v", sym, val)
	if env.Parent != nil {
		if _, ok := env.Parent.Get(sym); ok {
			return env.Parent.Set(sym, val)
		}
	}
	return env.Def(sym, val)
}
