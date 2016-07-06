package interpreter

import (
	"remexre.xyz/go-parallisp/interpreter/builtins"
	"remexre.xyz/go-parallisp/types"
)

// NewEnv creates and returns a new types.Env initialized with "reasonable"
// builtins.
func NewEnv() types.Env {
	return rootEnvImpl{
		"println": types.MustNewReflectFunction(builtins.Println),
	}
}
