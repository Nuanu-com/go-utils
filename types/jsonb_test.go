package types_test

import (
	"encoding/json"

	"github.com/Nuanu-com/go-utils/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSONB", func() {
	Describe("MarshalJSON", func() {
		It("returns the bytes", func() {
			data := types.JSONB{
				"name": "George",
			}

			res, err := json.Marshal(data)

			Expect(err).To(BeNil())
			Expect(res).To(Equal([]byte(`{"name":"George"}`)))
		})
	})

	Describe("UnmarshalJSON", func() {
		It("converts json string into JSONB object", func() {
			var obj types.JSONB

			err := json.Unmarshal([]byte(`{"name":"George"}`), &obj)

			Expect(err).To(BeNil())
			Expect(obj).To(Equal(types.JSONB{"name": "George"}))
		})
	})

	Describe("ToJSONB", func() {
		It("converts any thing into JSONB", func() {
			aa := map[string]any{
				"foo": "bar",
			}

			bb, err := types.ToJSONB(aa)

			Expect(err).To(BeNil())
			Expect(bb).To(Equal(types.JSONB{"foo": "bar"}))
		})
	})
})
