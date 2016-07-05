package interpreter

import "remexre.xyz/parallisp/types"

type rwEnvImpl struct {
	Parent    types.Env
	Variables map[types.Symbol]types.Expr
}

func (env *rwEnvImpl) Derive(vars map[types.Symbol]types.Expr) types.Env {
	return &rwEnvImpl{
		Parent:    env,
		Variables: vars,
	}
}

func (env *rwEnvImpl) Get(sym types.Symbol) (types.Expr, bool) {
	if val, ok := env.Variables[sym]; ok {
		return val, true
	} else if env.Parent != nil {
		return env.Parent.Get(sym)
	}
	return nil, false
}

func (env *rwEnvImpl) Def(sym types.Symbol, val types.Expr) error {
	env.Variables[sym] = val
	return nil
}

func (env *rwEnvImpl) List(recursive bool) []types.Symbol {
	var out []types.Symbol
	for sym := range env.Variables {
		out = append(out, sym)
	}
	if recursive && env.Parent != nil {
		out = append(out, env.Parent.List(true)...)
	}
	return out
}

func (env *rwEnvImpl) Set(sym types.Symbol, val types.Expr) error {
	if env.Parent != nil {
		if _, ok := env.Parent.Get(sym); ok {
			return env.Parent.Set(sym, val)
		}
	}
	return env.Def(sym, val)
}
