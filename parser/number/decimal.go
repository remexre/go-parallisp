package number

import (
	"math"

	"github.com/remexre/go-parcom"

	"remexre.xyz/go-parallisp/types"
)

var decimals = parcom.Map(parcom.AnyOf("0123456789"), func(n string) []byte {
	bs := []byte(n)
	for i, b := range bs {
		bs[i] = b - '0'
	}
	return bs
})

var decimalNumber = parcom.Map(parcom.Chain(
	parcom.Map(decimals, func(nums []uint8) int64 {
		var out int64
		for i, l := 0, len(nums)-1; i <= l; i++ {
			out += int64(nums[l-i]) * int64(math.Pow10(int(i)))
		}
		return out
	}),
	parcom.Opt(parcom.Map(parcom.Chain(
		parcom.Tag("."),
		parcom.Opt(decimals, []uint8{0}),
	), func(i []interface{}) []uint8 {
		return i[1].([]uint8)
	}), nil),
), func(integer int64, floatingDigits []uint8) types.Expr {
	if floatingDigits == nil {
		return types.Integer(integer)
	}
	floating := 0.
	for i, digit := range floatingDigits {
		floating += float64(digit) / math.Pow10(i+1)
	}
	return types.Floating(float64(integer) + floating)
})
