// github.com/Peter-Bird/Flash-DB
// pkg/db/repository.go

package FlashDB

type Repository interface {
	List() ([]map[string]interface{}, error)

	Get(id string) (map[string]interface{}, error)
	Save(id string, data map[string]interface{}) error
	Delete(id string) error
	Truncate() error
}
