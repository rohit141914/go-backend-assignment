package main

import (
	"backend-assignment/internal/api"
	"backend-assignment/internal/store"
	"log"
	"net/http"
	"os"
	"encoding/csv"
	// "fmt"
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
	// Load store master data from the CSV file
	stores, err := loadStoreData("StoreMasterAssignment.csv")
	if err != nil {
		log.Fatal("Error loading store data:", err)
	}

	// Register API routes and pass the stores data
	api.RegisterRoutes(stores)

	// Start the server
	log.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
