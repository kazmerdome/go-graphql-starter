package identity

// import (
// 	"errors"
// 	"strings"

// 	"github.com/kazmerdome/go-graphql-starter/pkg/config"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/util/token"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// const (
// 	ERR_INVALID_TOKEN = "invalid header token"
// 	ERR_UNAUTHORIZED  = "unauthorized"
// 	ERR_ACCESS_DENIED = "access denied"
// )

// type IdentityService interface {
// 	// Role Validations
// 	isWhitelistedRole(existingRole PolicyRole, requiredRoles []PolicyRole) bool
// 	// Token methods
// 	getObjectIDFromBearerToken(bearerToken string) (primitive.ObjectID, error)
// 	// db methods
// 	getIdentityDataFromRepository(oid primitive.ObjectID) (*Identity, error)
// 	getCurrentRoleFromObjectID(oid primitive.ObjectID) PolicyRole

// 	// Guards
// 	GetRoleGuard(bearerToken string) PolicyRole
// 	GetIdGuard(bearerToken string) (primitive.ObjectID, error)
// 	AuthGuard(requiredRoles []PolicyRole, bearerToken string) error
// }

// // AuthUser ...
// type identityService struct {
// 	s  shared.SharedService
// 	r  IdentityRepository
// 	pr PolicyResource
// }

// func NewIdentityService(s shared.SharedService, db repository.MongoDatabase) IdentityService {
// 	r := NewIdentityRepository(s, db.Collection("User")) // yes... it is the user collection... the user entity is just an extended identity!
// 	return &identityService{s: s, r: r, pr: BASE_SERVER}
// }

// func (r *identityService) isWhitelistedRole(existingRole PolicyRole, requiredRoles []PolicyRole) bool {
// 	isIncluded := false
// 	for _, reqiredRole := range requiredRoles {
// 		if reqiredRole == existingRole {
// 			isIncluded = true
// 		}
// 	}
// 	return isIncluded
// }

// func (r *identityService) getObjectIDFromBearerToken(bearer string) (primitive.ObjectID, error) {
// 	rawTokenParts := strings.Split(bearer, "Bearer ")
// 	if len(rawTokenParts) < 2 {
// 		return primitive.ObjectID{}, errors.New(ERR_INVALID_TOKEN)
// 	}
// 	// verify jwt token string
// 	userIDHex, err := token.VerifyJWTToken(rawTokenParts[1], r.s.Config.Get(config.ENV_JWT_SESSION_SECRET))
// 	if err != nil {
// 		return primitive.ObjectID{}, err
// 	}
// 	// create primitive.ObjectID
// 	oid, err := primitive.ObjectIDFromHex(userIDHex)
// 	if err != nil {
// 		return primitive.ObjectID{}, errors.New(ERR_UNAUTHORIZED)
// 	}
// 	return oid, nil
// }

// func (r *identityService) getIdentityDataFromRepository(oid primitive.ObjectID) (*Identity, error) {
// 	filter := IdentityWhereUniqueDTO{ID: oid}
// 	i, err := r.r.OneById(&filter)
// 	if err != nil || i == nil || i.Email == "" {
// 		return nil, errors.New(ERR_ACCESS_DENIED)
// 	}
// 	return i, err
// }

// // The default role is VISITOR !!!
// func (r *identityService) getCurrentRoleFromObjectID(oid primitive.ObjectID) PolicyRole {
// 	currentRole := VISITOR
// 	// fetch identity data from db
// 	data, err := r.getIdentityDataFromRepository(oid)
// 	if err != nil {
// 		return currentRole
// 	}
// 	// check []AppPolicy is exist, if not it means user is also not exist || invalid stored user
// 	if data.Policy == nil && len(data.Policy) < 1 {
// 		return currentRole
// 	}
// 	// get current role from user []AppPolicy
// 	for _, policy := range data.Policy {
// 		if policy.Resource == r.pr {
// 			currentRole = policy.Role
// 		}
// 	}

// 	return currentRole
// }

// func (r *identityService) GetRoleGuard(bearerToken string) PolicyRole {
// 	if bearerToken == "" {
// 		return VISITOR
// 	}
// 	oid, err := r.getObjectIDFromBearerToken(bearerToken)
// 	if err != nil {
// 		return VISITOR
// 	}
// 	return r.getCurrentRoleFromObjectID(oid)
// }

// func (r *identityService) GetIdGuard(bearerToken string) (primitive.ObjectID, error) {
// 	if bearerToken == "" {
// 		return primitive.ObjectID{}, errors.New(ERR_ACCESS_DENIED)
// 	}
// 	oid, err := r.getObjectIDFromBearerToken(bearerToken)
// 	if err != nil {
// 		return primitive.ObjectID{}, err
// 	}
// 	return oid, nil
// }

// func (r *identityService) AuthGuard(requiredRoles []PolicyRole, bearerToken string) error {
// 	if bearerToken == "" && r.isWhitelistedRole(VISITOR, requiredRoles) {
// 		return nil
// 	}
// 	oid, err := r.getObjectIDFromBearerToken(bearerToken)
// 	if err != nil {
// 		return err
// 	}
// 	currentRole := r.getCurrentRoleFromObjectID(oid)
// 	if r.isWhitelistedRole(currentRole, requiredRoles) {
// 		return nil
// 	}
// 	return errors.New(ERR_UNAUTHORIZED)
// }
