package interpreter

import (
	"remexre.xyz/go-parallisp/interpreter/natives"
	"remexre.xyz/go-parallisp/parser"
	"remexre.xyz/go-parallisp/types"
)

//go:generate ./make-prelude.sh

// NewEnv creates and returns a new types.Env initialized with "reasonable"
// builtins.
func NewEnv(m map[types.Symbol]types.Expr) types.Env {
	ms := []map[types.Symbol]types.Expr{m}
	for _, env := range DefaultEnvs {
		ms = append(ms, LoadedEnvs[env].All(false))
	}
	return types.NewRootEnv(ms...)
}

// DefaultEnvs is the list of environments included in NewEnv's return.
var DefaultEnvs = []string{
	"natives",
}

// LoadedEnvs is the list of environments already imported.
var LoadedEnvs = map[string]types.Env{
	"natives": natives.Env,
}

func init() {
	exprs, err := parser.Parse(Prelude)
	if err != nil {
		panic(err)
	}

	env := NewEnv(nil).Derive(nil)
	for _, expr := range exprs {
		_, err = types.EvalExpr(env, expr)
		if err != nil {
			panic(err)
		}
	}

	LoadedEnvs["import"] = types.NewRootEnv(map[types.Symbol]types.Expr{
		"import": types.SpecialFormFunc(Import),
	})
	LoadedEnvs["prelude"] = env
	DefaultEnvs = append(DefaultEnvs, "import", "prelude")
}
