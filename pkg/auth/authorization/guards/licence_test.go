package guards_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kazmerdome/go-graphql-starter/mocks"
	"github.com/kazmerdome/go-graphql-starter/pkg/auth/authorization/guards"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type licenceGuardFixture struct {
	g     guards.LicenceGuard
	mocks struct {
		*mocks.LicenceRepository
	}
	fakes struct {
		fakeLicence   licence.Licence
		testJwtSecret string
	}
}

func newLicenceGuardFixture(t *testing.T) *licenceGuardFixture {
	// fakes
	fakeLicence := licence.Licence{
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
	fakes := struct {
		fakeLicence   licence.Licence
		testJwtSecret string
	}{fakeLicence, "testJwtSecret"}

	// mocks
	mockedLicenceRepository := mocks.LicenceRepository{}
	var m = struct{ *mocks.LicenceRepository }{
		&mockedLicenceRepository,
	}

	// Setup
	c := config.NewConfigService(config.MODE_GLOBALENV)
	c.Set("JWT_SESSION_SECRET", fakes.testJwtSecret)
	l := logger.NewStandardLogger()
	s := *shared.NewSharedService(l, c)
	g := guards.NewLicenceGuard(s, &mockedLicenceRepository)

	return &licenceGuardFixture{g, m, fakes}
}

func TestGuardsWithoutBearerToken(t *testing.T) {
	var testFeature licence.Feature = "post"
	m := newLicenceGuardFixture(t)
	assert := assert.New(t)
	m.mocks.LicenceRepository.On("One", mock.Anything).Return(&m.fakes.fakeLicence, nil)

	// Context: Using AuthGuard
	// When try to get some public resource
	// It should be ok
	assert.Equal(m.g.AuthGuard(testFeature, []licence.Permission{"read"}, ""), nil)
	// When Try to get restricted resource
	// It should throw an unauthorized error
	err := m.g.AuthGuard(testFeature, []licence.Permission{"create"}, "")
	assert.NotNil(err)
	assert.Equal(err.Error(), "unauthorized")

	// Context: Using GetPermissionsGuard
	// When call GetPermissionsGuard
	// It should provide the the hardcoded default visitor licence
	permissions, err := m.g.GetPermissionsGuard("", testFeature)
	assert.NotNil(err)
	assert.Equal(err.Error(), "access denied")
	assert.Empty(permissions)

	// Context: Using GetIdGuard
	// When Try to get id with empty bearer token
	// It should throw an access denied error
	_, err = m.g.GetIdGuard("")
	assert.NotNil(err)
	assert.Equal(err.Error(), "access denied")
}

func TestGuardsWithBearerToken(t *testing.T) {
	var testFeature licence.Feature = "post"
	m := newLicenceGuardFixture(t)
	assert := assert.New(t)
	m.mocks.LicenceRepository.On("One", mock.Anything).Return(&m.fakes.fakeLicence, nil)
	licenceToken, err := token.GenerateJWTToken(m.fakes.fakeLicence.ID.Hex(), m.fakes.testJwtSecret, 100)
	bearerToken := fmt.Sprintf("Bearer %s", licenceToken)
	assert.Nil(err)

	// Context: Using AuthGuard
	// When try to get public resource with a token that has access to it
	// It should be ok
	err = m.g.AuthGuard(testFeature, []licence.Permission{"read"}, bearerToken)
	assert.Equal(err, nil)
	// When the found licence does not have permission to the resource, It should throw an error
	err = m.g.AuthGuard("category", []licence.Permission{"read"}, bearerToken)
	assert.Equal(err.Error(), "unauthorized")

	// Context: Using GetPermissionsGuard
	// When call GetPermissionsGuard
	// It should return with the fakeLicence post permissions
	permissions, err := m.g.GetPermissionsGuard(bearerToken, testFeature)
	var fakePermissions []licence.Permission
	for _, g := range m.fakes.fakeLicence.Grants {
		if g.Feature == testFeature {
			fakePermissions = g.Permissions
		}
	}
	assert.NoError(err)
	assert.Equal(permissions, fakePermissions)

	// Context: Using GetIdGuard
	// When Try to get id with bearer token
	// It should return with fakeLicence.Id
	oid, err := m.g.GetIdGuard(bearerToken)
	assert.Nil(err)
	assert.Equal(oid, m.fakes.fakeLicence.ID)
}

func TestGuardsWithInvalidBearerToken(t *testing.T) {
	var testFeature licence.Feature = "post"
	m := newLicenceGuardFixture(t)
	assert := assert.New(t)
	m.mocks.LicenceRepository.On("One", mock.Anything).Return(&m.fakes.fakeLicence, nil)
	licenceToken, err := token.GenerateJWTToken(m.fakes.fakeLicence.ID.Hex(), m.fakes.testJwtSecret, 100)
	assert.Nil(err)

	// Context: Using AuthGuard
	// When try to get resource with an invalid token [with two Bearer words]
	// It should throw an error
	invalidToken := fmt.Sprintf("Bearer Bearer %s", licenceToken)
	err = m.g.AuthGuard(testFeature, []licence.Permission{"read"}, invalidToken)
	assert.NotNil(err)
	assert.Equal(err.Error(), "token contains an invalid number of segments")
	// When try to get resource with an invalid token [without Bearer word]
	// It should throw an error
	invalidToken = licenceToken
	err = m.g.AuthGuard(testFeature, []licence.Permission{"read"}, invalidToken)
	assert.NotNil(err)
	assert.Equal(err.Error(), "invalid header token")

	// Context: Using GetPermissionsGuard
	// When call GetPermissionsGuard with an invalid token
	// It should returns with an empty hardcoded default visitor licence
	permissions, err := m.g.GetPermissionsGuard(invalidToken, testFeature)
	assert.NotNil(err)
	assert.Equal(err.Error(), "invalid header token")
	assert.Empty(permissions)

	// Context: Using GetIdGuard
	// When Try to get id with invalid bearer token
	// It should throw an access denied error
	_, err = m.g.GetIdGuard("")
	assert.NotNil(err)
	assert.Equal(err.Error(), "access denied")
}

func TestGuardsWithValidBearerTokenAndNotExistenceLicenceId(t *testing.T) {
	var testFeature licence.Feature = "post"
	m := newLicenceGuardFixture(t)
	assert := assert.New(t)
	m.mocks.LicenceRepository.On("One", mock.Anything).Return(nil, nil)
	// Create a valid token with a "deleted" licence id (a valid id but mimics that is an already deleted one)
	mimicId := primitive.NewObjectID().Hex()
	licenceToken, err := token.GenerateJWTToken(mimicId, m.fakes.testJwtSecret, 100)
	bearerToken := fmt.Sprintf("Bearer %s", licenceToken)
	assert.Nil(err)

	// Context: Using AuthGuard
	// When try to get resource with a valid token but with a deleted licence id
	// It should return an error
	err = m.g.AuthGuard(testFeature, []licence.Permission{"read"}, bearerToken)
	assert.NotNil(err)
	assert.Equal(err.Error(), "access denied")

	// Context: Using GetPermissionsGuard
	// When call GetPermissionsGuard with an invalid token
	// It should returns with an empty hardcoded default visitor licence
	permissions, err := m.g.GetPermissionsGuard(bearerToken, testFeature)
	assert.NotNil(err)
	assert.Equal(err.Error(), "access denied")
	assert.Empty(permissions)

	// Context: Using GetIdGuard
	// When Try to get id from a valid bearer token
	// It should provides that (does not matter is deleted already or not from the repositry)
	oid, err := m.g.GetIdGuard(bearerToken)
	assert.Nil(err)
	assert.Equal(oid.Hex(), mimicId)
}

func TestGuardsWithValidBearerTokenAndInvalidLicenceId(t *testing.T) {
	var testFeature licence.Feature = "post"
	m := newLicenceGuardFixture(t)
	assert := assert.New(t)
	// Create a valid token with an invalid objectID
	licenceToken, err := token.GenerateJWTToken("iamnotavalidobjectidlol", m.fakes.testJwtSecret, 100)
	bearerToken := fmt.Sprintf("Bearer %s", licenceToken)
	assert.Nil(err)

	// Context: Using AuthGuard
	// When try to get resource with a valid token but with an invalid licence id
	// It should return an error
	err = m.g.AuthGuard(testFeature, []licence.Permission{"read"}, bearerToken)
	assert.NotNil(err)
	assert.Equal(err.Error(), "unauthorized")

	// Context: Using GetIdGuard
	// When try to get id with a valid token but with an invalid licence id
	// It should return an error
	_, err = m.g.GetIdGuard(bearerToken)
	assert.NotNil(err)
	assert.Equal(err.Error(), "unauthorized")
}
