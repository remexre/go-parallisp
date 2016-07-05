package interpreter

import (
	"remexre.xyz/parallisp/interpreter/builtins"
	"remexre.xyz/parallisp/types"
)

// NewEnv creates and returns a new types.Env initialized with "reasonable"
// builtins.
func NewEnv() types.Env {
	return rootEnvImpl{
		"println": types.MustNewReflectFunction(builtins.Println),
	}
}
