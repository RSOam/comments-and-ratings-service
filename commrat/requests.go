package commrat

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateCommentRequest struct {
		ChargerID string `json:"chargerID"`
		UserToken string `json:"userToken"`
		Text      string `json:"text"`
	}
	CreateCommentResponse struct {
		Status string `json:"status"`
	}
	GetCommentRequest struct {
		Id string `json:"id"`
	}
	GetCommentResponse struct {
		Text    string `json:"text"`
		Created string `json:"created"`
	}
	GetCommentsRequest struct {
	}
	GetCommentsFilterRequest struct {
		ChargerID string `json:"chargerID"`
		UserID    string `json:"userID"`
	}
	GetCommentsFilterResponse struct {
		Comments []Comment `json:"comments"`
	}
	GetCommentsResponse struct {
		Comments []Comment `json:"comments"`
	}
	DeleteCommentRequest struct {
		Id string `json:"id"`
	}
	DeleteCommentResponse struct {
		Status string `json:"status"`
	}
	UpdateCommentRequest struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}
	UpdateCommentResponse struct {
		Status string `json:"status"`
	}

	//RATINGS
	CreateRatingRequest struct {
		ChargerID string  `json:"chargerID"`
		UserToken string  `json:"userToken"`
		Rating    float64 `json:"rating"`
	}
	CreateRatingResponse struct {
		Status string `json:"status"`
	}
	GetRatingRequest struct {
		Id string `json:"id"`
	}
	GetRatingResponse struct {
		Rating  float64 `json:"rating"`
		Created string  `json:"created"`
	}
	GetRatingsRequest struct {
	}
	GetRatingsResponse struct {
		Ratings []Rating `json:"ratings"`
	}
	DeleteRatingRequest struct {
		Id string `json:"id"`
	}
	DeleteRatingResponse struct {
		Status string `json:"status"`
	}
	UpdateRatingRequest struct {
		Id     string  `json:"id"`
		Rating float64 `json:"rating"`
	}
	UpdateRatingResponse struct {
		Status string `json:"status"`
	}
	GetRatingsFilterRequest struct {
		ChargerID string `json:"chargerID"`
		UserID    string `json:"userID"`
	}
	GetRatingsFilterResponse struct {
		Ratings []Rating `json:"ratings"`
	}
	//OTHER
	PostChargerUpdateRequest struct {
		AverageRating float64 `json:"averageRating"`
	}
	PostChargerUpdateResponse struct {
		Status string `json:"status"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := CreateCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	req.UserToken = r.Header.Get("Authorization")
	if err != nil {
		return nil, err
	}
	return req, nil
}
func decodeGetCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := GetCommentRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	return req, nil
}
func decodeGetCommentsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := GetCommentRequest{}
	return req, nil
}
func decodeGetCommentsFilterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := GetCommentsFilterRequest{}
	req.ChargerID = r.URL.Query().Get("charger")
	req.UserID = r.URL.Query().Get("user")
	return req, nil
}
func decodeDeleteCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := DeleteCommentRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	return req, nil
}
func decodeUpdateRatingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := UpdateRatingRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//RATINGS
func decodeCreateRatingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := CreateRatingRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	req.UserToken = r.Header.Get("Authorization")
	if err != nil {
		return nil, err
	}
	return req, nil
}
func decodeGetRatingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := GetRatingRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	return req, nil
}
func decodeGetRatingsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := GetRatingRequest{}
	return req, nil
}
func decodeDeleteRatingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := DeleteCommentRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	return req, nil
}
func decodeUpdateCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := UpdateCommentRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func decodeGetRatingsFilterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := GetRatingsFilterRequest{}
	req.ChargerID = r.URL.Query().Get("charger")
	req.UserID = r.URL.Query().Get("user")
	return req, nil
}
