package commrat

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateComment     endpoint.Endpoint
	GetComment        endpoint.Endpoint
	GetComments       endpoint.Endpoint
	GetCommentsFilter endpoint.Endpoint
	DeleteComment     endpoint.Endpoint
	UpdateComment     endpoint.Endpoint
	//
	CreateRating     endpoint.Endpoint
	GetRating        endpoint.Endpoint
	GetRatings       endpoint.Endpoint
	GetRatingsFilter endpoint.Endpoint
	DeleteRating     endpoint.Endpoint
	UpdateRating     endpoint.Endpoint
}

func MakeEndpoints(s CommRatService) Endpoints {
	return Endpoints{
		CreateComment:     makeCreateCommentEndpoint(s),
		GetComment:        makeGetCommentEndpoint(s),
		GetComments:       makeGetCommentsEndpoint(s),
		GetCommentsFilter: makeGetCommentsFilterEndpoint(s),
		DeleteComment:     makeDeleteCommentEndpoint(s),
		UpdateComment:     makeUpdateCommentEndpoint(s),
		//
		CreateRating:     makeCreateRatingEndpoint(s),
		GetRating:        makeGetRatingEndpoint(s),
		GetRatings:       makeGetRatingsEndpoint(s),
		GetRatingsFilter: makeGetRatingsFilterEndpoint(s),
		DeleteRating:     makeDeleteRatingEndpoint(s),
		UpdateRating:     makeUpdateRatingEndpoint(s),
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
func makeGetCommentsFilterEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCommentsFilterRequest)
		comments, err := s.GetCommentsFilter(ctx, req.ChargerID, req.UserID)
		return GetCommentsFilterResponse{
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
func makeUpdateCommentEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCommentRequest)
		status, err := s.UpdateComment(ctx, req.Id, req.Text)
		return CreateCommentResponse{Status: status}, err
	}
}

//RATINGS
func makeCreateRatingEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRatingRequest)
		status, err := s.CreateRating(ctx, req.Rating, req.UserID, req.ChargerID)
		return CreateRatingResponse{Status: status}, err
	}
}
func makeGetRatingEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRatingRequest)
		rating, err := s.GetRating(ctx, req.Id)
		return GetRatingResponse{
			Rating:  rating.Rating,
			Created: rating.Created,
		}, err
	}
}
func makeGetRatingsEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ratings, err := s.GetRatings(ctx)
		return GetRatingsResponse{
			Ratings: ratings,
		}, err
	}
}
func makeDeleteRatingEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRatingRequest)
		status, err := s.DeleteRating(ctx, req.Id)
		return DeleteRatingResponse{
			Status: status,
		}, err
	}
}
func makeGetRatingsFilterEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRatingsFilterRequest)
		ratings, err := s.GetRatingsFilter(ctx, req.ChargerID, req.UserID)
		return GetRatingsFilterResponse{
			Ratings: ratings,
		}, err
	}
}
func makeUpdateRatingEndpoint(s CommRatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRatingRequest)
		status, err := s.UpdateRating(ctx, req.Id, req.Rating)
		return CreateRatingResponse{Status: status}, err
	}
}
