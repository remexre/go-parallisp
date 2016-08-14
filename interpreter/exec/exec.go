package exec

import (
	"os/exec"

	"github.com/remexre/go-parallisp/types"
)

// Exec executes a command.
func Exec(cmd types.String, argsIn ...types.String) (types.String, error) {
	args := make([]string, len(argsIn))
	for i, arg := range argsIn {
		args[i] = string(arg)
	}
	output, err := exec.Command(string(cmd)).Output()
	return types.String(output), err
}
