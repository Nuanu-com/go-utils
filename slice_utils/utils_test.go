package slice_utils_test

import (
	"github.com/Nuanu-com/go-utils/slice_utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Slice Utils", Label("Utils"), func() {
	Context("First()", func() {
		It("Returns the first element of the slices", func(ctx SpecContext) {
			three := 3
			Expect(slice_utils.First([]int{3, 5, 6, 7, 7})).To(Equal(&three))
		})
	})

	Context("Map()", func() {
		It("Transforms the slices", func(ctx SpecContext) {
			Expect(slice_utils.Map[int, int]([]int{1, 2, 3, 4, 5}, func(i int) int { return i * 2 })).To(Equal([]int{2, 4, 6, 8, 10}))
		})
	})

	Context("GroupBy()", func() {
		It("Groups the data", func() {
			Expect(slice_utils.GroupBy([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) string {
				if i%2 == 0 {
					return "even"
				} else {
					return "odd"
				}
			})).To(Equal(
				map[string][]int{
					"odd":  []int{1, 3, 5, 7, 9},
					"even": []int{2, 4, 6, 8},
				},
			))
		})
	})

	Context("Reduce", func() {
		It("Reduces slice into final data", func() {
			result := slice_utils.Reduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, func(sum int, current int) int {
				return sum + current
			})

			Expect(result).To(Equal(45))
		})
	})

	Context("Filter", func() {
		It("Filters the data", func() {
			result := slice_utils.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(item int) bool {
				return item%2 == 0
			})

			Expect(result).To(Equal([]int{2, 4, 6, 8, 10}))
		})
	})

	Context("Get", func() {
		It("safely access the slice", func() {
			s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			res, exists := slice_utils.Get(s, 3)
			Expect(exists).To(BeTrue())
			Expect(res).To(Equal(4))

			res, exists = slice_utils.Get(s, 10)

			Expect(exists).To(BeFalse())
			Expect(res).To(Equal(0))

			res, exists = slice_utils.Get(s, 9)
			Expect(exists).To(BeFalse())
			Expect(res).To(Equal(0))

			res, exists = slice_utils.Get(s, 8)
			Expect(exists).To(BeTrue())
			Expect(res).To(Equal(9))
		})
	})
})
