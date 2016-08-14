package interpreterTypes

import (
	"fmt"

	"github.com/remexre/go-parallisp/types"
)

// FunctionLike is a type for functions or macros.
type FunctionLike struct {
	doc   string
	macro bool

	name     types.Symbol
	args     []types.Symbol
	variadic bool

	body []types.Expr
	env  types.Env
}

// NewFunctionLike creates a new function or macro.
func NewFunctionLike(macro bool, env types.Env, name types.Symbol, args types.Vector, body ...types.Expr) (*FunctionLike, error) {
	var argSyms []types.Symbol
	variadic := false
	for _, arg := range args {
		sym, ok := arg.(types.Symbol)
		if !ok {
			return nil, fmt.Errorf("invalid function argument %v", arg)
		}
		if sym[0] == '&' {
			switch string(sym) {
			case "&rest":
				variadic = true
			default:
				return nil, fmt.Errorf("unrecognized meta-argument %s", sym)
			}
		} else {
			argSyms = append(argSyms, sym)
		}
	}

	doc := ""
	if len(body) >= 2 {
		if docStr, ok := body[0].(types.String); ok {
			doc = string(docStr)
			body = body[1:]
		}
	}

	return &FunctionLike{
		doc:   doc,
		macro: macro,

		name:     name,
		args:     argSyms,
		variadic: variadic,

		body: body,
		env:  env,
	}, nil
}

// CallRaw calls the FunctionLike and returns its output. It returns the wrong
// output for both functions and macros!
func (f *FunctionLike) CallRaw(env types.Env, args ...types.Expr) (types.Expr, error) {
	if f.variadic && len(args) < len(f.args)-1 {
		return nil, fmt.Errorf("%s: incorrect argn", f.name)
	} else if !f.variadic && len(args) != len(f.args) {
		return nil, fmt.Errorf("%s: incorrect argn", f.name)
	}

	childEnv := env.Derive(nil)
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

// Call evaluates the function or macro.
func (f *FunctionLike) Call(env types.Env, args ...types.Expr) (types.Expr, error) {
	if f.macro {
		code, err := f.CallRaw(env, args...)
		if err != nil {
			return nil, err
		}

		return types.EvalExpr(env, code)
	}
	argVals := make([]types.Expr, len(args))
	for i, arg := range args {
		var err error
		argVals[i], err = types.EvalExpr(env, arg)
		if err != nil {
			return nil, err
		}
	}
	return f.CallRaw(f.env, argVals...)
}

// Doc returns the documentation for this expression.
func (f *FunctionLike) Doc() string {
	return f.doc
}

// Eval evaluates a special form.
func (f *FunctionLike) Eval(env types.Env) (types.Expr, error) {
	return nil, fmt.Errorf("interpreterTypes.FunctionLike: cannot eval a function-like: %s", f)
}

// String converts the special form to a string.
func (f *FunctionLike) String() string {
	if f.macro {
		return fmt.Sprintf("macro-%s", f.name)
	}
	return fmt.Sprintf("function-%s", f.name)
}

// Type returns the type of the special form.
func (f *FunctionLike) Type() string {
	if f.macro {
		return "macro"
	}
	return "function"
}
