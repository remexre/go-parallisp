package types

type rwEnvImpl struct {
	Parent    Env
	Variables map[Symbol]Expr
}

func (env *rwEnvImpl) Derive(vars map[Symbol]Expr) Env {
	if vars == nil {
		vars = make(map[Symbol]Expr)
	}
	return &rwEnvImpl{
		Parent:    env,
		Variables: vars,
	}
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
	if env.Parent != nil {
		if _, ok := env.Parent.Get(sym); ok {
			return env.Parent.Set(sym, val)
		}
	}
	return env.Def(sym, val)
}
