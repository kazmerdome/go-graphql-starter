package validator_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/util/validator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator Test Suite")
}

var _ = Describe("Validator Test Suite", func() {
	type test struct {
		Name string `json:"name" bson:"name" validate:"required"`
	}

	When("calling validator", func() {
		It("should validate the data and pass", func() {
			t := new(test)
			t.Name = "test"
			err := validator.Validate(t)
			Expect(err).To(BeNil())
		})
		It("should validate the data and fail | the validator needs to be loaded from cache", func() {
			t := new(test)
			err := validator.Validate(t)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("Key: 'test.Name' Error:Field validation for 'Name' failed on the 'required' tag"))
		})
	})
})
