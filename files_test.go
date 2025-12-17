package go_utils_test

import (
	"context"

	go_utils "github.com/Nuanu-com/go-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileService", func() {
	fileService := go_utils.NewFileService()

	Describe("Upload", func() {
		It("Saves the data", func() {
			filename, err := fileService.Upload(
				context.Background(),
				nil,
				"foo.jpeg",
				"image/jpeg",
			)

			Expect(err).To(BeNil())
			Expect(filename).To(Equal("foo.jpeg"))
		})
	})
})
