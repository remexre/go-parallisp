package parallisp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParallisp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parallisp Suite")
}
