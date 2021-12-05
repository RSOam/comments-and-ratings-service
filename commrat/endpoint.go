package commrat

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateComment endpoint.Endpoint
	GetComment    endpoint.Endpoint
	GetComments   endpoint.Endpoint
	DeleteComment endpoint.Endpoint
}

func MakeEndpoints(s CommRatService) Endpoints {
	return Endpoints{
		CreateComment: makeCreateCommentEndpoint(s),
		GetComment:    makeGetCommentEndpoint(s),
		GetComments:   makeGetCommentsEndpoint(s),
		DeleteComment: makeDeleteCommentEndpoint(s),
	}
}

func makeCreateCommentEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCommentRequest)
		status, err := s.CreateComment(ctx, req.Text, req.UserID, req.ChargerID)
		return CreateCommentResponse{Status: status}, err
	}
}
func makeGetCommentEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCommentRequest)
		comment, err := s.GetComment(ctx, req.Id)
		return GetCommentResponse{
			Text:    comment.Text,
			Created: comment.Created,
		}, err
	}
}
func makeGetCommentsEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		comments, err := s.GetComments(ctx)
		return GetCommentsResponse{
			Comments: comments,
		}, err
	}
}
func makeDeleteCommentEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCommentRequest)
		status, err := s.DeleteComment(ctx, req.Id)
		return DeleteCommentResponse{
			Status: status,
		}, err
	}
}
