package types

import (
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
	return fn, nil
}

func (fn *reflectFunction) String() string {
	return fmt.Sprintf("%d", fn.fn.UnsafeAddr())
}

func (*reflectFunction) Type() string {
	return "function"
}
