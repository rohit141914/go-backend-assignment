package jobs

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"backend-assignment/internal/store"
)

type JobPayload struct {
	Count  int `json:"count"`
	Visits []struct {
		StoreID   string   `json:"store_id"`
		ImageURLs []string `json:"image_url"`
		VisitTime string   `json:"visit_time"`
	} `json:"visits"`
}

type JobStatus struct {
	Status string `json:"status"`
	Error  []struct {
		StoreID string `json:"store_id"`
		Error   string `json:"error"`
	} `json:"error,omitempty"`
}

type jobManager struct {
	mu       sync.Mutex
	jobStore map[string]*JobStatus
}

func NewJobManager() *jobManager {
	return &jobManager{
		jobStore: make(map[string]*JobStatus),
	}
}

func (p *JobPayload) Validate() error {
	if p.Count != len(p.Visits) {
		return errors.New("count does not match the number of visits")
	}
	for _, visit := range p.Visits {
		if visit.StoreID == "" || len(visit.ImageURLs) == 0 {
			return errors.New("store_id or image_url is missing")
		}
	}
	return nil
}

func (jm *jobManager) AddJob(payload JobPayload) string {
	jobID := fmt.Sprintf("job_%d", rand.Intn(100000))
	jm.mu.Lock()
	jm.jobStore[jobID] = &JobStatus{Status: "ongoing"}
	jm.mu.Unlock()

	go jm.processJob(jobID, payload)
	return jobID
}

func (jm *jobManager) processJob(jobID string, payload JobPayload) {
	time.Sleep(time.Millisecond * 100) // Simulate some delay

	var failedStores []struct {
		StoreID string `json:"store_id"`
		Error   string `json:"error"`
	}

	for _, visit := range payload.Visits {
		if !store.IsValidStoreID(visit.StoreID) {
			failedStores = append(failedStores, struct {
				StoreID string `json:"store_id"`
				Error   string `json:"error"`
			}{
				StoreID: visit.StoreID,
				Error:   "invalid store ID",
			})
			continue
		}

		for _, imageURL := range visit.ImageURLs {
			if err := processImage(imageURL); err != nil {
				failedStores = append(failedStores, struct {
					StoreID string `json:"store_id"`
					Error   string `json:"error"`
				}{
					StoreID: visit.StoreID,
					Error:   fmt.Sprintf("failed to process image: %v", err),
				})
			}
		}
	}

	jm.mu.Lock()
	if len(failedStores) > 0 {
		jm.jobStore[jobID] = &JobStatus{Status: "failed", Error: failedStores}
	} else {
		jm.jobStore[jobID] = &JobStatus{Status: "completed"}
	}
	jm.mu.Unlock()
}

func (jm *jobManager) GetJobStatus(jobID string) (*JobStatus, error) {
	jm.mu.Lock()
	defer jm.mu.Unlock()

	status, exists := jm.jobStore[jobID]
	if !exists {
		return nil, errors.New("job ID does not exist")
	}
	return status, nil
}

func processImage(imageURL string) error {
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(300))) // Simulate GPU processing
	return nil // Replace with actual image download and processing logic
}
