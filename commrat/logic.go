package commrat

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	consulapi "github.com/hashicorp/consul/api"
)

type service struct {
	db     CommRatDB
	logger log.Logger
	consul consulapi.Client
}

func NewService(db CommRatDB, logger log.Logger, consul consulapi.Client) CommRatService {
	return &service{
		db:     db,
		logger: logger,
		consul: consul,
	}
}

func (s service) CreateComment(ctx context.Context, text string, userID string, chargerID string) (string, error) {
	logger := log.With(s.logger, "method: ", "CreateComment")

	if err := s.db.CreateComment(ctx, text, userID, chargerID); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create Comment", nil)
	return "Ok", nil
}
func (s service) GetComment(ctx context.Context, id string) (Comment, error) {
	logger := log.With(s.logger, "method", "GetComment")
	comment, err := s.db.GetComment(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return comment, err
	}
	logger.Log("Get Comment", id)
	return comment, nil
}
func (s service) GetComments(ctx context.Context) ([]Comment, error) {
	logger := log.With(s.logger, "method", "GetComments")
	comments, err := s.db.GetComments(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return comments, err
	}
	logger.Log("Get Comments")
	return comments, nil
}
func (s service) GetCommentsFilter(ctx context.Context, chargerID string, userID string) ([]Comment, error) {
	logger := log.With(s.logger, "method", "GetCommentsFilter")
	comments, err := s.db.GetCommentsFilter(ctx, chargerID, userID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return comments, err
	}
	logger.Log("Get CommentsFilter")
	return comments, nil
}

func (s service) DeleteComment(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteComment")
	err := s.db.DeleteComment(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Delete Comment", id)
	return "Ok", nil
}
func (s service) UpdateComment(ctx context.Context, id string, text string) (string, error) {
	logger := log.With(s.logger, "method: ", "UpdateComment")

	if err := s.db.UpdateComment(ctx, id, text); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("update Comment", id)
	return "Ok", nil
}

//RATING
func (s service) CreateRating(ctx context.Context, rating float64, userID string, chargerID string) (string, error) {
	logger := log.With(s.logger, "method: ", "CreateRating")

	if err := s.db.CreateRating(ctx, rating, userID, chargerID); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create Rating", nil)
	return "Ok", nil
}
func (s service) GetRating(ctx context.Context, id string) (Rating, error) {
	logger := log.With(s.logger, "method", "GetRating")
	rating, err := s.db.GetRating(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return rating, err
	}
	logger.Log("Get Rating", id)
	return rating, nil
}
func (s service) GetRatings(ctx context.Context) ([]Rating, error) {
	logger := log.With(s.logger, "method", "GetRatings")
	chargers, err := s.db.GetRatings(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return chargers, err
	}
	logger.Log("Get Ratings")
	return chargers, nil
}
func (s service) GetRatingsFilter(ctx context.Context, chargerID string, userID string) ([]Rating, error) {
	logger := log.With(s.logger, "method", "GetRatingsFilter")
	ratings, err := s.db.GetRatingsFilter(ctx, chargerID, userID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return ratings, err
	}
	logger.Log("Get RatingsFilter")
	return ratings, nil
}
func (s service) DeleteRating(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteRating")
	err := s.db.DeleteRating(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Delete Rating", id)
	return "Ok", nil
}
func (s service) UpdateRating(ctx context.Context, id string, rating float64) (string, error) {
	logger := log.With(s.logger, "method: ", "UpdateRating")

	if err := s.db.UpdateRating(ctx, id, rating); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("update Rating", id)
	return "Ok", nil
}
