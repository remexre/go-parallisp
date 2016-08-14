package parser_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/remexre/go-parallisp/parser"
)

var _ = Describe("Parse", func() {
	Describe("Simple tests", func() {
		for _, test := range simpleTests {
			func(test simpleTest) {
				It(fmt.Sprintf("Parses `%s`", test.data), func() {
					out, err := parser.Parse(test.data)
					Expect(err).ToNot(HaveOccurred())
					Expect(out).To(HaveLen(1))
					if test.expr == nil {
						Expect(out[0]).To(BeNil())
					} else {
						Expect(out[0]).To(Equal(test.expr))
					}
				})
			}(test)
		}
	})
})
