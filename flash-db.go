// github.com/Peter-Bird/db
// pkg/db/repository.go

package FlashDB

import (
	"errors"
	"fmt"
	"sync"
)

type FlashDB struct {
	jsonData map[string]interface{}
	mu       sync.RWMutex
}

// Ensure FlashDB implements Repository interface
var _ Repository = (*FlashDB)(nil)

// NewFlashDB initializes the database
func NewFlashDB() *FlashDB {
	return &FlashDB{
		jsonData: make(map[string]interface{}),
	}
}

// List retrieves all records
func (db *FlashDB) List() ([]map[string]interface{}, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	result := make([]map[string]interface{}, 0, len(db.jsonData))
	for _, value := range db.jsonData {
		if record, ok := value.(map[string]interface{}); ok {
			result = append(result, record)
		}
	}
	return result, nil
}

// Get retrieves a record by ID
func (db *FlashDB) Get(id string) (map[string]interface{}, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if data, ok := db.jsonData[id]; ok {
		if result, ok := data.(map[string]interface{}); ok {
			return result, nil
		}
		return nil, errors.New("record exists but is in an unexpected format")
	}
	return nil, fmt.Errorf("record with ID %s not found", id)
}

// Save adds or updates a record
func (db *FlashDB) Save(id string, data map[string]interface{}) error {
	if data == nil {
		return errors.New("data cannot be nil")
	}
	db.mu.Lock()
	defer db.mu.Unlock()

	db.jsonData[id] = data
	return nil
}

// Delete removes a record by ID
func (db *FlashDB) Delete(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.jsonData[id]; ok {
		delete(db.jsonData, id)
		return nil
	}
	return fmt.Errorf("record with ID %s not found", id)
}

// Truncate removes all records from the database
func (db *FlashDB) Truncate() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.jsonData = make(map[string]interface{})

	return nil
}
