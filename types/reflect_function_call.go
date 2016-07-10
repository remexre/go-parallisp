package types

import (
	"fmt"
	"reflect"
)

func (fn *reflectFunction) Call(env Env, exprs ...Expr) (Expr, error) {
	name := fn.NameOr("parallisp.types.reflectFunction")

	// Check argument number.
	if fn.minArgN > 0 && len(exprs) < fn.minArgN {
		return nil, fmt.Errorf("%s: insufficient arguments: wanted %d, got %d", name, fn.minArgN, len(exprs))
	} else if fn.maxArgN > 0 && len(exprs) > fn.maxArgN {
		return nil, fmt.Errorf("%s: too many arguments: wanted %d, got %d", name, fn.maxArgN, len(exprs))
	}

	// Evaluate the arguments.
	for i, expr := range exprs {
		var err error
		exprs[i], err = EvalExpr(env, expr)
		if err != nil {
			return nil, err
		}
	}

	// Convert the args to reflect.Values.
	var args []reflect.Value
	if fn.t.IsVariadic() {
		for _, expr := range exprs {
			if expr == nil {
				exprType := reflect.TypeOf((*Expr)(nil)).Elem()
				val := reflect.Zero(exprType)
				args = append(args, val)
			} else {
				args = append(args, reflect.ValueOf(expr))
			}
		}
	} else {
		args = make([]reflect.Value, len(exprs))
		for i, expr := range exprs {
			val := reflect.ValueOf(expr)
			if expr == nil {
				exprType := reflect.TypeOf((*Expr)(nil)).Elem()
				val = reflect.Zero(exprType)
			}

			if t1, t2 := val.Type(), fn.t.In(i); !t1.ConvertibleTo(t2) {
				return nil, fmt.Errorf("%s: cannot convert %s to %s", name, t1, t2)
			}
			args[i] = val
		}
	}

	// Call the function.
	outVal := fn.fn.Call(args)

	// Return.
	numOut := fn.t.NumOut()
	if numOut == 0 {
		return nil, nil
	}

	// Check for errors.
	if numOut == 2 {
		err := outVal[1].Interface()
		if err != nil {
			if errOut, ok := err.(error); ok {
				return nil, errOut
			}
			return nil, fmt.Errorf("%v", err)
		}
	}

	out := outVal[0].Interface()
	if out == nil {
		return nil, nil
	}
	return out.(Expr), nil
}
