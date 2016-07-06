package types

import (
	"fmt"
	"reflect"
)

func (fn *reflectFunction) Call(_ Env, exprs ...Expr) (Expr, error) {
	// Check argument number.
	if fn.minArgN > 0 && len(exprs) < fn.minArgN {
		return nil, fmt.Errorf("parallisp.types: insufficient arguments: wanted %d, got %d", fn.minArgN, len(exprs))
	} else if fn.maxArgN > 0 && len(exprs) > fn.maxArgN {
		return nil, fmt.Errorf("parallisp.types: too many arguments: wanted %d, got %d", fn.maxArgN, len(exprs))
	}

	// Convert the args to reflect.Values.
	var args []reflect.Value
	if fn.t.IsVariadic() {
		for _, expr := range exprs {
			args = append(args, reflect.ValueOf(expr))
		}
	} else {
		args = make([]reflect.Value, len(exprs))
		for i, expr := range exprs {
			val := reflect.ValueOf(expr)
			if t1, t2 := val.Type(), fn.t.In(i); !t1.ConvertibleTo(t2) {
				return nil, fmt.Errorf("parallisp.types: cannot convert %s to %s", t1, t2)
			}
			args[i] = reflect.ValueOf(expr)
		}
	}

	// Call the function.
	out := fn.fn.Call(args)

	// Return.
	if fn.t.NumOut() == 0 {
		return nil, nil
	}
	return out[0].Interface().(Expr), nil
}
