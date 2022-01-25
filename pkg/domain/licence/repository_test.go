package licence_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/mocks"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestLicencePackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: Licence")
}

type licenceRepositoryFixture struct {
	repository licence.LicenceRepository
	mocks      struct {
		*mocks.MongoCollection
		*mocks.MongoDatabase
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
	f.mocks.MongoDatabase = &mocks.MongoDatabase{}

	// setup
	c := config.NewConfigService(config.MODE_GLOBALENV)
	l := logger.NewStandardLogger()
	s := shared.NewSharedService(l, c)
	f.mocks.MongoDatabase.On("Collection", mock.Anything, mock.Anything).Return(f.mocks.MongoCollection)
	f.repository = licence.NewLicenceRepository(*s, f.mocks.MongoDatabase)

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
