package types_test

import (
	"encoding/json"
	"time"

	"github.com/Nuanu-com/go-utils/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date", func() {
	Describe("UnmarshalJSON", func() {
		It("parse a valid data", func() {
			var date types.Date

			err := json.Unmarshal([]byte(`"2025-05-01"`), &date)
			Expect(err).To(BeNil())
			Expect(date.Time.Format(types.StandardDateFormat)).To(Equal("2025-05-01"))
		})

		It("returns error when given invalid data", func() {
			var date types.Date

			err := json.Unmarshal([]byte(`"2025-05-01T"`), &date)
			Expect(err).To(MatchError(`parsing time "2025-05-01T": extra text: "T"`))
		})
	})

	Describe("MarshalJSON", func() {
		It("marshal into a valid string", func() {
			date := types.Date{Time: time.Date(2026, 8, 19, 19, 10, 0, 0, time.Local)}

			res, err := json.Marshal(date)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`"2026-08-19"`)))
		})
	})

	Describe("MarshalText", func() {
		It("marshal into valid string", func() {
			date := types.Date{Time: time.Date(2026, 8, 19, 19, 10, 0, 0, time.Local)}

			res, err := date.MarshalText()
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`2026-08-19`)))
		})
	})

	Describe("UnmarshalText", func() {
		It("unmarshal string into date", func() {
			var date types.Date

			err := date.UnmarshalText([]byte(`2026-08-19`))

			Expect(err).To(BeNil())
		})

		It("returns error when the string invalid", func() {
			var date types.Date

			err := date.UnmarshalText([]byte(`2026-08-19T`))

			Expect(err).To(MatchError(`parsing time "2026-08-19T": extra text: "T"`))
		})
	})
})
