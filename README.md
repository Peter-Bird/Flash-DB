# FlashDB

FlashDB is an in-memory database implementation written in Go, designed for lightweight data storage and retrieval in real-time applications. It implements a `Repository` interface, providing essential CRUD operations.

## Features

- **In-memory storage**: Fast, ephemeral storage ideal for testing or temporary data.
- **Concurrency support**: Built with thread-safe mechanisms using Go's `sync.RWMutex`.
- **CRUD operations**: Includes functions to list, retrieve, save, delete, and truncate data.

## Getting Started

### Prerequisites

- Go 1.18+ installed on your system.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Peter-Bird/Flash-DB.git
   cd Flash-DB
   ```

2. Include the package in your Go project:

   ```go
   import "github.com/Peter-Bird/Flash-DB"
   ```

### Usage

#### Initializing the Database

```go
package main

import (
    "fmt"
    "github.com/Peter-Bird/Flash-DB"
)

func main() {
    db := FlashDB.NewFlashDB()

    // Add some data
    err := db.Save("1", map[string]interface{}{"name": "Example"})
    if err != nil {
        fmt.Println("Error saving data:", err)
    }

    // List all records
    data, _ := db.List()
    fmt.Println("All data:", data)
}
```

#### API

##### `List() ([]map[string]interface{}, error)`
Returns all records stored in the database.

##### `Get(id string) (map[string]interface{}, error)`
Retrieves a specific record by ID.

##### `Save(id string, data map[string]interface{}) error`
Saves or updates a record with the given ID.

##### `Delete(id string) error`
Deletes a record by ID.

##### `Truncate() error`
Clears all records in the database.

## Contributing

Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the need for simple, lightweight data storage for Go applications.
