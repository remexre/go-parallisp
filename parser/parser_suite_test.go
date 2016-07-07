package parser_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/remexre/go-parcom"

	"testing"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

type test struct {
	data      string
	expr      interface{}
	remaining string
	ok        bool
}

func do(p parcom.Parser, ts []test) {
	for _, t := range ts {
		func(t test) {
			It(fmt.Sprintf("Parses `%s`", t.data), func() {
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

type simpleTest struct {
	data string
	expr interface{}
}

func doSimple(p parcom.Parser, sts []simpleTest) {
	tests := make([]test, len(sts))
	for i, st := range sts {
		tests[i] = test{st.data, st.expr, "", true}
	}
	do(p, tests)
}
