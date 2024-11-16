package api

import (
	"backend-assignment/internal/jobs"
	"backend-assignment/internal/store"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var (
	jobMap = make(map[string]*jobs.Job)
	mu     sync.Mutex // Protect access to jobMap
)

// SubmitJobHandler handles job submission
func SubmitJobHandler(stores []store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Count  int `json:"count"`
			Visits []struct {
				StoreID   string   `json:"store_id"`
				ImageURL  []string `json:"image_url"`
				VisitTime string   `json:"visit_time"`
			} `json:"visits"`
		}

		// Decode the request body
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Error: Failed to parse request body. Please ensure it is a valid JSON.", http.StatusBadRequest)
			return
		}

		// Validate the request
		if requestBody.Count != len(requestBody.Visits) {
			errorMessage := map[string]string{
				"error": "The 'count' field does not match the number of 'visits'.",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMessage)
			return
		}

		// Generate unique job ID
		mu.Lock()
		jobID := "job_" + time.Now().Format("20060102150405")
		mu.Unlock()

		// Create a new job
		storeIDs := []string{}
		for _, visit := range requestBody.Visits {
			storeIDs = append(storeIDs, visit.StoreID)
		}
		job := jobs.NewJob(jobID, storeIDs)
		mu.Lock()
		jobMap[jobID] = job
		mu.Unlock()

		// Process the job asynchronously
		go func() {
			job.ProcessJob()
		}()

		// Respond with the job ID
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"job_id": jobID,
		})
	}
}

// GetJobInfoHandler handles fetching job status
func GetJobInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobID := r.URL.Query().Get("jobid")
		if jobID == "" {
			http.Error(w, "JobID is required", http.StatusBadRequest)
			return
		}

		mu.Lock()
		job, exists := jobMap[jobID]
		mu.Unlock()

		if !exists {
			http.Error(w, "Job not found", http.StatusBadRequest)
			return
		}

		// Respond with the current job status and any errors
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"job_id": jobID,
			"status": job.GetJobStatus(),
		}

		// If there are errors, add them to the response
		if len(job.Errors) > 0 {
			response["error"] = job.Errors
		}

		json.NewEncoder(w).Encode(response)
	}
}
