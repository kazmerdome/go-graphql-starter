package licence

import (
	"errors"
	"strings"

	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/guard"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	LICENCE_ID_TOKEN_KEY = "lid"

	ERR_INVALID_TOKEN = "invalid header token"
	ERR_UNAUTHORIZED  = "unauthorized"
)

type LicenceGuard interface {
	// Validations
	isWhitelistedPermission(licencePermissions []Permission, requiredPermissions []Permission) bool
	getFeaturePermissionsFromLicence(feature Feature, licence Licence) []Permission
	// Token methods
	getLicenceIDFromBearerToken(bearerToken string) (primitive.ObjectID, error)
	// db methods
	getLicenceFromRepository(oid primitive.ObjectID) (*Licence, error)
	// Guards
	GetPermissionsGuard(bearerToken string, feature Feature) ([]Permission, error)
	GetIdGuard(bearerToken string) (primitive.ObjectID, error)
	AuthGuard(feature Feature, requiredPermissions []Permission, bearerToken string) error
}

type licenceGuard struct {
	guard.GuardConfig
	licenceRepository LicenceRepository
	defaultLicence    *Licence
}

func newLicenceGuard(c guard.GuardConfig, r LicenceRepository) LicenceGuard {
	return &licenceGuard{GuardConfig: c, licenceRepository: r, defaultLicence: GetVisitorLicence()}
}

/*
 * Validations
 */
// At least one!
func (r *licenceGuard) isWhitelistedPermission(licencePermissions []Permission, requiredPermissions []Permission) bool {
	isIncluded := false
	for _, requiredPermission := range requiredPermissions {
		for _, licencePermission := range licencePermissions {
			if licencePermission == requiredPermission {
				isIncluded = true
			}
		}
	}
	return isIncluded
}

func (r *licenceGuard) getFeaturePermissionsFromLicence(feature Feature, l Licence) []Permission {
	var permissions []Permission
	for _, grant := range l.Grants {
		if grant.Feature == feature {
			permissions = grant.Permissions
		}
	}
	return permissions
}

/*
 * Token methods
 */
func (r *licenceGuard) getLicenceIDFromBearerToken(bearer string) (primitive.ObjectID, error) {
	rawTokenParts := strings.Split(bearer, "Bearer ")
	if len(rawTokenParts) < 2 {
		return primitive.ObjectID{}, errors.New(ERR_INVALID_TOKEN)
	}

	// verify jwt token string
	claimData, err := token.VerifyJWTToken(rawTokenParts[1], r.GetConfig().Get(config.ENV_JWT_SESSION_SECRET))
	if err != nil || claimData == nil || claimData[LICENCE_ID_TOKEN_KEY] == "" {
		return primitive.ObjectID{}, err
	}
	lidHex := claimData[LICENCE_ID_TOKEN_KEY]

	// create primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(lidHex)
	if err != nil {
		return primitive.ObjectID{}, errors.New(ERR_UNAUTHORIZED)
	}
	return oid, nil
}

/*
 * Repository methods
 */
func (r *licenceGuard) getLicenceFromRepository(oid primitive.ObjectID) (*Licence, error) {
	filter := LicenceWhereDTO{ID: &oid}
	i, err := r.licenceRepository.One(&filter)
	if err != nil || i == nil || i.ID.Hex() == "" {
		return nil, errors.New(ERR_ACCESS_DENIED)
	}
	return i, err
}

/*
 * Guards
 */
func (r *licenceGuard) GetPermissionsGuard(bearerToken string, feature Feature) ([]Permission, error) {
	// get licence
	var l = r.defaultLicence
	if bearerToken == "" {
		return []Permission{}, errors.New(ERR_ACCESS_DENIED)
	}
	oid, err := r.getLicenceIDFromBearerToken(bearerToken)
	if err != nil {
		return []Permission{}, err
	}
	foundLicence, err := r.getLicenceFromRepository(oid)
	if err != nil {
		return []Permission{}, err
	}
	if foundLicence != nil {
		l = foundLicence
	}

	return r.getFeaturePermissionsFromLicence(feature, *l), nil
}

func (r *licenceGuard) GetIdGuard(bearerToken string) (primitive.ObjectID, error) {
	if bearerToken == "" {
		return primitive.ObjectID{}, errors.New(ERR_ACCESS_DENIED)
	}
	oid, err := r.getLicenceIDFromBearerToken(bearerToken)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return oid, nil
}

func (r *licenceGuard) AuthGuard(feature Feature, requiredPermissions []Permission, bearerToken string) error {
	// get licence
	var l = r.defaultLicence
	if bearerToken != "" {
		oid, err := r.getLicenceIDFromBearerToken(bearerToken)
		if err != nil {
			return err
		}
		foundLicence, err := r.getLicenceFromRepository(oid)
		if err != nil {
			return err
		}
		if foundLicence != nil {
			l = foundLicence
		}
	}
	// get feature permissions
	licencePermissions := r.getFeaturePermissionsFromLicence(feature, *l)

	// check that the licence is granted to the required permissions
	if r.isWhitelistedPermission(licencePermissions, requiredPermissions) {
		return nil
	}
	return errors.New(ERR_UNAUTHORIZED)
}
