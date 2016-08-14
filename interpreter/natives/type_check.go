package natives

import "github.com/remexre/go-parallisp/types"

func allIntegers(args []types.Expr) ([]types.Integer, bool) {
	var out []types.Integer
	for _, arg := range args {
		if arg, ok := arg.(types.Integer); ok {
			out = append(out, arg)
			continue
		}
		return nil, false
	}
	return out, true
}

func allFloatings(args []types.Expr) ([]types.Floating, bool) {
	var out []types.Floating
	for _, arg := range args {
		if arg, ok := arg.(types.Floating); ok {
			out = append(out, arg)
			continue
		}
		return nil, false
	}
	return out, true
}

func allNumbers(args []types.Expr) ([]types.Floating, bool) {
	var out []types.Floating
	for _, arg := range args {
		switch arg := arg.(type) {
		case types.Floating:
			out = append(out, arg)
		case types.Integer:
			out = append(out, types.Floating(arg))
		default:
			return nil, false
		}
	}
	return out, true
}

func allStrings(args []types.Expr) ([]types.String, bool) {
	var out []types.String
	for _, arg := range args {
		if arg, ok := arg.(types.String); ok {
			out = append(out, arg)
			continue
		}
		return nil, false
	}
	return out, true
}
