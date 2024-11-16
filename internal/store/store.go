package store

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Store represents a store's data
type Store struct {
	AreaCode  string
	StoreName string
	StoreID   string
}

// LoadStores reads the CSV file and returns a list of stores
func LoadStores(filename string) ([]Store, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %v", err)
	}

	var stores []Store
	// Skip the header row (first row)
	for i, record := range records {
		if i == 0 {
			continue
		}

		// Parse each record
		store := Store{
			AreaCode:  record[0],
			StoreName: record[1],
			StoreID:   record[2],
		}
		stores = append(stores, store)
	}

	return stores, nil
}
