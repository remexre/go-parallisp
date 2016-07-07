package builtins

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

type function struct {
	name     types.Symbol
	args     []types.Symbol
	variadic bool

	body []types.Expr
	env  types.Env
}

// Call evaluates the function.
func (f *function) Call(args ...types.Expr) (types.Expr, error) {
	if f.variadic && len(args) < len(f.args)-1 {
		return nil, fmt.Errorf("%s: incorrect argn", f.name)
	} else if !f.variadic && len(args) != len(f.args) {
		return nil, fmt.Errorf("%s: incorrect argn", f.name)
	}

	childEnv := f.env.Derive(nil)
	for i, argName := range f.args {
		var argVal types.Expr
		if f.variadic && i == len(f.args)-1 {
			argVal = types.NewConsList(args[i:]...)
		} else {
			argVal = args[i]
		}
		if err := childEnv.Def(argName, argVal); err != nil {
			return nil, err
		}
	}

	return Progn(childEnv, f.body...)
}

// Eval evaluates a special form.
func (*function) Eval(env types.Env) (types.Expr, error) {
	return nil, errors.New("builtins: cannot eval a function")
}

func (f *function) MinArgN() int {
	if f.variadic {
		return len(f.args) - 1
	}
	return len(f.args)
}
func (f *function) MaxArgN() int {
	if f.variadic {
		return -1
	}
	return len(f.args)
}

// String converts the special form to a string.
func (f *function) String() string {
	return fmt.Sprintf("user-def-function-%s", f.name)
}

// Type returns the type of the special form.
func (*function) Type() string {
	return "function"
}
