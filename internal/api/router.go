package api

import (
	"backend-assignment/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, stores []store.Store) {
	router.HandleFunc("/api/submit", SubmitJobHandler(stores)).Methods(http.MethodPost)
	router.HandleFunc("/api/status", GetJobInfoHandler()).Methods(http.MethodGet)
}
