package licence

import (
	"context"

	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/misc"
)

const (
	ERR_ACCESS_DENIED = "access denied"
)

type LicenceService interface {
	Licence(ctx context.Context, where *LicenceWhereDTO, search *string) (*Licence, error)
	Licences(ctx context.Context, where *LicenceWhereDTO, orderBy *LicenceOrderByENUM, skip *int, limit *int, search *string) ([]*Licence, error)
	LicenceCount(ctx context.Context, where *LicenceWhereDTO, search *string) (*int, error)
	CreateLicence(ctx context.Context, data LicenceCreateDTO) (*Licence, error)
	UpdateLicence(ctx context.Context, where LicenceWhereUniqueDTO, data LicenceUpdateDTO) (*Licence, error)
	DeleteLicence(ctx context.Context, where LicenceWhereUniqueDTO) (*Licence, error)
}

type licenceService struct {
	shared.SharedService
	r LicenceRepository
}

func NewLicenceService(s shared.SharedService, r LicenceRepository) LicenceService {
	return &licenceService{SharedService: s, r: r}
}

// Licence
func (r *licenceService) Licence(ctx context.Context, where *LicenceWhereDTO, search *string) (*Licence, error) {
	if search != nil {
		where.OR = misc.MongoSearchFieldParser(SEARCH_FILEDS, *search)
	}
	return r.r.One(where)
}

// Licences
func (r *licenceService) Licences(ctx context.Context, where *LicenceWhereDTO, orderBy *LicenceOrderByENUM, skip *int, limit *int, search *string) ([]*Licence, error) {
	if search != nil {
		where.OR = misc.MongoSearchFieldParser(SEARCH_FILEDS, *search)
	}
	return r.r.List(where, orderBy, skip, limit, nil)
}

// LicenceCount
func (r *licenceService) LicenceCount(ctx context.Context, where *LicenceWhereDTO, search *string) (*int, error) {
	if search != nil {
		where.OR = misc.MongoSearchFieldParser(SEARCH_FILEDS, *search)
	}
	return r.r.Count(where)
}

// CreateLicence
func (r *licenceService) CreateLicence(ctx context.Context, data LicenceCreateDTO) (*Licence, error) {
	return r.r.Create(&data)
}

// UpdateLicence
func (r *licenceService) UpdateLicence(ctx context.Context, where LicenceWhereUniqueDTO, data LicenceUpdateDTO) (*Licence, error) {
	return r.r.Update(where.ID, &data)
}

// DeleteLicence
func (r *licenceService) DeleteLicence(ctx context.Context, where LicenceWhereUniqueDTO) (*Licence, error) {
	return r.r.Delete(where.ID)
}
