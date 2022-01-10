package licence_test

import (
	"context"
	"testing"
	"time"

	"github.com/kazmerdome/go-graphql-starter/mocks"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type serviceFixture struct {
	service licence.LicenceService
	mocks   struct{ *mocks.LicenceRepository }
	fakes   struct {
		licences []*licence.Licence
	}
}

func newServiceFixture(t *testing.T) *serviceFixture {
	// fakes
	var createFakeLicence func() *licence.Licence = func() *licence.Licence {
		return &licence.Licence{
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
	}
	fakes := struct{ licences []*licence.Licence }{[]*licence.Licence{
		createFakeLicence(),
		createFakeLicence(),
		createFakeLicence(),
	}}

	// mocks
	mockedLicenceRepository := mocks.LicenceRepository{}
	var m = struct{ *mocks.LicenceRepository }{
		&mockedLicenceRepository,
	}

	// setup
	c := config.NewConfigService(config.MODE_GLOBALENV)
	l := logger.NewStandardLogger()
	ls := licence.NewLicenceService(*shared.NewSharedService(l, c), &mockedLicenceRepository)

	return &serviceFixture{ls, m, fakes}
}

// Context: Using Licence
func TestLicence(t *testing.T) {
	f := newServiceFixture(t)
	assert := assert.New(t)
	ctx := context.TODO()

	// When try to get a licence that does exist
	// It shoud return with the mocked licence
	w := licence.LicenceWhereDTO{ID: &f.fakes.licences[0].ID}
	f.mocks.LicenceRepository.On("One", &w).Return(f.fakes.licences[0], nil).Times(1)
	l, err := f.service.Licence(ctx, &w, nil)
	assert.NoError(err)
	assert.Equal(l.ID, f.fakes.licences[0].ID)

	// When try to get a licence with search
	// It shoud return with the mocked licences and extended where filter
	f = newServiceFixture(t)
	cq := "customQuery"
	w = licence.LicenceWhereDTO{}
	f.mocks.LicenceRepository.On("One", &w).Return(f.fakes.licences[0], nil).Times(1)
	l, err = f.service.Licence(ctx, &w, &cq)
	assert.Equal(w.OR[0], primitive.M(primitive.M{"id": primitive.M{"$regex": cq}}))
	assert.NoError(err)
	assert.Equal(l.ID, f.fakes.licences[0].ID)
}

// Context: Using Licences
func TestLicences(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	w := licence.LicenceWhereDTO{}

	// When try to get a licences that do exist
	// It shoud return with the mocked licences
	f := newServiceFixture(t)
	f.mocks.LicenceRepository.On("List", &w, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(f.fakes.licences, nil).Times(1)
	l, err := f.service.Licences(ctx, &w, nil, nil, nil, nil)
	assert.NoError(err)
	for i := range f.fakes.licences {
		assert.Equal(l[i].ID, f.fakes.licences[i].ID)
	}

	// When try to get a licences with search
	// It shoud return with the mocked licences and extended where filter
	f = newServiceFixture(t)
	cq := "customQuery"
	f.mocks.LicenceRepository.On("List", &w, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(f.fakes.licences, nil).Times(1)
	l, err = f.service.Licences(ctx, &w, nil, nil, nil, &cq)
	assert.NoError(err)
	assert.Equal(w.OR[0], primitive.M(primitive.M{"id": primitive.M{"$regex": cq}}))
	for i := range f.fakes.licences {
		assert.Equal(l[i].ID, f.fakes.licences[i].ID)
	}
}

// Context: Using LicenceCount
func TestLicenceCount(t *testing.T) {
	f := newServiceFixture(t)
	assert := assert.New(t)
	ctx := context.TODO()
	w := licence.LicenceWhereDTO{}
	length := len(f.fakes.licences)

	// When try to get a licenceCount
	// It shoud return with the length of the mocked licences
	f.mocks.LicenceRepository.On("Count", &w).Return(&length, nil).Times(1)
	l, err := f.service.LicenceCount(ctx, &w, nil)
	assert.NoError(err)
	assert.Equal(l, &length)

	// When try to get a licenceCount with search
	// It shoud return with the extended where filter
	cq := "customQuery"
	f.mocks.LicenceRepository.On("Count", &w).Return(&length, nil).Times(1)
	_, err = f.service.LicenceCount(ctx, &w, &cq)
	assert.NoError(err)
	assert.Equal(w.OR[0], primitive.M(primitive.M{"id": primitive.M{"$regex": cq}}))
}

// Context: Using CreateLicence
func TestCreateLicence(t *testing.T) {
	f := newServiceFixture(t)
	assert := assert.New(t)
	ctx := context.TODO()
	newItem := licence.LicenceCreateDTO{
		Grants:    f.fakes.licences[0].Grants,
		UsedAt:    f.fakes.licences[0].UsedAt,
		UpdatedAt: f.fakes.licences[0].UpdatedAt,
		CreatedAt: f.fakes.licences[0].CreatedAt,
	}

	// When try to create a licence
	// It shoud create it so should returns with mocked licence
	f.mocks.LicenceRepository.On("Create", &newItem).Return(f.fakes.licences[0], nil).Times(1)
	l, err := f.service.CreateLicence(ctx, newItem)
	assert.NoError(err)
	assert.Equal(l, f.fakes.licences[0])
}

// Context: Using UpdateLicence
func TestUpdateLicence(t *testing.T) {
	f := newServiceFixture(t)
	assert := assert.New(t)
	ctx := context.TODO()
	where := licence.LicenceWhereUniqueDTO{
		ID: f.fakes.licences[0].ID,
	}
	udpateDTO := licence.LicenceUpdateDTO{}

	// When try to update a licence
	// It shoud call the repository with the correct updateDTO
	f.mocks.LicenceRepository.On("Update", where.ID, &udpateDTO).Return(nil, nil).Times(1)
	_, err := f.service.UpdateLicence(ctx, where, udpateDTO)
	assert.NoError(err)
}

// Context: Using DeleteLicence
func TestDeleteLicence(t *testing.T) {
	f := newServiceFixture(t)
	assert := assert.New(t)
	ctx := context.TODO()
	where := licence.LicenceWhereUniqueDTO{
		ID: f.fakes.licences[0].ID,
	}

	// When try to delete a licence
	// It shoud call the repository with the correct id
	f.mocks.LicenceRepository.On("Delete", where.ID).Return(nil, nil).Times(1)
	_, err := f.service.DeleteLicence(ctx, where)
	assert.NoError(err)
}
