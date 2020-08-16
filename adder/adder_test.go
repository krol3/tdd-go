package adder_test

import (
	. "github.com/krol3/tdd-go.git/adder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adder", func() {

	Describe("Add", func() {

		Context("when summands are positive", func() {
			It("adds two numbers", func() {
				sum, err := Add(2, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(sum).To(Equal(5))
			})
		})

		Context("when summands are positive", func() {

			It("adds two numbers", func() {
				sum, err := Add(2, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(sum).To(Equal(5))
			})

		})

		Context("when summand is negative", func() {

			It("returns an err", func() {
				_, err := Add(-1, -1)
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
