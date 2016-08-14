package number

import (
	"fmt"

	"github.com/remexre/go-parallisp/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("applySign", func() {
	for _, t := range applySignTests() {
		func(t applySignTest) {
			It(fmt.Sprintf("Works for %d * %#v -> %#v", t.sign, t.in, t.out), func() {
				out := applySign(t.sign, t.in)
				if t.out == nil {
					Expect(out).To(BeNil())
				} else {
					Expect(out).To(Equal(t.out))
				}
			})
		}(t)
	}
})

type applySignTest struct {
	sign int64
	in   types.Expr
	out  types.Expr
}

func applySignTests() []applySignTest {
	var out []applySignTest
	for _, sign := range []int64{1, 0, -1} {
		for _, expr := range []types.Expr{
			nil,
			types.Integer(123),
			types.Floating(123.45),
			types.Cons{nil, nil},
			types.Vector{},
			types.String("Hello String"),
			types.Symbol("hello-symbol"),
		} {
			test := applySignTest{sign, expr, nil}
			if !(sign == 1 || sign == -1) {
				test.out = test.in
			} else if i, ok := expr.(types.Integer); ok {
				test.out = types.Integer(sign) * i
			} else if f, ok := expr.(types.Floating); ok {
				test.out = types.Floating(sign) * f
			} else {
				test.out = test.in
			}
			out = append(out, test)
		}
	}
	return out
}
