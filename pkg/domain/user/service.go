package user

import "github.com/kazmerdome/go-graphql-starter/pkg/module/provider/service"

type UserService interface {
	// Licence(ctx context.Context, where *LicenceWhereDTO, search *string) (*Licence, error)
	// Licences(ctx context.Context, where *LicenceWhereDTO, orderBy *LicenceOrderByENUM, skip *int, limit *int, search *string) ([]*Licence, error)
	// LicenceCount(ctx context.Context, where *LicenceWhereDTO, search *string) (*int, error)
	// CreateLicence(ctx context.Context, data LicenceCreateDTO) (*Licence, error)
	// UpdateLicence(ctx context.Context, where LicenceWhereUniqueDTO, data LicenceUpdateDTO) (*Licence, error)
	// DeleteLicence(ctx context.Context, where LicenceWhereUniqueDTO) (*Licence, error)
}

type userService struct {
	service.ServiceConfig
	userRepository UserRepository
}

func newUserService(c service.ServiceConfig, userRepository UserRepository) UserService {
	return &userService{ServiceConfig: c, userRepository: userRepository}
}
