package licence_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLicenceModule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Module: Licence")
}
