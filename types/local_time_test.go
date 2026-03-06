package types_test

import (
	"time"

	"github.com/Nuanu-com/go-utils/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LocalTime", func() {
	Describe("String()", func() {
		It("returns formatted time", func() {
			l := types.LocalTime(time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local))

			Expect(l.String()).To(Equal("2025-07-01T00:00:00Z"))
		})

		It("Convert time to local time", func() {
			l := types.LocalTime(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))

			Expect(l.String()).To(Equal("2025-07-01T08:00:00Z"))
		})
	})
})

var _ = Describe("LocalTimeConverter", func() {
	It("parse the date time in local time zone", func() {
		l := types.LocalTimeConverter("2025-07-01T08:00:00Z")

		val, ok := l.Interface().(types.LocalTime)

		Expect(ok).To(BeTrue())
		res := types.LocalTime(time.Date(2025, 7, 1, 8, 0, 0, 0, time.Local))
		Expect(val).To(Equal(res))
	})
})
