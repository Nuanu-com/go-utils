package numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Numbers Suite")
}
