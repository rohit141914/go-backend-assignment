package store

import (
	"encoding/csv"
	"os"
)

// Store represents the store data
type Store struct {
	AreaCode  string
	StoreName string
	StoreID   string
}

// LoadStoreData loads and parses the store CSV file
func LoadStoreData(filename string) ([]Store, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip the header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var stores []Store
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		stores = append(stores, Store{
			AreaCode:  record[0],
			StoreName: record[1],
			StoreID:   record[2],
		})
	}

	return stores, nil
}

// IsValidStoreID checks if the store ID exists in the store list
func IsValidStoreID(storeID string, stores []Store) bool {
	for _, store := range stores {
		if store.StoreID == storeID {
			return true
		}
	}
	return false
}
