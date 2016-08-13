package types

import (
	"errors"
	"fmt"
	"reflect"
)

type reflectFunction struct {
	name    string
	fn      reflect.Value
	t       reflect.Type
	minArgN int
	maxArgN int
}

func (fn *reflectFunction) Name() string {
	return fn.NameOr(fmt.Sprintf("%p", fn.fn.Interface()))
}
func (fn *reflectFunction) NameOr(defaultName string) string {
	if fn.name == "" {
		return defaultName
	}
	return fn.name
}

func (fn *reflectFunction) Eval(env Env) (Expr, error) {
	return nil, errors.New("parallisp.types: cannot eval a function")
}

// LiteralAsm converts an expression to its representation in AT&T syntax x86-64
// assembly.
func (reflectFunction) LiteralAsm() string {
	panic("types.Function.LiteralAsm: cannot make a function into a literal")
}

func (fn *reflectFunction) String() string {
	return "reflected-function-" + fn.Name()
}

func (*reflectFunction) Type() string {
	return "function"
}
