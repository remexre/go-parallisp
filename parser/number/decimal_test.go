package number

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/parallisp"
)

var _ = Describe("Decimal Subparser", func() {
	do(decimalNumber, decimalTests)
})

var decimalTests = []test{
	{"123", parallisp.Integer(123)},
	{"123.", parallisp.Floating(123)},
	{"123.0", parallisp.Floating(123)},
	{"123.45", parallisp.Floating(123.45)},
	{"123.450", parallisp.Floating(123.45)},
	{"123.045", parallisp.Floating(123.045)},
	{"0", parallisp.Integer(0)},
	{"0.", parallisp.Floating(0)},
	{"0.0", parallisp.Floating(0)},
	{"0.45", parallisp.Floating(0.45)},
	{"0.450", parallisp.Floating(0.45)},
	{"0.045", parallisp.Floating(0.045)},
}
