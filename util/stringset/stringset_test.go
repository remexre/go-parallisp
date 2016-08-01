package stringset_test

import (
	"remexre.xyz/go-parallisp/util/stringset"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringSet", func() {
	var set, set2 stringset.StringSet
	BeforeEach(func() {
		set = stringset.New("a", "b", "c")
	})
	Describe("New", func() {
		It("works", func() {
			Expect(set.Contains("")).To(BeFalse())
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
			Expect(set.Contains("d")).To(BeFalse())
			Expect(set.Length()).To(Equal(3))
		})
	})
	Describe("Difference", func() {
		BeforeEach(func() {
			set2 = set.Difference(stringset.New("c", "d"))
		})
		It("works", func() {
			Expect(set2.Equals(stringset.New("a", "b")))
		})
	})
	Describe("Intersection", func() {
		BeforeEach(func() {
			set2 = set.Intersection(stringset.New("c", "d"))
		})
		It("works", func() {
			Expect(set2.Equals(stringset.New("c")))
		})
	})
	Describe("Symmetric Difference", func() {
		BeforeEach(func() {
			set2 = set.SymDifference(stringset.New("c", "d"))
		})
		It("works", func() {
			Expect(set2.Equals(stringset.New("a", "b", "d")))
		})
	})
	Describe("Union", func() {
		BeforeEach(func() {
			set2 = set.Intersection(stringset.New("c", "d"))
		})
		It("works", func() {
			Expect(set2.Equals(stringset.New("a", "b", "c", "d", "e")))
		})
	})
})
