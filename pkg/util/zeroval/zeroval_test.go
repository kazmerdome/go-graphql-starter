package zeroval_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/util/zeroval"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZeroval(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zeroval Test Suite")
}

var _ = Describe("Zeroval Test Suite", func() {
	type testdata struct{}

	When("it is zeroval", func() {
		refType := testdata{}
		valType := ""

		It("should validate the data and pass", func() {
			Expect(refType).To(BeZero())
			Expect(valType).To(BeZero())

			Expect(zeroval.IsZeroVal(refType)).To(BeTrue())
			Expect(zeroval.IsZeroVal(valType)).To(BeTrue())
		})
	})

	When("it is not zeroval", func() {
		refType := new(testdata)
		valType := "!"

		It("should validate the data and pass", func() {
			Expect(refType).ToNot(BeZero())
			Expect(valType).ToNot(BeZero())

			Expect(zeroval.IsZeroVal(refType)).To(BeFalse())
			Expect(zeroval.IsZeroVal(valType)).To(BeFalse())
		})
	})
})
