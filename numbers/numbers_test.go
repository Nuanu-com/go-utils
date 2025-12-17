package numbers_test

import (
	"github.com/Nuanu-com/go-utils/numbers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NumberMinusPercent", func() {
	It("returns the valid data", func() {
		res := numbers.NumberMinusPercent(100_000, 10)

		Expect(res).To(Equal(float64(90_000)))

		Expect(numbers.NumberMinusPercent(100_000, 0.5)).To(Equal(float64(99500)))

		Expect(numbers.NumberMinusPercent(100_000, 3.5)).To(Equal(float64(96500)))
	})
})

var _ = Describe("PercentOf", func() {
	It("returns the valid data", func() {
		Expect(numbers.PercentOf(100, 20)).To(Equal(float64(20)))
		Expect(numbers.PercentOf(100, 10)).To(Equal(float64(10)))
		Expect(numbers.PercentOf(50, 10)).To(Equal(float64(5)))
	})
})

var _ = Describe("RevPercentOf", func() {
	It("returns the reverse percent", func() {
		Expect(numbers.RevPercentOf(10, 10)).To(Equal(float64(100)))
		Expect(numbers.RevPercentOf(5, 10)).To(Equal(float64(50)))

		Expect(numbers.RevPercentOf(500, 0.5)).To(Equal(float64(100_000)))
	})
})

var _ = Describe("RoundNumber", func() {
	It("returns rounded number", func() {
		Expect(numbers.RoundToEven(3.2)).To(Equal(3))
		Expect(numbers.RoundToEven(3.5)).To(Equal(4))
	})
})

var _ = Describe("Round", func() {
	It("returns rounded number", func() {
		Expect(numbers.Round(1480.5)).To(Equal(1481))
		Expect(numbers.Round(1480.4)).To(Equal(1480))
		Expect(numbers.Round(1480.6)).To(Equal(1481))
	})
})

var _ = Describe("RoundAtPoint51", func() {
	It("rounds the number when the decimal point is greater than or equal to 0.51", func() {
		Expect(numbers.RoundAtPoint51(439.5)).To(Equal(439))
		Expect(numbers.RoundAtPoint51(439.4)).To(Equal(439))
		Expect(numbers.RoundAtPoint51(439.51)).To(Equal(440))
		Expect(numbers.RoundAtPoint51(439.6)).To(Equal(440))
		Expect(numbers.RoundAtPoint51(439.9)).To(Equal(440))
		Expect(numbers.RoundAtPoint51(439.5099999)).To(Equal(439))
	})
})
