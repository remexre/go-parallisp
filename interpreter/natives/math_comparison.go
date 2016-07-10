package natives

import (
	"errors"
	"fmt"

	"remexre.xyz/go-parallisp/types"
)

// Gt tests if one number is greater than another.
func Gt(args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, errors.New(">: incorrect argn")
	}
	if args, ok := allIntegers(args); ok {
		if args[0] > args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	if args, ok := allNumbers(args); ok {
		if args[0] > args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	return nil, fmt.Errorf(">: cannot compare non-numbers %s", args)
}

// Gte tests if one number is greater or equal to than another.
func Gte(args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, errors.New(">=: incorrect argn")
	}
	if args, ok := allIntegers(args); ok {
		if args[0] >= args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	if args, ok := allNumbers(args); ok {
		if args[0] >= args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	return nil, fmt.Errorf(">=: cannot compare non-numbers %s", args)
}

// Lt tests if one number is less than another.
func Lt(args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, errors.New("<: incorrect argn")
	}
	if args, ok := allIntegers(args); ok {
		if args[0] < args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	if args, ok := allNumbers(args); ok {
		if args[0] < args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	return nil, fmt.Errorf("<: cannot compare non-numbers %s", args)
}

// Lte tests if one number is less or equal to than another.
func Lte(args ...types.Expr) (types.Expr, error) {
	if len(args) != 2 {
		return nil, errors.New("<=: incorrect argn")
	}
	if args, ok := allIntegers(args); ok {
		if args[0] <= args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	if args, ok := allNumbers(args); ok {
		if args[0] <= args[1] {
			return types.Symbol("t"), nil
		}
		return nil, nil
	}
	return nil, fmt.Errorf("<=: cannot compare non-numbers %s", args)
}
