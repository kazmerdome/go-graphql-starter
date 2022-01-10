package identity

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// /**
//  * Model
//  */
// type Identity struct {
// 	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Email        string             `json:"email" bson:"email,omitempty"`
// 	AuthStrategy []AuthStrategy     `json:"auth_strategy" bson:"auth_strategy"`
// 	Policy       []Policy           `json:"policy" bson:"policy"`
// 	Username     string             `json:"username" bson:"username,omitempty" validate:"omitempty,min=2"`
// 	LastLogin    *time.Time         `json:"last_login,omitempty" bson:"last_login,omitempty"`
// 	LastSeen     *time.Time         `json:"last_seen,omitempty" bson:"last_seen,omitempty"`
// 	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
// 	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
// }
// type AuthStrategy struct {
// 	Type   AuthStrategyType `json:"type" bson:"type,omitempty" validate:"required"`
// 	Secret string           `json:"secret" bson:"secret,omitempty" validate:"omitempty,min=6"`
// }
// type Policy struct {
// 	Resource PolicyResource `json:"resource" bson:"resource,omitempty" validate:"required"`
// 	Role     PolicyRole     `json:"role" bson:"role,omitempty" validate:"required"`
// }

// /**
//  * Dto
//  */
// // Read
// type IdentityWhereUniqueDTO struct {
// 	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// }
// type IdentityWhereDTO struct {
// 	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Email    *string             `json:"email" bson:"email,omitempty"`
// 	Policy   []*Policy           `json:"policy" bson:"policy,omitempty" validate:"omitempty,dive"`
// 	Username *string             `json:"username" bson:"username,omitempty"`
// 	OR       []bson.M            `json:"$or,omitempty" bson:"$or,omitempty"`
// }

// // Create
// type IdentityCreateDTO struct {
// 	ID           *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Email        string              `json:"email" bson:"email,omitempty" validate:"required,email"`
// 	AuthStrategy []AuthStrategy      `json:"auth_strategy" bson:"auth_strategy" validate:"required,dive"`
// 	Policy       []Policy            `json:"policy" bson:"policy" validate:"required,dive"`
// 	Username     string              `json:"username" bson:"username,omitempty" validate:"required,min=2"`
// 	LastLogin    *time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
// 	CreatedAt    time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
// 	UpdatedAt    time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
// }

// // Update
// type IdentityUpdateDTO struct {
// 	Email        string          `json:"email" bson:"email,omitempty"`
// 	AuthStrategy []*AuthStrategy `json:"auth_strategy" bson:"auth_strategy" validate:"omitempty,dive"`
// 	Policy       []*Policy       `json:"policy" bson:"policy" validate:"omitempty,dive"`
// 	Username     string          `json:"username" bson:"username,omitempty" validate:"omitempty,min=2"`
// 	LastLogin    *time.Time      `json:"last_login,omitempty" bson:"last_login,omitempty"`
// 	LastSeen     *time.Time      `json:"last_seen,omitempty" bson:"last_seen,omitempty"`
// 	CreatedAt    time.Time       `json:"created_at,omitempty" bson:"created_at,omitempty"`
// 	UpdatedAt    time.Time       `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
// }

// /**
//  * Types
//  */

// // IdentityAuthStrategyType
// type AuthStrategyType string

// const (
// 	GOOGLE_OAUTH AuthStrategyType = "GOOGLE_OAUTH"
// 	BASIC        AuthStrategyType = "BASIC"
// )

// // IdentityPolicyResource
// type PolicyResource string

// var (
// 	BASE_SERVER PolicyResource = "BASE_SERVER"
// )

// // AppPolicyRole
// type PolicyRole string

// const (
// 	VISITOR PolicyRole = "VISITOR"
// 	USER    PolicyRole = "USER"
// 	EDITOR  PolicyRole = "EDITOR"
// 	ADMIN   PolicyRole = "ADMIN"
// )
