package repository

import (

	// "github.com/aklinkert/go-logging"
	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/helper"
	"github.com/fatih/structs"
	"github.com/lib/pq"
)

type Repositories struct {
	db           *database.DB
	defaultJoins pq.StringArray
	preloads     pq.StringArray
}

// NewRepositories returns a new base repositories that implements TransactionRepositories
func NewRepositories(db *database.DB, defaultJoins ...string) TransactionRepository {
	return &Repositories{
		defaultJoins: defaultJoins,
		db:           db,
		preloads:     pq.StringArray{},
	}
}

func (r *Repositories) DB() *database.DB {
	dbConn := r.db

	for _, join := range r.defaultJoins {
		dbConn = dbConn.Joins(join)
	}

	for _, preload := range r.preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}

func (r *Repositories) SetPreloads(preloads ...string) *Repositories {
	r.preloads = preloads

	return r
}

func (r *Repositories) HandleError(res *database.DB) error {
	if res.Error != nil && res.Error != database.ErrRecordNotFound {
		return helper.ParseError(res.Error)
	}

	return nil
}

func (r *Repositories) Query(target interface{}, query string, values ...interface{}) (interface{}, error) {
	res := r.DB().Raw(query, values...).Scan(&target)

	return target, r.HandleError(res)
}

func (r *Repositories) Get(target interface{}, config interface{}, model interface{}) (interface{}, error) {

	res := r.ToQuery(config, model).Find(&target)

	return target, r.HandleError(res)
}

func (r *Repositories) GetId(target interface{}) interface{} {
	s := structs.New(target)

	f, ok := s.FieldOk("Id")
	if !ok {
		return nil
	}
	return f.Value()
}

func (r *Repositories) HasId(target interface{}) bool {
	s := structs.New(target)

	_, ok := s.FieldOk("Id")

	return ok
}

func (r *Repositories) Create(target interface{}) (interface{}, error) {
	res := r.db.Create(&target)
	return target, r.HandleError(res)
}

func (r *Repositories) CreateBatch(target interface{}, batchSize int) (interface{}, error) {
	res := r.db.CreateInBatches(&target, batchSize)
	return target, r.HandleError(res)
}

func (r *Repositories) CreateTx(target interface{}, tx *database.DB) (interface{}, error) {
	res := tx.Create(&target)
	return target, r.HandleError(res)
}

func (r *Repositories) Insert(target interface{}) (interface{}, error) {
	return r.Create(&target)
}

func (r *Repositories) InsertTx(target interface{}, tx *database.DB) (interface{}, error) {
	return r.CreateTx(&target, tx)
}

func (r *Repositories) Save(target interface{}) (interface{}, error) {
	res := r.db.Save(&target)
	return target, r.HandleError(res)
}

func (r *Repositories) SaveTx(target interface{}, tx *database.DB) (interface{}, error) {
	res := tx.Save(&target)
	return target, r.HandleError(res)
}

func (r *Repositories) Update(target interface{}) (interface{}, error) {
	m, err := r.Get(target, nil, target)

	if err != nil {
		return nil, err
	}

	res := r.db.Model(m).Updates(target)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return res, nil
}

func (r *Repositories) Delete(target interface{}) error {
	res := r.db.Delete(target)
	return r.HandleError(res)
}

func (r *Repositories) DeleteTx(target interface{}, tx *database.DB) error {
	res := tx.Delete(target)
	return r.HandleError(res)
}

func (r *Repositories) DeletePermanently(target interface{}) error {
	res := r.db.Unscoped().Delete(target)
	return r.HandleError(res)
}

func (r *Repositories) Remove(target interface{}) error {
	return r.Delete(target)
}

func (r *Repositories) RemoveTX(target interface{}, tx *database.DB) error {
	return r.DeleteTx(target, tx)
}

func (r *Repositories) Preload(target interface{}, preloads ...string) (interface{}, error) {
	dbConn := r.db

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	res := dbConn.Find(&target)
	if res.Error != nil {
		s := dbConn.Save(&target)
		return &target, r.HandleError(s)
	}
	return &target, r.HandleError(res)
}

func (r *Repositories) Upsert(target interface{}) (interface{}, error) {
	res := r.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&target)
	return target, r.HandleError(res)
}

func (r *Repositories) SoftRemove(target interface{}) error {
	return r.Remove(target)
}

func (r *Repositories) Recover(target interface{}) (interface{}, error) {
	res := r.db.Model(target).Update("deleted_at", nil)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return res, nil
}

func (r *Repositories) Count(target interface{}, condition FindOption) (*int64, error) {
	var count int64

	res := r.BuildFindQuery(condition).Find(&target).Count(&count)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return &count, nil
}

func (r *Repositories) CountBy(target interface{}, condition string) (*int64, error) {
	res := r.DB().Model(&target).Where(condition)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return &res.RowsAffected, nil
}

func (r *Repositories) Find(target interface{}, condition FindOption) (interface{}, error) {
	res := r.BuildFindQuery(condition).Find(&target)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return &target, nil
}

func (r *Repositories) FindBy(target interface{}, condition map[string]interface{}) (interface{}, error) {
	res := r.BuildFindWhereQuery(condition).Find(&target)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return &target, nil
}

func (r *Repositories) FindAndCount(target interface{}, condition FindOption) (interface{}, *int64, error) {
	res := r.BuildFindQuery(condition).Find(&target)

	if res.Error != nil {
		return nil, nil, r.HandleError(res)
	}

	return &target, &res.RowsAffected, nil
}

func (r *Repositories) FindAndCountBy(target interface{}, condition map[string]interface{}) (interface{}, *int64, error) {
	res := r.BuildFindWhereQuery(condition).Find(&target)

	if res.Error != nil {
		return nil, nil, r.HandleError(res)
	}

	return &target, &res.RowsAffected, nil
}

func (r *Repositories) FindOne(target interface{}, condition FindOption) (interface{}, error) {
	res := r.BuildFindQuery(condition).First(&target)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return &target, nil
}

func (r *Repositories) FindOneBy(target interface{}, condition map[string]interface{}) (interface{}, error) {
	res := r.BuildFindWhereQuery(condition).First(&target)

	if res.Error != nil {
		return nil, r.HandleError(res)
	}

	return &target, nil
}

func (r *Repositories) Clear(target interface{}) error {
	res := r.db.Migrator().DropTable(target)

	if res != nil {
		return res
	}

	return nil
}

func (r *Repositories) HandleOneError(res *database.DB) error {
	if err := r.HandleError(res); err != nil {
		return err
	}

	if res.RowsAffected != 1 {
		return ErrNotFound
	}

	return nil
}

// func (r *Repositories) GetAll(target interface{}) error {

// 	res := r.DB().
// 		Unscoped().
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetBatch(target interface{}, limit, offset int) error {

// 	res := r.DB().
// 		Unscoped().
// 		Limit(limit).
// 		Offset(offset).
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetWhere(target interface{}, condition string) error {

// 	res := r.DB().
// 		Where(condition).
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetWhereBatch(target interface{}, condition string, limit, offset int) error {

// 	res := r.DB().
// 		Where(condition).
// 		Limit(limit).
// 		Offset(offset).
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetByField(target interface{}, field string, value interface{}) error {

// 	res := r.DB().
// 		Where(fmt.Sprintf("%v = ?", field), value).
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetByFields(target interface{}, filters map[string]interface{}) error {

// 	db := r.DB()
// 	for field, value := range filters {
// 		db = db.Where(fmt.Sprintf("%v = ?", field), value)
// 	}

// 	res := db.Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetByFieldBatch(target interface{}, field string, value interface{}, limit, offset int) error {

// 	res := r.DB().
// 		Where(fmt.Sprintf("%v = ?", field), value).
// 		Limit(limit).
// 		Offset(offset).
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetByFieldsBatch(target interface{}, filters map[string]interface{}, limit, offset int) error {

// 	db := r.DB()
// 	for field, value := range filters {
// 		db = db.Where(fmt.Sprintf("%v = ?", field), value)
// 	}

// 	res := db.
// 		Limit(limit).
// 		Offset(offset).
// 		Find(target)

// 	return r.HandleError(res)
// }

// func (r *Repositories) GetOneByField(target interface{}, field string, value interface{}) error {

// 	res := r.DB().
// 		Where(fmt.Sprintf("%v = ?", field), value).
// 		First(target)

// 	return r.HandleOneError(res)
// }

// func (r *Repositories) GetOneByFields(target interface{}, filters map[string]interface{}) error {

// 	db := r.DB()
// 	for field, value := range filters {
// 		db = db.Where(fmt.Sprintf("%v = ?", field), value)
// 	}

// 	res := db.First(target)
// 	return r.HandleOneError(res)
// }

// func (r *Repositories) GetOneByID(target interface{}, id string) error {

// 	res := r.DB().
// 		Where("id = ?", id).
// 		First(target)

// 	return r.HandleOneError(res)
// }
