package types_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"remexre.xyz/go-parallisp/types"
)

var _ = Describe("NewConsList", func() {
	It("Works", func() {
		Expect(types.NewConsList(
			types.Integer(1),
			types.Integer(2),
			types.Integer(3),
		)).To(Equal(types.NewCons(
			types.Integer(1),
			types.NewCons(
				types.Integer(2),
				types.NewCons(
					types.Integer(3),
					nil,
				),
			),
		)))
	})
})

var _ = Describe("NewImproperConsList", func() {
	It("Works", func() {
		Expect(types.NewImproperConsList(
			types.Integer(1),
			types.Integer(2),
			types.Integer(3),
		)).To(Equal(types.NewCons(
			types.Integer(1),
			types.NewCons(
				types.Integer(2),
				types.Integer(3),
			),
		)))
	})
})
