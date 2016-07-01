package parser_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp"

	"testing"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

type test struct {
	data      string
	expr      parallisp.Expr
	remaining string
	ok        bool
}

func do(p parcom.Parser, ts []test) {
	for _, t := range ts {
		func(t test) {
			It(fmt.Sprintf(`Parses "%s"`, t.data), func() {
				remaining, out, ok := p(t.data)
				if t.ok {
					Expect(ok).To(BeTrue())
					if t.expr == nil {
						Expect(out).To(BeNil())
					} else {
						Expect(out).To(Equal(t.expr))
					}
					Expect(remaining).To(Equal(t.remaining))
				} else {
					Expect(ok).To(BeFalse())
				}
			})
		}(t)
	}
}
