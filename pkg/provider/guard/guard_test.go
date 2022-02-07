package guard_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/guard"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGuardProviderConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider: Guard")
}

var _ = Describe("Guard Provider Suite", func() {
	It("should inherit all of the underlying methods from providerConfig", func() {
		l := logger.NewStandardLogger()
		providerConfig := provider.NewProviderConfig(l, nil)
		guardConfig := guard.NewGuardConfig(providerConfig)
		Expect(guardConfig.GetConfig).ShouldNot(BeNil())
		Expect(guardConfig.GetLogger).ShouldNot(BeNil())
		Expect(guardConfig.GetConfig()).Should(BeNil())
		Expect(guardConfig.GetLogger()).Should(Equal(l))
	})
})
