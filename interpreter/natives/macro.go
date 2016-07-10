package natives

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

type macro struct {
	name     types.Symbol
	args     []types.Symbol
	variadic bool

	body []types.Expr
}

// CallMacro evaluates the macro.
func (m *macro) CallMacro(env types.Env, args ...types.Expr) (types.Expr, error) {
	if m.variadic && len(args) < len(m.args)-1 {
		return nil, fmt.Errorf("%s: incorrect argn", m.name)
	} else if !m.variadic && len(args) != len(m.args) {
		return nil, fmt.Errorf("%s: incorrect argn", m.name)
	}

	childEnv := env.Derive(nil)
	for i, argName := range m.args {
		var argVal types.Expr
		if m.variadic && i == len(m.args)-1 {
			argVal = types.NewConsList(args[i:]...)
		} else {
			argVal = args[i]
		}
		if err := childEnv.Def(argName, argVal); err != nil {
			return nil, err
		}
	}

	return Progn(childEnv, m.body...)
}

// CallSpecialForm evaluates the macro, then its return.
func (m *macro) CallSpecialForm(env types.Env, args ...types.Expr) (types.Expr, error) {
	code, err := m.CallMacro(env, args...)
	if err != nil {
		return nil, err
	}

	return types.EvalExpr(env, code)
}

// Eval evaluates a special form.
func (m *macro) Eval(env types.Env) (types.Expr, error) {
	return nil, errors.New("builtins: cannot eval a macro")
}

// String converts the special form to a string.
func (m *macro) String() string {
	return fmt.Sprintf("macro-%s", m.name)
}

// Type returns the type of the special form.
func (m *macro) Type() string {
	return "macro"
}
