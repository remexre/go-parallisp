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

func (fn *reflectFunction) MinArgN() int { return fn.minArgN }
func (fn *reflectFunction) MaxArgN() int { return fn.maxArgN }

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

func (fn *reflectFunction) String() string {
	return "function-" + fn.Name()
}

func (*reflectFunction) Type() string {
	return "function"
}
