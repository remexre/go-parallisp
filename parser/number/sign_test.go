package number

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"remexre.xyz/parallisp"
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
	in   parallisp.Expr
	out  parallisp.Expr
}

func applySignTests() []applySignTest {
	var out []applySignTest
	for _, sign := range []int64{1, 0, -1} {
		for _, expr := range []parallisp.Expr{
			nil,
			parallisp.Integer(123),
			parallisp.Floating(123.45),
			parallisp.Cons{nil, nil},
			parallisp.Vector{},
			parallisp.String("Hello String"),
			parallisp.Symbol("hello-symbol"),
		} {
			test := applySignTest{sign, expr, nil}
			if !(sign == 1 || sign == -1) {
				test.out = test.in
			} else if i, ok := expr.(parallisp.Integer); ok {
				test.out = parallisp.Integer(sign) * i
			} else if f, ok := expr.(parallisp.Floating); ok {
				test.out = parallisp.Floating(sign) * f
			} else {
				test.out = test.in
			}
			out = append(out, test)
		}
	}
	return out
}
