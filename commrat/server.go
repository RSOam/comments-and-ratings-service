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

	r.Methods("GET").Path("/comments/{id}").Handler(ht.NewServer(
		endpoints.GetComment,
		decodeGetCommentRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/comments").Handler(ht.NewServer(
		endpoints.GetComments,
		decodeGetCommentsRequest,
		encodeResponse,
	))
	r.Methods("DELETE").Path("/comments/{id}").Handler(ht.NewServer(
		endpoints.DeleteComment,
		decodeDeleteChargerRequest,
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
