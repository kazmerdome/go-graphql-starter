package category

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category
// CategoryCreateDTO
// CategoryUpdateDTO
type Category struct {
	ID          *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string              `json:"title" bson:"title,omitempty" validate:"required,min=2"`
	Slug        string              `json:"slug" bson:"slug,omitempty" validate:"required,min=2"`
	Description *string             `json:"description" bson:"description,omitempty"`
	Order       *int                `json:"order" bson:"order,omitempty" validate:"number"`
	CreatedAt   time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Enums
type CategoryOrderByENUM string

// Read
type CategoryWhereUniqueDTO struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
type CategoryWhereDTO struct {
	ID    *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title *string             `json:"title" bson:"title,omitempty"`
	Slug  *string             `json:"slug" bson:"slug,omitempty"`
	Order *int                `json:"order" bson:"order,omitempty"`
}
