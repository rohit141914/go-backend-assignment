package main

import (
	"backend-assignment/internal/api"
	"backend-assignment/internal/store"
	"log"
	"net/http"
	"os"
	"encoding/csv"
	// "fmt"
	"github.com/gorilla/mux"

)

// Load the store master data from CSV file
func loadStoreData(filename string) ([]store.Store, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip the first line containing headers
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var stores []store.Store
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		stores = append(stores, store.Store{
			AreaCode: record[0],
			StoreName: record[1],
			StoreID: record[2],
		})
	}

	return stores, nil
}

func main() {
	// Load stores from the CSV file
	stores, err := store.LoadStores("StoreMasterAssignment.csv")
	if err != nil {
		log.Fatalf("Error loading store data: %v", err)
	}

	// Create a new router
	router := mux.NewRouter()

	// Register API routes
	api.RegisterRoutes(router, stores)

	// Start the server
	port := "8080"
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}