package licence

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * Types
 */

// Feature
type Feature string

const (
	POST     Feature = "post"
	CATEGORY Feature = "category"
	LICENCE  Feature = "licence"
)

var Features []Feature = []Feature{POST, CATEGORY, LICENCE}

// Permission
type Permission string

const (
	READ   Permission = "read"
	CREATE Permission = "create"
	UPDATE Permission = "update"
	DELETE Permission = "delete"
)

/**
 * Model
 */
type Licence struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Grants    []Grant            `json:"permissions" bson:"permissions"`
	UsedAt    *time.Time         `json:"used_at,omitempty" bson:"used_at,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type Grant struct {
	Feature     Feature      `json:"feature" bson:"feature,omitempty" validate:"required"`
	Version     string       `json:"version" bson:"version,omitempty" validate:"required"`
	Permissions []Permission `json:"permission" bson:"permission,omitempty" validate:"required"`
}

/**
 * Dto
 */
// Read
type LicenceWhereUniqueDTO struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
type LicenceWhereDTO struct {
	ID     *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Grants []*Grant            `json:"grants" bson:"grants,omitempty" validate:"omitempty,dive"`
	OR     []bson.M            `json:"$or,omitempty" bson:"$or,omitempty"`
}

// Create
type LicenceCreateDTO struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Grants    []Grant             `json:"grants" bson:"grants" validate:"required,dive"`
	UsedAt    *time.Time          `json:"used_at,omitempty" bson:"used_at,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Update
type LicenceUpdateDTO struct {
	Grants    []*Grant   `json:"grants" bson:"grants" validate:"omitempty,dive"`
	UsedAt    *time.Time `json:"used_at,omitempty" bson:"used_at,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Enums
type LicenceOrderByENUM string

/**
 * Presets
 */

// Visitor Default Licence
func GetVisitorLicence() *Licence {
	id := primitive.NewObjectID()
	defaultTime := time.Now()

	var grants []Grant
	for _, f := range Features {
		grant := Grant{
			Feature:     f,
			Version:     "1",
			Permissions: []Permission{READ},
		}

		// Blacklisted features
		if f == LICENCE {
			grant.Permissions = []Permission{}
		}

		grants = append(grants, grant)
	}

	return &Licence{
		Grants:    grants,
		ID:        id,
		UsedAt:    nil,
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}
}

func GetSuperAdminLicence() *Licence {
	id := primitive.NewObjectID()
	defaultTime := time.Now()

	var grants []Grant
	for _, f := range Features {
		grant := Grant{
			Feature:     f,
			Version:     "1",
			Permissions: []Permission{READ, CREATE, UPDATE, DELETE},
		}
		grants = append(grants, grant)
	}

	return &Licence{
		Grants:    grants,
		ID:        id,
		UsedAt:    nil,
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}
}
