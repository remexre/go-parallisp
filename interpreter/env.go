package interpreter

import (
	"github.com/remexre/go-parallisp/debug"
	debugModule "github.com/remexre/go-parallisp/interpreter/debug"
	"github.com/remexre/go-parallisp/interpreter/exec"
	"github.com/remexre/go-parallisp/interpreter/fs"
	"github.com/remexre/go-parallisp/interpreter/natives"
	"github.com/remexre/go-parallisp/interpreter/process"
	"github.com/remexre/go-parallisp/types"
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
	"debug":   debugModule.Env,
	"exec":    exec.Env,
	"fs":      fs.Env,
	"natives": natives.Env,
	"process": process.Env,
}

func init() {
	debug.Log("import", `importing "prelude"`)
	_, prelude, err := Interpret(Prelude, "")
	if err != nil {
		panic(err)
	}
	debug.Log("import", `imported "prelude"`)

	LoadedEnvs["import"] = types.NewRootEnv(map[types.Symbol]types.Expr{
		"import": types.NativeFunc(Import),
	})
	LoadedEnvs["prelude"] = prelude
	DefaultEnvs = append(DefaultEnvs, "import", "prelude")
}
