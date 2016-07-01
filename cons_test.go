package parallisp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"remexre.xyz/parallisp"
)

var _ = Describe("NewConsList", func() {
	It("Works", func() {
		Expect(parallisp.NewConsList(
			parallisp.Integer(1),
			parallisp.Integer(2),
			parallisp.Integer(3),
		)).To(Equal(parallisp.Cons{
			parallisp.Integer(1),
			parallisp.Cons{
				parallisp.Integer(2),
				parallisp.Cons{
					parallisp.Integer(3),
					nil,
				},
			},
		}))
	})
})

var _ = Describe("NewImproperConsList", func() {
	It("Works", func() {
		Expect(parallisp.NewImproperConsList(
			parallisp.Integer(1),
			parallisp.Integer(2),
			parallisp.Integer(3),
		)).To(Equal(parallisp.Cons{
			parallisp.Integer(1),
			parallisp.Cons{
				parallisp.Integer(2),
				parallisp.Integer(3),
			},
		}))
	})
})
