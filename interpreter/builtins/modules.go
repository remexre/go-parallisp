package builtins

import (
	"remexre.xyz/go-parallisp/interpreter/builtins/fs"
	"remexre.xyz/go-parallisp/interpreter/builtins/process"
	"remexre.xyz/go-parallisp/types"
)

// Modules is a map of builtin modules.
var Modules = map[string]types.Env{
	"fs":      fs.Env,
	"process": process.Env,
}
