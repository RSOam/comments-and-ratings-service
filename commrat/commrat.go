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

type CommRatDB interface {
	CreateComment(ctx context.Context, text string, userID string, chargerID string) error
	GetComment(ctx context.Context, id string) (Comment, error)
	GetComments(ctx context.Context) ([]Comment, error)
	DeleteComment(ctx context.Context, id string) error
}
