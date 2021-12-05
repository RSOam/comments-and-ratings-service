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
		UserID    string `json:"userID"`
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
	GetCommentsResponse struct {
		Comments []Comment `json:"comments"`
	}
	DeleteCommentRequest struct {
		Id string `json:"id"`
	}
	DeleteCommentResponse struct {
		Status string `json:"status"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateCommentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := CreateCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
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
func decodeDeleteChargerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := DeleteCommentRequest{}
	vals := mux.Vars(r)
	req.Id = vals["id"]
	return req, nil
}
