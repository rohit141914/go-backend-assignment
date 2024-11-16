package jobs

import (
	"fmt"
	"sync"
	"time"
)

// Job struct represents a processing job
type Job struct {
	ID     string
	Status string
	mu     sync.Mutex // To ensure thread-safe updates
}

// NewJob creates a new Job instance
func NewJob(id string) *Job {
	return &Job{
		ID:     id,
		Status: "ongoing", // Initial status is ongoing
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

// ProcessJob simulates processing the job with a 5-second delay
func (j *Job) ProcessJob() {
	// Simulate a 5-second processing delay
	time.Sleep(5 * time.Second)

	// After the delay, update the status to "completed"
	j.SetJobStatus("completed")
	fmt.Printf("Job %s processing completed\n", j.ID)
}
