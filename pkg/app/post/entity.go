package post

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post
// PostCreateDTO
// PostUpdateDTO
type Post struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string              `json:"title" bson:"title,omitempty" validate:"required,min=2"`
	Slug      string              `json:"slug" bson:"slug,omitempty" validate:"required,min=2"`
	Category  *primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty" validate:"required"`
	Content   *string             `json:"content" bson:"content,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Enums
type PostOrderByENUM string

// Read
type PostWhereUniqueDTO struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
type PostWhereDTO struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    *string             `json:"title" bson:"title,omitempty"`
	Slug     *string             `json:"slug" bson:"slug,omitempty"`
	Category *primitive.ObjectID `json:"category" bson:"category,omitempty"`
}
