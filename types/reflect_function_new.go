package types

import (
	"fmt"
	"reflect"
)

// MustNewReflectFunction wraps a Go function using reflection and returns it,
// panicking on error.
func MustNewReflectFunction(name string, fn interface{}) Function {
	expr, err := NewReflectFunction(name, fn)
	if err != nil {
		panic(err)
	}
	return expr
}

// NewReflectFunction wraps a Go function using reflection and returns it, or an
// error if one occurred.
func NewReflectFunction(name string, fn interface{}) (Function, error) {
	val := reflect.ValueOf(fn)
	typ := val.Type()
	if typ.Kind() != reflect.Func {
		return nil, fmt.Errorf("parallisp.types: not a function: %v", val)
	} else if numOut := typ.NumOut(); numOut > 2 {
		return nil, fmt.Errorf("parallisp.types: invalid function signature")
	}
	out := &reflectFunction{
		name:    name,
		fn:      val,
		t:       typ,
		minArgN: typ.NumIn(),
		maxArgN: typ.NumIn(),
	}
	if typ.IsVariadic() {
		out.minArgN--
		out.maxArgN = -1
	}
	return out, nil
}
