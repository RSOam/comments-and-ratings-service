package commrat

import (
	"context"
)

type CommRatService interface {
	CreateComment(ctx context.Context, name string, userID string, chargerID string) (string, error)
	GetComment(ctx context.Context, id string) (Comment, error)
	GetComments(ctx context.Context) ([]Comment, error)
	DeleteComment(ctx context.Context, id string) (string, error)
}
