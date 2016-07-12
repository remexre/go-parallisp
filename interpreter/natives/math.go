package natives

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	"remexre.xyz/go-parallisp/types"
)

// Add adds numbers or concatenates strings.
func Add(args ...types.Expr) (types.Expr, error) {
	if args, ok := allIntegers(args); ok {
		var i types.Integer
		for _, arg := range args {
			i += arg
		}
		return i, nil
	}
	if args, ok := allNumbers(args); ok {
		var i types.Floating
		for _, arg := range args {
			i += arg
		}
		return i, nil
	}
	if args, ok := allStrings(args); ok {
		buf := new(bytes.Buffer)
		for _, arg := range args {
			buf.WriteString(string(arg))
		}
		return types.String(buf.String()), nil
	}
	return nil, fmt.Errorf("+: cannot add non-numbers or concatenate non-strings %s", args)
}

// Subtract subtracts numbers.
func Subtract(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return types.Integer(0), nil
	}
	if args, ok := allIntegers(args); ok {
		i := args[0]
		for _, arg := range args[1:] {
			i -= arg
		}
		return i, nil
	}
	if args, ok := allNumbers(args); ok {
		i := args[0]
		for _, arg := range args[1:] {
			i -= arg
		}
		return i, nil
	}
	return nil, fmt.Errorf("-: cannot subtract non-numbers %s", args)
}

// Multiply multiplies numbers or repeats a string.
func Multiply(args ...types.Expr) (types.Expr, error) {
	if args, ok := allIntegers(args); ok {
		i := types.Integer(1)
		for _, arg := range args {
			i *= arg
		}
		return i, nil
	}
	if args, ok := allNumbers(args); ok {
		i := types.Floating(1)
		for _, arg := range args {
			i *= arg
		}
		return i, nil
	}
	if len(args) == 2 {
		str, ok1 := args[0].(types.String)
		n, ok2 := args[1].(types.Integer)
		if ok1 && ok2 && n >= 0 {
			return types.String(strings.Repeat(string(str), int(n))), nil
		}
	}
	return nil, fmt.Errorf("*: cannot multiply or repeat %s", args)
}

// Divide divides numbers.
func Divide(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return types.Integer(0), nil
	}
	if args, ok := allIntegers(args); ok {
		i := args[0]
		for _, arg := range args[1:] {
			i /= arg
		}
		return i, nil
	}
	if args, ok := allNumbers(args); ok {
		i := args[0]
		for _, arg := range args[1:] {
			i /= arg
		}
		return i, nil
	}
	return nil, fmt.Errorf("/: cannot divide non-numbers %s", args)
}

// Modulo performs the modulus operation on integers.
func Modulo(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return types.Integer(0), nil
	}
	if args, ok := allIntegers(args); ok {
		i := args[0]
		for _, arg := range args[1:] {
			i %= arg
		}
		return i, nil
	}
	return nil, fmt.Errorf("%%: cannot divide non-integers %s", args)
}

// Pow exponentiates numbers.
func Pow(a, b types.Expr) (types.Expr, error) {
	args := []types.Expr{a, b}
	if ints, ok := allIntegers(args); ok {
		return types.Integer(math.Pow(float64(ints[0]), float64(ints[1]))), nil
	} else if nums, ok := allNumbers(args); ok {
		return types.Floating(math.Pow(float64(nums[0]), float64(nums[1]))), nil
	}
	return nil, fmt.Errorf("pow: cannot exponentiate non-numbers %s %s", a, b)
}
