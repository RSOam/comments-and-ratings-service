package commrat

import (
	"context"
	"net/http"

	ht "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/comments").Handler(ht.NewServer(
		endpoints.CreateComment,
		decodeCreateCommentRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/comments/").Handler(ht.NewServer(
		endpoints.GetCommentsFilter,
		decodeGetCommentsFilterRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/comments/{id}").Handler(ht.NewServer(
		endpoints.GetComment,
		decodeGetCommentRequest,
		encodeResponse,
	))
	r.Methods("PUT").Path("/comments/{id}").Handler(ht.NewServer(
		endpoints.UpdateComment,
		decodeUpdateCommentRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/comments").Handler(ht.NewServer(
		endpoints.GetComments,
		decodeGetCommentsRequest,
		encodeResponse,
	))
	r.Methods("DELETE").Path("/comments/{id}").Handler(ht.NewServer(
		endpoints.DeleteComment,
		decodeDeleteCommentRequest,
		encodeResponse,
	))
	//RATINGS
	r.Methods("POST").Path("/ratings").Handler(ht.NewServer(
		endpoints.CreateRating,
		decodeCreateRatingRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/ratings/{id}").Handler(ht.NewServer(
		endpoints.GetRating,
		decodeGetRatingRequest,
		encodeResponse,
	))
	r.Methods("PUT").Path("/ratings/{id}").Handler(ht.NewServer(
		endpoints.UpdateRating,
		decodeUpdateRatingRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/ratings/").Handler(ht.NewServer(
		endpoints.GetRatingsFilter,
		decodeGetRatingsFilterRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/ratings").Handler(ht.NewServer(
		endpoints.GetRatings,
		decodeGetRatingsRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/ratings/{id}").Handler(ht.NewServer(
		endpoints.DeleteRating,
		decodeDeleteRatingRequest,
		encodeResponse,
	))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
