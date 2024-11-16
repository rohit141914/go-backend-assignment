package jobs

// Define additional logic for handling job status if needed here.
// For example, we might want to store errors or other information related to job processing.

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

// CreateJobError creates a new job error entry
func CreateJobError(storeID, err string) *JobError {
	return &JobError{
		StoreID: storeID,
		Error:   err,
	}
}
