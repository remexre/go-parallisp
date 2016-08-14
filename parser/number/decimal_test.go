package number

import (
	. "github.com/onsi/ginkgo"

	"github.com/remexre/go-parallisp/types"
)

var _ = Describe("Decimal Subparser", func() {
	do(decimalNumber, decimalTests)
})

var decimalTests = []test{
	{"123", types.Integer(123)},
	{"123.", types.Floating(123)},
	{"123.0", types.Floating(123)},
	{"123.45", types.Floating(123.45)},
	{"123.450", types.Floating(123.45)},
	{"123.045", types.Floating(123.045)},
	{"0", types.Integer(0)},
	{"0.", types.Floating(0)},
	{"0.0", types.Floating(0)},
	{"0.45", types.Floating(0.45)},
	{"0.450", types.Floating(0.45)},
	{"0.045", types.Floating(0.045)},
}
