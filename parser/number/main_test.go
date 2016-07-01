package number

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Number", func() {
	testList, err := mainTests()
	It("Generates Tests", func() {
		Expect(err).To(BeNil())
	})
	for name, tests := range testList {
		Describe(fmt.Sprintf("Parses %s numbers", name), func() {
			do(Parse, tests)
		})
	}
})

var allTests = map[string][]test{
	"decimal": decimalTests,
}

func mainTests() (map[string][]test, error) {
	out := make(map[string][]test)
	for name, tests := range allTests {
		for _, t := range tests {
			out[name] = append(out[name], t)
			out[name] = append(out[name], test{
				"+" + t.data,
				t.expr,
			})
			out[name] = append(out[name], test{
				"-" + t.data,
				applySign(-1, t.expr),
			})
		}
	}
	return out, nil
}
