package commrat

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type service struct {
	db     CommRatDB
	logger log.Logger
}

func NewService(db CommRatDB, logger log.Logger) CommRatService {
	return &service{
		db:     db,
		logger: logger,
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
	chargers, err := s.db.GetComments(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return chargers, err
	}
	logger.Log("Get Comments")
	return chargers, nil
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
