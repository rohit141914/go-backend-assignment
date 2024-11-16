package jobs

import (
	"fmt"
	// "net/http"
	"sync"
	"time"
)

// Job struct represents a processing job
type Job struct {
	ID     string
	Status string
	Errors []StoreError // To track failed store IDs and their errors
	mu     sync.Mutex   // To ensure thread-safe updates
	Stores []string     // List of valid store IDs for the job
}

// StoreError struct represents an error for a specific store_id
type StoreError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

// NewJob creates a new Job instance
func NewJob(id string, stores []string) *Job {
	return &Job{
		ID:     id,
		Status: "ongoing", // Initial status is ongoing
		Stores: stores,
	}
}

// SetJobStatus updates the status of the job
func (j *Job) SetJobStatus(status string) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.Status = status
}

// GetJobStatus returns the current status of the job
func (j *Job) GetJobStatus() string {
	j.mu.Lock()
	defer j.mu.Unlock()
	return j.Status
}

// AddError adds an error related to a specific store_id
func (j *Job) AddError(storeID, errorMsg string) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.Errors = append(j.Errors, StoreError{StoreID: storeID, Error: errorMsg})
}

// ProcessJob simulates processing the job with a 5-second delay
func (j *Job) ProcessJob() {
	// Set the status to "ongoing"
	j.SetJobStatus("ongoing")

	// Simulate a 5-second processing delay
	time.Sleep(5 * time.Second)

	// Simulate checking the store_id and downloading images
	for _, storeID := range j.Stores {
		if !j.isStoreIDValid(storeID) {
			// If store_id is invalid, add an error
			j.AddError(storeID, "Store ID does not exist")
		} else if !j.downloadImageForStore(storeID) {
			// If image download fails, add an error
			j.AddError(storeID, "Image download failed")
		}
	}

	// If there are any errors, mark the job as failed
	if len(j.Errors) > 0 {
		j.SetJobStatus("failed")
	} else {
		// If no errors, mark as completed
		j.SetJobStatus("completed")
	}

	fmt.Printf("Job %s processing completed with status: %s\n", j.ID, j.Status)
}

// Simulate checking if store ID is valid
func (j *Job) isStoreIDValid(storeID string) bool {
	// Assuming store IDs that start with "RP" are valid for this example
	if storeID[:2] == "RP" {
		return true
	}
	return false
}

// Simulate image download for a store (you can replace this with real image downloading logic)
func (j *Job) downloadImageForStore(storeID string) bool {
	// In this simulation, we assume image download fails for any store ID that ends with "3"
	if storeID[len(storeID)-1] == '3' {
		return false // Simulate failure for store IDs ending in '3'
	}
	return true
}
