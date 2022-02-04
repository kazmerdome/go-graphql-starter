package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User
type User struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string              `json:"email" bson:"email,omitempty"`
	Username  string              `json:"username" bson:"username,omitempty" validate:"omitempty,min=2"`
	Details   []*Detail           `json:"details" bson:"details" validate:"omitempty,dive"`
	Licence   *primitive.ObjectID `json:"licence" bson:"licence"`
	LastLogin *time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastSeen  *time.Time          `json:"last_seen,omitempty" bson:"last_seen,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Detail struct {
	Name  DetailNameType `json:"name" bson:"name,omitempty" validate:"omitempty,min=2"`
	Value string         `json:"value" bson:"value,omitempty" validate:"omitempty,min=2"`
}

type DetailNameType string

const (
	Firstname DetailNameType = "firstname"
	Lastname  DetailNameType = "lastname"
	Phone     DetailNameType = "phone"
)

// Enums
type UserOrderByENUM string

// Read
type UserWhereUniqueDTO struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
type UserWhereDTO struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    *string             `json:"email" bson:"email,omitempty"`
	Username *string             `json:"username" bson:"username,omitempty"`
	Licence  *primitive.ObjectID `json:"licence" bson:"licence,omitempty"`
	Details  *[]*Detail          `json:"details" bson:"details,omitempty" validate:"omitempty,dive"`

	OR []bson.M `json:"$or,omitempty" bson:"$or,omitempty"`
}

// // UserCreateDTO
// type UserCreateDTO struct {
// 	identity.IdentityCreateDTO `bson:",inline"`
// 	Firstname                  *string `json:"firstname" bson:"firstname,omitempty" validate:"omitempty,min=2"`
// 	Lastname                   *string `json:"lastname" bson:"lastname,omitempty" validate:"omitempty,min=2"`
// 	Phone                      *string `json:"phone" bson:"phone,omitempty" validate:"omitempty,min=2"`
// }

// // UserUpdateDTO
// type UserUpdateDTO struct {
// 	identity.IdentityUpdateDTO `bson:",inline"`
// 	Firstname                  *string `json:"firstname" bson:"firstname,omitempty" validate:"omitempty,min=2"`
// 	Lastname                   *string `json:"lastname" bson:"lastname,omitempty" validate:"omitempty,min=2"`
// 	Phone                      *string `json:"phone" bson:"phone,omitempty" validate:"omitempty,min=2"`
// }
