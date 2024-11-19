package FlashDB_test

import (
	"encoding/json"
	"testing"

	db "github.com/Peter-Bird/Flash-DB"
	"github.com/stretchr/testify/assert"
)

func TestFlashDB(t *testing.T) {
	db := db.NewFlashDB()

	// Test Save
	t.Run("Save new record", func(t *testing.T) {
		data := map[string]interface{}{"key": "value"}
		err := db.Save("1", data)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	// Test Get
	t.Run("Get existing record", func(t *testing.T) {
		data, err := db.Get("1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if data["key"] != "value" {
			t.Fatalf("expected key to be 'value', got %v", data["key"])
		}
	})

	t.Run("Get non-existing record", func(t *testing.T) {
		_, err := db.Get("2")
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	// Test List
	t.Run("List all records", func(t *testing.T) {
		data := map[string]interface{}{"another_key": "another_value"}
		_ = db.Save("2", data)

		list, err := db.List()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(list) != 2 {
			t.Fatalf("expected 2 records, got %d", len(list))
		}
	})

	// Test Delete
	t.Run("Delete existing record", func(t *testing.T) {
		err := db.Delete("1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = db.Get("1")
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("Delete non-existing record", func(t *testing.T) {
		err := db.Delete("3")
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestFlashDBWithJSON(t *testing.T) {
	// Initialize the database
	flashDB := db.NewFlashDB()

	// JSON test data
	jsonData1 := `{"key1": "value1", "key2": 42}`
	jsonData2 := `{"key3": "value3", "key4": 84}`
	id1 := "jsonRecord1"
	id2 := "jsonRecord2"
	nonExistentID := "nonexistentJSON"

	// Helper function to unmarshal JSON
	unmarshalJSON := func(data string) map[string]interface{} {
		var result map[string]interface{}
		err := json.Unmarshal([]byte(data), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		return result
	}

	// Unmarshal JSON test data
	testData1 := unmarshalJSON(jsonData1)
	testData2 := unmarshalJSON(jsonData2)

	// Test Save method with JSON data
	t.Run("Save JSON Data", func(t *testing.T) {
		err := flashDB.Save(id1, testData1)
		assert.NoError(t, err, "expected no error when saving valid JSON data")

		err = flashDB.Save(id2, testData2)
		assert.NoError(t, err, "expected no error when saving valid JSON data")
	})

	// Test Get method with JSON data
	t.Run("Get JSON Data", func(t *testing.T) {
		// Retrieve existing JSON record
		data, err := flashDB.Get(id1)
		assert.NoError(t, err, "expected no error when retrieving an existing JSON record")
		assert.Equal(t, testData1, data, "expected retrieved JSON data to match saved data")

		// Attempt to retrieve non-existent JSON record
		data, err = flashDB.Get(nonExistentID)
		assert.Error(t, err, "expected an error when retrieving a non-existent JSON record")
		assert.Nil(t, data, "expected retrieved data to be nil for non-existent JSON record")
	})

	// Test List method with JSON data
	t.Run("List JSON Data", func(t *testing.T) {
		// Retrieve all JSON records
		dataList, err := flashDB.List()
		assert.NoError(t, err, "expected no error when listing all JSON records")
		assert.Len(t, dataList, 2, "expected two JSON records in the database")
		assert.Contains(t, dataList, testData1, "expected listed JSON data to include testData1")
		assert.Contains(t, dataList, testData2, "expected listed JSON data to include testData2")
	})

	// Test Delete method with JSON data
	t.Run("Delete JSON Data", func(t *testing.T) {
		// Delete an existing JSON record
		err := flashDB.Delete(id1)
		assert.NoError(t, err, "expected no error when deleting an existing JSON record")

		// Verify the JSON record was deleted
		_, err = flashDB.Get(id1)
		assert.Error(t, err, "expected an error when retrieving a deleted JSON record")

		// Attempt to delete a non-existent JSON record
		err = flashDB.Delete(nonExistentID)
		assert.Error(t, err, "expected an error when deleting a non-existent JSON record")
	})
}
