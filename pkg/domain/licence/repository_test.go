package licence_test

import (
	"github.com/kazmerdome/go-graphql-starter/mocks"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type licenceRepositoryFixture struct {
	repository licence.LicenceRepository
	mocks      struct {
		*mocks.MongoCollection
		*mocks.MongodbAdapter
	}
	data struct {
		filter   *licence.LicenceWhereDTO
		licences []*licence.Licence
	}
}

func newLicenceRepositoryFixture() *licenceRepositoryFixture {
	f := new(licenceRepositoryFixture)

	// data
	id := primitive.NewObjectID()
	f.data.filter = &licence.LicenceWhereDTO{
		ID:     &id,
		Grants: []*licence.Grant{},
	}
	f.data.licences = []*licence.Licence{
		{
			ID: id,
		},
	}

	// mocks
	f.mocks.MongoCollection = &mocks.MongoCollection{}
	f.mocks.MongodbAdapter = &mocks.MongodbAdapter{}
	f.mocks.MongodbAdapter.On("Collection", mock.Anything, mock.Anything).Return(f.mocks.MongoCollection)

	// setup
	c := config.NewConfig(config.MODE_GLOBALENV)
	l := logger.NewStandardLogger()
	moduleConfig := module.NewModuleConfig(l, c)

	module := licence.NewLicenceModule(moduleConfig, f.mocks.MongodbAdapter)

	f.repository = module.GetRepository()

	return f
}

var _ = Describe("Licence Repository Suite", func() {
	var f *licenceRepositoryFixture

	BeforeEach(func() {
		f = newLicenceRepositoryFixture()
	})

	AfterEach(func() {
		f = nil
	})

	/*
	 * One
	 */
	Context("Testing One Method", func() {
		When("try to get a licence with filter", func() {
			It("should call the underlaying FindOne collection method", func() {
				f.mocks.MongoCollection.On("FindOne", mock.Anything, &f.data.filter).Return(&mongo.SingleResult{})
				_, err := f.repository.One(f.data.filter)
				Expect(err).To(BeNil())
			})
		})
		When("try to get a licence without filter", func() {
			It("it should call the underlaying FindOne collection method", func() {
				f.mocks.MongoCollection.On("FindOne", mock.Anything, mock.Anything).Return(&mongo.SingleResult{})
				_, err := f.repository.One(nil)
				Expect(err).To(BeNil())
			})
		})
	})

	/*
	 * List
	 */
	Context("Testing List Method", func() {
		When("try to get a licences with filters", func() {
			It("should call the underlaying Find collection method", func() {
			})
		})
	})
})
