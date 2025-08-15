package types_test

import (
	"encoding/json"
	"time"

	"github.com/Nuanu-com/go-utils/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LocalDateTime", func() {
	Describe("UnmarshalJSON", func() {
		It("parses a valid date", func() {
			var data types.LocalDateTime

			err := json.Unmarshal([]byte(`"2025-05-01T08:00:00Z"`), &data)
			Expect(err).To(BeNil())

			Expect(data.Time).To(Equal(time.Date(2025, 5, 1, 8, 0, 0, 0, time.Local)))
		})

		It("returns error when the format invalid", func() {
			var data types.LocalDateTime

			err := json.Unmarshal([]byte(`"2025-05-01TX08:00:00Z"`), &data)
			Expect(err).To(
				MatchError(
					`parsing time "2025-05-01TX08:00:00Z" as "2006-01-02T15:04:05Z": cannot parse "X08:00:00Z" as "15"`,
				),
			)
		})
	})

	Describe("MarshalJSON", func() {
		It("marshal the data", func() {
			data := types.LocalDateTime{
				Time: time.Date(2025, 5, 1, 8, 0, 0, 0, time.Local),
			}

			res, err := json.Marshal(data)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`"2025-05-01T08:00:00Z"`)))
		})
	})

	Describe("MarshalText", func() {
		It("marshal it into string", func() {
			data := types.LocalDateTime{
				Time: time.Date(2025, 5, 1, 8, 0, 0, 0, time.Local),
			}

			res, err := data.MarshalText()
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`2025-05-01T08:00:00Z`)))
		})
	})

	Describe("UnmarshalText", func() {
		It("marshal the correct data", func() {
			var data types.LocalDateTime

			err := data.UnmarshalText([]byte(`2025-05-01T08:00:00Z`))
			Expect(err).To(BeNil())
			Expect(data.Time).To(Equal(time.Date(2025, 5, 1, 8, 0, 0, 0, time.Local)))
		})

		It("returns error when the data invalid", func() {
			var data types.LocalDateTime

			err := data.UnmarshalText([]byte(`2025-05-01TXX08:00:00Z`))
			Expect(err).To(
				MatchError(
					`parsing time "2025-05-01TXX08:00:00Z" as "2006-01-02T15:04:05Z": cannot parse "XX08:00:00Z" as "15"`,
				),
			)
		})
	})
})
