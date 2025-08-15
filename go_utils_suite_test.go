package go_utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoUtils Suite")
}
