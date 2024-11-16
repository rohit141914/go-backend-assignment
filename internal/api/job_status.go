package api

import (
	"sync"
)

// Declare a global variable to store job statuses and a mutex to sync access
var (
	jobStatuses = make(map[string]string)
	mu          sync.Mutex
)
