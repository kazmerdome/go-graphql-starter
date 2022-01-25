package guard_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/kazmerdome/go-graphql-starter/mocks"
	"github.com/kazmerdome/go-graphql-starter/pkg/auth/authorization/guard"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestLicenceGuard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authorization Guards - Licence Guard Suite")
}

type licenceGuardFixture struct {
	guard guard.LicenceGuard
	mocks struct {
		*mocks.LicenceRepository
	}
	data struct {
		licence   licence.Licence
		jwtSecret string
	}
}

func newLicenceGuardFixture() *licenceGuardFixture {
	f := new(licenceGuardFixture)

	// data
	f.data.jwtSecret = "testJwtSecret"
	f.data.licence = licence.Licence{
		ID: primitive.NewObjectID(),
		Grants: []licence.Grant{
			{
				Feature:     licence.POST,
				Version:     "1",
				Permissions: []licence.Permission{licence.READ},
			},
		},
		UsedAt:    nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// mocks
	f.mocks.LicenceRepository = &mocks.LicenceRepository{}

	// setup
	c := config.NewConfigService(config.MODE_GLOBALENV)
	c.Set("JWT_SESSION_SECRET", f.data.jwtSecret)
	l := logger.NewStandardLogger()
	s := shared.NewSharedService(l, c)
	f.guard = guard.NewLicenceGuard(*s, f.mocks.LicenceRepository)

	return f
}

var _ = Describe("Licence Guard Suite", func() {
	/*
	 * Without Bearer Token
	 */
	Context("Testing Guards Without BearerToken", func() {
		var f *licenceGuardFixture
		var testFeature licence.Feature = "post"

		BeforeEach(func() {
			f = newLicenceGuardFixture()
		})

		AfterEach(func() {
			f = nil
		})

		Context("using AuthGuard", func() {
			When("try to get some public resource", func() {
				It("should pass (return with nil error)", func() {
					Expect(f.guard.AuthGuard(testFeature, []licence.Permission{"read"}, "")).To(BeNil())
				})
			})
			When("try to get restricted resource", func() {
				It("should throw an unauthorized error", func() {
					err := f.guard.AuthGuard(testFeature, []licence.Permission{"create"}, "")
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("unauthorized"))
				})
			})
		})

		Context("using GetPermissionsGuard", func() {
			When("call GetPermissionsGuard", func() {
				It("should provide the the default (hardcoded) visitor licence", func() {
					permissions, err := f.guard.GetPermissionsGuard("", testFeature)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("access denied"))
					Expect(permissions).To(BeEmpty())
				})
			})
		})

		Context("using GetIdGuard", func() {
			When("try to get id with empty bearer token", func() {
				It("should throw an access denied error", func() {
					_, err := f.guard.GetIdGuard("")
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("access denied"))
				})
			})
		})
	})

	/*
	 * With Valid Bearer Token
	 */
	Context("Testing Guards With Valid BearerToken and Valid LicenceID in it", func() {
		var m *licenceGuardFixture
		var bearerToken string
		var testFeature licence.Feature = "post"

		BeforeEach(func() {
			m = newLicenceGuardFixture()
			licenceToken, err := token.GenerateJWTToken(map[string]string{"lid": m.data.licence.ID.Hex()}, m.data.jwtSecret, 100)
			bearerToken = fmt.Sprintf("Bearer %s", licenceToken)
			Expect(err).To(BeNil())
			m.mocks.LicenceRepository.On("One", mock.Anything).Return(&m.data.licence, nil)
		})

		AfterEach(func() {
			m = nil
			bearerToken = ""
		})

		Context("using AuthGuard", func() {
			When("try to get public resource with a token that has access to it", func() {
				It("should be ok", func() {
					err := m.guard.AuthGuard(testFeature, []licence.Permission{"read"}, bearerToken)
					Expect(err).To(BeNil())
				})
			})
		})

		Context("using GetPermissionsGuard", func() {
			When("try to call GetPermissionsGuard", func() {
				It("should return the fakeLicence post permissions", func() {
					permissions, err := m.guard.GetPermissionsGuard(bearerToken, testFeature)
					var fakePermissions []licence.Permission
					for _, g := range m.data.licence.Grants {
						if g.Feature == testFeature {
							fakePermissions = g.Permissions
						}
					}
					Expect(err).To(BeNil())
					Expect(permissions).To(Equal(fakePermissions))
				})
			})
		})

		Context("using GetIdGuard", func() {
			When("try to get id with bearer token", func() {
				It("should return with fakeLicence.Id", func() {
					oid, err := m.guard.GetIdGuard(bearerToken)
					Expect(err).To(BeNil())
					Expect(oid).To(Equal(m.data.licence.ID))
				})
			})
		})
	})

	/*
	 * With Invalid Bearer Token
	 */
	Context("Testing Guards With Invalid BearerToken", func() {
		var m *licenceGuardFixture
		var err error
		var testFeature licence.Feature = "post"
		var licenceToken string

		BeforeEach(func() {
			m = newLicenceGuardFixture()
			licenceToken, err = token.GenerateJWTToken(map[string]string{"lid": m.data.licence.ID.Hex()}, m.data.jwtSecret, 100)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			m = nil
			err = nil
			licenceToken = ""
		})

		Context("Using AuthGuard", func() {
			When("try to get resource with an invalid token [with two Bearer words]", func() {
				It("should throw an error", func() {
					invalidToken := fmt.Sprintf("Bearer Bearer %s", licenceToken)
					err := m.guard.AuthGuard(testFeature, []licence.Permission{"read"}, invalidToken)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("token contains an invalid number of segments"))
				})
			})
			When("try to get resource with an invalid token [without Bearer word]", func() {
				It("should throw an error", func() {
					invalidToken := licenceToken
					err := m.guard.AuthGuard(testFeature, []licence.Permission{"read"}, invalidToken)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("invalid header token"))
				})
			})
		})

		Context("Using GetPermissionsGuard", func() {
			When("call GetPermissionsGuard with an invalid token", func() {
				It("should throw an error and returns with empty licence", func() {
					invalidToken := licenceToken
					permissions, err := m.guard.GetPermissionsGuard(invalidToken, testFeature)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("invalid header token"))
					Expect(permissions).To(BeEmpty())
				})
			})
		})

		Context("Using GetIdGuard", func() {
			When("try to get id with invalid bearer token", func() {
				It("should throw an access denied error", func() {
					_, err = m.guard.GetIdGuard("")
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("access denied"))
				})
			})
		})
	})

	/*
	 * With Valid Bearer Token And Deleted LicenceID
	 */
	Context("Testing Guards With Valid BearerToken And Deleted Licence", func() {
		var f *licenceGuardFixture
		var err error
		var testFeature licence.Feature = "post"
		var bearerToken string
		var mimicId primitive.ObjectID

		BeforeEach(func() {
			f = newLicenceGuardFixture()
			// Create a valid token with a "deleted" licence id (a valid id but is an already deleted one)
			mimicId = primitive.NewObjectID()
			var licenceToken string
			licenceToken, err = token.GenerateJWTToken(map[string]string{"lid": mimicId.Hex()}, f.data.jwtSecret, 100)
			Expect(err).To(BeNil())
			bearerToken = fmt.Sprintf("Bearer %s", licenceToken)
		})

		AfterEach(func() {
			f = nil
			err = nil
			bearerToken = ""
			mimicId = primitive.ObjectID{}
		})

		Context("Using AuthGuard", func() {
			When("try to get resource with a valid token but with a deleted licence id", func() {
				It("should return with an error", func() {
					f.mocks.LicenceRepository.On("One", mock.Anything).Return(nil, nil)
					err = f.guard.AuthGuard(testFeature, []licence.Permission{"read"}, bearerToken)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("access denied"))
				})
			})
		})

		Context("Using GetPermissionsGuard", func() {
			When("call GetPermissionsGuard with an valid token but with a deleted licence id", func() {
				It("should throw an error and returns with empty licence", func() {
					f.mocks.LicenceRepository.On("One", mock.Anything).Return(nil, nil)
					permissions, err := f.guard.GetPermissionsGuard(bearerToken, testFeature)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("access denied"))
					Expect(permissions).To(BeEmpty())
				})
			})
		})

		Context("Using GetIdGuard", func() {
			When("Try to get id from a valid bearer token", func() {
				It("should provides that (does not matter is deleted already or not from the repositry)", func() {
					oid, err := f.guard.GetIdGuard(bearerToken)
					Expect(err).To(BeNil())
					Expect(oid.Hex()).To(Equal(mimicId.Hex()))
				})
			})
		})
	})

	/*
	 * With Valid Bearer Token And Invalid LicenceID
	 */
	Context("Testing Guards With Valid BearerToken And Invalid LicenceID", func() {
		var f *licenceGuardFixture
		var err error
		var testFeature licence.Feature = "post"
		var bearerToken string

		BeforeEach(func() {
			f = newLicenceGuardFixture()
			var licenceToken string
			licenceToken, err = token.GenerateJWTToken(map[string]string{"lid": "iamnotavalidobjectidlol"}, f.data.jwtSecret, 100)
			Expect(err).To(BeNil())
			bearerToken = fmt.Sprintf("Bearer %s", licenceToken)
		})

		AfterEach(func() {
			f = nil
			err = nil
			bearerToken = ""
		})

		Context("Using AuthGuard", func() {
			When("try to get resource with a valid token but with an invalid licence id", func() {
				It("should return with an error", func() {
					err = f.guard.AuthGuard(testFeature, []licence.Permission{"read"}, bearerToken)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("unauthorized"))
				})
			})
		})

		Context("Using GetPermissionsGuard", func() {
			When("try to get permissions with a valid token but with an invalid licence id in it", func() {
				It("should throw an error and returns with empty licence", func() {
					permissions, err := f.guard.GetPermissionsGuard(bearerToken, testFeature)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("unauthorized"))
					Expect(permissions).To(BeEmpty())
				})
			})
		})

		Context("Using GetIdGuard", func() {
			When("try to get id with a valid token but with an invalid licence id", func() {
				It("should throw an error", func() {
					_, err = f.guard.GetIdGuard(bearerToken)
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal("unauthorized"))
				})
			})
		})
	})
})
