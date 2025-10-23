package types_test

import (
	"encoding/json"
	"fmt"

	"github.com/Nuanu-com/go-utils/types"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type foo struct {
	data int
}

// MarshalText implements encoding.TextMarshaler.
func (f foo) MarshalText() (text []byte, err error) {
	return []byte("Foo Bar"), nil
}

func (f *foo) UnmarshalText(data []byte) error {
	f.data = 100
	return nil
}

type customUnimplemented struct {
	data int
}

var _ = Describe("Null[T]", func() {
	Describe("JSON Marshaler", func() {
		It("returns the valid []byte", func() {
			data := types.NewNull("Chicken", true)

			res, err := json.Marshal(data)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`"Chicken"`)))

			data2 := types.NewNull(12, true)

			res, err = json.Marshal(data2)

			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`12`)))
		})

		It("returns nil byte when data invalid", func() {
			data := types.NewNull("", false)

			res, err := json.Marshal(data)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`null`)))
		})

		It("handles Date", func() {
			date := types.NewNull[types.Date](types.MustParseDate("2026-04-01"), true)

			res, err := json.Marshal(date)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`"2026-04-01"`)))
		})

		It("handles LocalDateTime", func() {
			date := types.NewNull[types.LocalDateTime](
				types.MustParseLocalDateTime("2025-05-01T08:00:00Z"),
				true,
			)

			res, err := json.Marshal(date)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`"2025-05-01T08:00:00Z"`)))
		})
	})

	Describe("MarshalJSON", func() {
		It("returns the object", func() {
			var result types.Null[int]

			err := json.Unmarshal([]byte(`10`), &result)

			Expect(err).To(BeNil())
			Expect(result.V).To(Equal(10))
			Expect(result.Valid).To(BeTrue())

			var result2 types.Null[string]

			err = json.Unmarshal([]byte(`"Fish & Chip"`), &result2)

			Expect(err).To(BeNil())
			Expect(result2.Valid).To(BeTrue())
			Expect(result2.V).To(Equal("Fish & Chip"))
		})

		It("returns empty when the data is null", func() {
			var result types.Null[int]

			err := json.Unmarshal([]byte(`null`), &result)
			Expect(err).To(BeNil())
			Expect(result.V).To(Equal(0))
			Expect(result.Valid).To(BeFalse())

			var result2 types.Null[string]

			err = json.Unmarshal([]byte(`null`), &result2)
			Expect(err).To(BeNil())
			Expect(result2.V).To(Equal(""))
			Expect(result2.Valid).To(BeFalse())
		})

		It("handles Date", func() {
			var result types.Date

			err := json.Unmarshal([]byte(`"2026-04-01"`), &result)
			Expect(err).To(BeNil())
			Expect(result.Format(types.StandardDateFormat)).To(Equal(`2026-04-01`))
		})

		It("handles LocalDateTime", func() {
			var result types.LocalDateTime

			err := json.Unmarshal([]byte(`"2025-05-01T08:00:00Z"`), &result)
			Expect(err).To(BeNil())
			Expect(result.Format(types.StandardDateTimeFormat)).To(Equal(`2025-05-01T08:00:00Z`))
		})
	})

	Describe("UnmarshalText", func() {
		It("unmarshal primitives", func() {
			var strData types.Null[string]

			err := strData.UnmarshalText([]byte(`Not Today`))

			Expect(err).To(BeNil())
			Expect(strData.V).To(Equal("Not Today"))

			var intData types.Null[int]

			err = intData.UnmarshalText([]byte(`10`))

			Expect(err).To(BeNil())
			Expect(intData.V).To(Equal(10))

			var floatData types.Null[float64]

			err = floatData.UnmarshalText([]byte(`10.89`))

			Expect(err).To(BeNil())
			Expect(floatData.V).To(Equal(10.89))

			var boolData types.Null[bool]

			err = boolData.UnmarshalText([]byte(`true`))

			Expect(err).To(BeNil())
			Expect(boolData.V).To(BeTrue())
		})

		It("handles null value", func() {
			var strData types.Null[string]

			err := strData.UnmarshalText([]byte(`null`))

			Expect(err).To(BeNil())
			Expect(strData.V).To(Equal(""))
			Expect(strData.Valid).To(BeFalse())
		})

		It("handles custom value too", func() {
			var data types.Null[foo]

			err := data.UnmarshalText([]byte("fox"))

			Expect(err).To(BeNil())
			Expect(data.V.data).To(Equal(100))
		})

		It("returns error when custom type does not implement Unmarshaler", func() {
			var data types.Null[customUnimplemented]

			err := data.UnmarshalText([]byte("Fo"))
			Expect(err).To(MatchError("Please implement encoding.TextUnmarshaler for type types_test.customUnimplemented"))
		})

		It("handles UUID", func() {
			var data types.Null[uuid.UUID]
			err := data.UnmarshalText([]byte("df44acc2-9111-4d30-953f-0795f8fe0a50"))
			Expect(err).To(BeNil())
			Expect(data.V.String()).To(Equal("df44acc2-9111-4d30-953f-0795f8fe0a50"))
		})

		It("handles Date", func() {
			var data types.Null[types.Date]
			err := data.UnmarshalText([]byte("2025-05-01"))
			Expect(err).To(BeNil())
			Expect(data.V.Format(types.StandardDateFormat)).To(Equal("2025-05-01"))
		})

		It("handles LocalDateTime", func() {
			var data types.Null[types.LocalDateTime]
			err := data.UnmarshalText([]byte("2025-05-01T08:00:00Z"))
			Expect(err).To(BeNil())
			Expect(data.V.Format(types.StandardDateTimeFormat)).To(Equal("2025-05-01T08:00:00Z"))
		})

		Context("Given a slice", func() {
			It("handles primitive slice", func() {
				var data types.Null[[]int]

				err := data.UnmarshalText([]byte("1,3,4,5,6,7,8"))
				Expect(err).To(BeNil())
				Expect(data.V).To(Equal([]int{1, 3, 4, 5, 6, 7, 8}))
			})

			It("it handles uuid", func() {
				var data types.Null[[]uuid.UUID]

				err := data.UnmarshalText([]byte("df44acc2-9111-4d30-953f-0795f8fe0a50,df44acc2-9111-4d30-953f-0795f8fe0a50"))
				Expect(err).To(BeNil())
				Expect(data.V).To(Equal([]uuid.UUID{
					uuid.MustParse("df44acc2-9111-4d30-953f-0795f8fe0a50"),
					uuid.MustParse("df44acc2-9111-4d30-953f-0795f8fe0a50"),
				}))
			})

			It("handles Date", func() {
				var data types.Null[[]types.Date]
				err := data.UnmarshalText([]byte("2025-05-01,2025-11-01"))
				Expect(err).To(BeNil())
				Expect(data.V[0].Format(types.StandardDateFormat)).To(Equal("2025-05-01"))
				Expect(data.V[1].Format(types.StandardDateFormat)).To(Equal("2025-11-01"))
			})

			It("handles LocalDateTime", func() {
				var data types.Null[[]types.LocalDateTime]
				err := data.UnmarshalText([]byte("2025-05-01T08:00:00Z,2025-05-03T08:00:00Z,2025-05-06T08:00:00Z"))
				Expect(err).To(BeNil())
				Expect(data.V[0].Format(types.StandardDateTimeFormat)).To(Equal("2025-05-01T08:00:00Z"))
				Expect(data.V[1].Format(types.StandardDateTimeFormat)).To(Equal("2025-05-03T08:00:00Z"))
				Expect(data.V[2].Format(types.StandardDateTimeFormat)).To(Equal("2025-05-06T08:00:00Z"))
			})
		})
	})

	Describe("MarshalText", func() {
		It("handles null", func() {
			data := types.NewNull[string]("", false)
			res, err := data.MarshalText()
			Expect(err).To(BeNil())
			Expect(res).To(BeNil())
		})
	})

	It("handles primitive types", func() {
		dataStr := types.NewNull[string]("OI", true)

		resStr, err := dataStr.MarshalText()

		Expect(err).To(BeNil())
		Expect(resStr).To(Equal([]byte("OI")))

		dataBool := types.NewNull[bool](true, true)

		resBool, err := dataBool.MarshalText()

		Expect(err).To(BeNil())
		Expect(resBool).To(Equal([]byte("true")))

		dataInt := types.NewNull[int](10, true)

		resInt, err := dataInt.MarshalText()

		Expect(err).To(BeNil())
		Expect(resInt).To(Equal([]byte("10")))

		dataInt64 := types.NewNull[int64](int64(10), true)
		resInt64, err := dataInt64.MarshalText()

		Expect(err).To(BeNil())
		Expect(resInt64).To(Equal([]byte("10")))
	})

	It("handles custom type", func() {
		data := types.NewNull[foo](foo{}, true)

		res, err := data.MarshalText()
		Expect(err).To(BeNil())
		Expect(res).To(Equal([]byte("Foo Bar")))
	})

	It("returns error when custom type does not implement Marshaler", func() {
		data := types.NewNull[customUnimplemented](customUnimplemented{data: 100}, true)

		_, err := data.MarshalText()

		Expect(err).To(MatchError("Please implement encoding.TextMarshaler for type types_test.customUnimplemented"))
	})

	It("handles Date", func() {
		data := types.NewNull[types.Date](types.MustParseDate("2025-04-01"), true)
		res, err := data.MarshalText()

		Expect(err).To(BeNil())
		Expect(res).To(Equal([]byte("2025-04-01")))
	})

	It("handles LocalDateTime", func() {
		data := types.NewNull[types.LocalDateTime](types.MustParseLocalDateTime("2025-05-01T08:00:00Z"), true)
		res, err := data.MarshalText()

		Expect(err).To(BeNil())
		Expect(res).To(Equal([]byte("2025-05-01T08:00:00Z")))
	})

	Context("Given slices", func() {
		It("handles primitive slice", func() {
			data := types.NewNull[[]int]([]int{1, 3, 4, 5, 6, 7, 8}, true)

			res, err := data.MarshalText()

			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte("1,3,4,5,6,7,8")))
		})

		It("handles uuid slice", func() {
			data := types.NewNull[[]uuid.UUID]([]uuid.UUID{
				uuid.MustParse("278c00cb-c6fe-4eb6-a215-876537a1dda6"),
				uuid.MustParse("568e139f-7d1f-4456-a9e6-f887f26a9a48"),
				uuid.MustParse("153326f5-9db9-4522-ac14-625b282be053"),
			}, true)

			res, err := data.MarshalText()

			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte("278c00cb-c6fe-4eb6-a215-876537a1dda6,568e139f-7d1f-4456-a9e6-f887f26a9a48,153326f5-9db9-4522-ac14-625b282be053")))
		})

		It("handles Date", func() {
			data := types.NewNull[[]types.Date]([]types.Date{
				types.MustParseDate("2025-01-03"),
				types.MustParseDate("2025-04-03"),
				types.MustParseDate("2025-05-03"),
			}, true)

			res, err := data.MarshalText()

			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte("2025-01-03,2025-04-03,2025-05-03")))
		})

		It("handles LocalDateTime", func() {
			data := types.NewNull[[]types.LocalDateTime]([]types.LocalDateTime{
				types.MustParseLocalDateTime("2025-05-01T08:00:00Z"),
				types.MustParseLocalDateTime("2025-05-02T08:00:00Z"),
				types.MustParseLocalDateTime("2025-05-03T08:00:00Z"),
			}, true)

			res, err := data.MarshalText()

			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte("2025-05-01T08:00:00Z,2025-05-02T08:00:00Z,2025-05-03T08:00:00Z")))
		})
	})
})

var _ = Describe("MapNull", func() {
	It("Converts the data when the value is not null", func() {
		data1 := types.MapNull(types.NewNull(1, true), func(i int) types.Null[string] { return types.NewNull(fmt.Sprintf("%d", i), true) })
		Expect(data1).To(Equal(types.NewNull("1", true)))

		data2 := types.MapNull(types.NewNull(1, false), func(i int) types.Null[string] { return types.NewNull(fmt.Sprintf("%d", i), true) })
		Expect(data2).To(Equal(types.NewNull("", false)))
	})
})
