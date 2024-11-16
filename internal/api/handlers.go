package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"backend-assignment/internal/store"
	"time"
)

// Handle job submission, validate store IDs and return job ID
func SubmitJobHandler(w http.ResponseWriter, r *http.Request, stores []store.Store) {
	var job map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	visits, ok := job["visits"].([]interface{})
	if !ok {
		http.Error(w, "Invalid visits format", http.StatusBadRequest)
		return
	}

	// Validate each store ID
	for _, v := range visits {
		visit, ok := v.(map[string]interface{})
		if !ok {
			http.Error(w, "Invalid visit data format", http.StatusBadRequest)
			return
		}
		storeID, ok := visit["store_id"].(string)
		if !ok || !store.IsValidStoreID(storeID, stores) {
			http.Error(w, fmt.Sprintf("Invalid store ID: %s", storeID), http.StatusBadRequest)
			return
		}
	}

	// Simulate job creation and set the initial status to "ongoing"
	jobID := "job_78688" // Generate a unique job ID
	mu.Lock() // Locking to ensure safe concurrent access to the map
	jobStatuses[jobID] = "ongoing"
	mu.Unlock()

	// Simulate processing the job and updating the status
	go processJob(jobID)

	// Respond with the job ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"job_id":"%s"}`, jobID)
}

// Handle checking the job status
func GetJobInfoHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, "Job ID is required", http.StatusBadRequest)
		return
	}

	// Check if the job exists in the status map
	mu.Lock()
	status, exists := jobStatuses[jobID]
	mu.Unlock()

	if !exists {
		http.Error(w, "Job ID not found", http.StatusBadRequest)
		return
	}

	// Return the job status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"%s","job_id":"%s"}`, status, jobID)
}

// Simulate processing the job
func processJob(jobID string) {
	// Simulating the job processing with a random delay
	time.Sleep(5 * time.Second) // Simulate processing time

	// Simulating job completion
	mu.Lock()
	jobStatuses[jobID] = "completed"
	mu.Unlock()
}
