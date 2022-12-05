package repository

import "github.com/driver005/database"

// Repository is a generic DB handler that cares about default error handling
type Repository interface {
	SetPreloads(preloads ...string) *Repositories
	Query(target interface{}, query string, values ...interface{}) (interface{}, error)
	Get(target interface{}, config interface{}, model interface{}) (interface{}, error)
	GetId(target interface{}) interface{}
	HasId(target interface{}) bool
	Create(target interface{}) (interface{}, error)
	CreateBatch(target interface{}, batchSize int) (interface{}, error)
	Insert(target interface{}) (interface{}, error)
	Save(target interface{}) (interface{}, error)
	Delete(target interface{}) error
	DeletePermanently(target interface{}) error
	Remove(target interface{}) error
	RemoveTX(target interface{}, tx *database.DB) error
	Preload(target interface{}, preloads ...string) (interface{}, error)
	Upsert(target interface{}) (interface{}, error)

	GetAll(target interface{}) error
	GetBatch(target interface{}, limit, offset int) error
	GetWhere(target interface{}, condition string) error
	GetWhereBatch(target interface{}, condition string, limit, offset int) error
	GetByField(target interface{}, field string, value interface{}) error
	GetByFields(target interface{}, filters map[string]interface{}) error
	GetByFieldBatch(target interface{}, field string, value interface{}, limit, offset int) error
	GetByFieldsBatch(target interface{}, filters map[string]interface{}, limit, offset int) error
	GetOneByField(target interface{}, field string, value interface{}) error
	GetOneByFields(target interface{}, filters map[string]interface{}) error
	GetOneByID(target interface{}, id string) error

	DB() *database.DB
	HandleError(res *database.DB) error
	HandleOneError(res *database.DB) error
}

// TransactionRepository extends Repository with modifier functions that accept a transaction
type TransactionRepository interface {
	Repository
	CreateTx(target interface{}, tx *database.DB) (interface{}, error)
	InsertTx(target interface{}, tx *database.DB) (interface{}, error)
	SaveTx(target interface{}, tx *database.DB) (interface{}, error)
	DeleteTx(target interface{}, tx *database.DB) error
	// CreateTx(target interface{}, tx *database.DB) error
	// SaveTx(target interface{}, tx *database.DB) error
	// DeleteTx(target interface{}, tx *database.DB) error
}
