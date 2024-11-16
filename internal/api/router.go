package api

import (
	"net/http"
	"backend-assignment/internal/store"
)

// Register API routes and handlers
func RegisterRoutes(stores []store.Store) {
	http.HandleFunc("/api/submit/", func(w http.ResponseWriter, r *http.Request) {
		SubmitJobHandler(w, r, stores)
	})

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		GetJobInfoHandler(w, r)
	})
}
