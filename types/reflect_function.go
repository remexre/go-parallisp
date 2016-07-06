package types

import (
	"errors"
	"fmt"
	"reflect"
)

type reflectFunction struct {
	fn      reflect.Value
	t       reflect.Type
	minArgN int
	maxArgN int
}

func (fn *reflectFunction) MinArgN() int { return fn.minArgN }
func (fn *reflectFunction) MaxArgN() int { return fn.maxArgN }

func (fn *reflectFunction) Eval(env Env) (Expr, error) {
	return nil, errors.New("parallisp.types: cannot eval a function")
}

func (fn *reflectFunction) String() string {
	return fmt.Sprintf("function-%p", fn.fn.Interface())
}

func (*reflectFunction) Type() string {
	return "function"
}
