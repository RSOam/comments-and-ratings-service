package commrat

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChargerID primitive.ObjectID `json:"chargerID"`
	UserID    primitive.ObjectID `json:"userID"`
	Text      string             `json:"text"`
	Created   string             `json:"created"`
	Modified  string             `json:"modified"`
}
type Rating struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChargerID primitive.ObjectID `json:"chargerID"`
	UserID    primitive.ObjectID `json:"userID"`
	Rating    float64            `json:"rating"`
	Created   string             `json:"created"`
	Modified  string             `json:"modified"`
}

type CommRatDB interface {
	CreateComment(ctx context.Context, text string, userID string, chargerID string) error
	GetComment(ctx context.Context, id string) (Comment, error)
	GetComments(ctx context.Context) ([]Comment, error)
	GetCommentsFilter(ctx context.Context, chargerID string, userID string) ([]Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, id string, text string) error
	//
	CreateRating(ctx context.Context, rating float64, userID string, chargerID string) error
	GetRating(ctx context.Context, id string) (Rating, error)
	GetRatings(ctx context.Context) ([]Rating, error)
	GetRatingsFilter(ctx context.Context, chargerID string, userID string) ([]Rating, error)
	DeleteRating(ctx context.Context, id string) error
	UpdateRating(ctx context.Context, id string, rating float64) error
}
