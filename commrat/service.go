package commrat

import (
	"context"
)

type CommRatService interface {
	CreateComment(ctx context.Context, text string, userToken string, chargerID string) (string, error)
	GetComment(ctx context.Context, id string) (Comment, error)
	UpdateComment(ctx context.Context, id string, text string) (string, error)
	GetComments(ctx context.Context) ([]Comment, error)
	GetCommentsFilter(ctx context.Context, chargerID string, userID string) ([]Comment, error)
	DeleteComment(ctx context.Context, id string) (string, error)
	//
	CreateRating(ctx context.Context, rating float64, userToken string, chargerID string) (string, error)
	GetRating(ctx context.Context, id string) (Rating, error)
	UpdateRating(ctx context.Context, id string, rating float64) (string, error)
	GetRatings(ctx context.Context) ([]Rating, error)
	GetRatingsFilter(ctx context.Context, chargerID string, userID string) ([]Rating, error)
	DeleteRating(ctx context.Context, id string) (string, error)
}
