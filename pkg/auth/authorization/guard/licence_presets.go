package guard

import (
	"time"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Visitor Default Licence
func GetVisitorLicence() *licence.Licence {
	id := primitive.NewObjectID()
	defaultTime := time.Now()

	var grants []licence.Grant
	for _, f := range licence.Features {
		grant := licence.Grant{
			Feature:     f,
			Version:     "1",
			Permissions: []licence.Permission{licence.READ},
		}

		// Blacklisted features
		if f == licence.LICENCE {
			grant.Permissions = []licence.Permission{}
		}

		grants = append(grants, grant)
	}

	return &licence.Licence{
		Grants:    grants,
		ID:        id,
		UsedAt:    nil,
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}
}

func GetSuperAdminLicence() *licence.Licence {
	id := primitive.NewObjectID()
	defaultTime := time.Now()

	var grants []licence.Grant
	for _, f := range licence.Features {
		grant := licence.Grant{
			Feature:     f,
			Version:     "1",
			Permissions: []licence.Permission{licence.READ, licence.CREATE, licence.UPDATE, licence.DELETE},
		}
		grants = append(grants, grant)
	}

	return &licence.Licence{
		Grants:    grants,
		ID:        id,
		UsedAt:    nil,
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}
}
