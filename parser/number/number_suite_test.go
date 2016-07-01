package number

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"remexre.xyz/go-parcom"
	"remexre.xyz/parallisp"

	"testing"
)

func TestNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Number Suite")
}

type test struct {
	data string
	expr parallisp.Expr
}

func do(p parcom.Parser, ts []test) {
	for _, t := range ts {
		func(t test) {
			It(fmt.Sprintf(`Parses "%s"`, t.data), func() {
				remaining, out, ok := p(t.data)
				Expect(ok).To(BeTrue())
				Expect(out).To(BeNumerically("~", t.expr))
				Expect(remaining).To(BeEmpty())
			})
		}(t)
	}
}
