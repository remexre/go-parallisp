package natives

import (
	"math"

	"remexre.xyz/go-parallisp/types"
)

// Eq tests for equality.
func Eq(a, b types.Expr) types.Expr {
	if a == nil || b == nil {
		if a == b {
			return types.Symbol("t")
		}
		return nil
	}

	switch a := a.(type) {
	case types.Floating:
		if b, ok := b.(types.Floating); ok && math.Abs(float64(a-b)) < 0.00000001 {
			return types.Symbol("t")
		}
	case types.Integer:
		if b, ok := b.(types.Integer); ok && a == b {
			return types.Symbol("t")
		}
	case types.String:
		if b, ok := b.(types.String); ok && a == b {
			return types.Symbol("t")
		}
	case types.Symbol:
		if b, ok := b.(types.Symbol); ok && a == b {
			return types.Symbol("t")
		}
	case types.Cons:
		if b, ok := b.(types.Cons); ok {
			return eqCons(a, b)
		}
	case types.Vector:
		if b, ok := b.(types.Vector); ok {
			return eqVector(a, b)
		}
	}
	return nil
}

func eqCons(a, b types.Cons) types.Expr {
	if Eq(a.Car(), b.Car()) != nil {
		return Eq(a.Cdr(), b.Cdr())
	}
	return nil
}

func eqVector(a, b types.Vector) types.Expr {
	if len(a) != len(b) {
		return nil
	}
	for i, ai := range a {
		if Eq(ai, b[i]) == nil {
			return nil
		}
	}
	return types.Symbol("t")
}
