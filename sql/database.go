package sql

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Sql interface {
	Connect(string, string, string) pg.DB
	Create(...interface{}) orm.Result
}

type Database struct {
	db *pg.DB
}

type Story struct {
    Id       int64
    Title    string
    AuthorId int64
}

func (d *Database) Connect(user string, password string, database string) *pg.DB {
	d.db = pg.Connect(&pg.Options{
        User:     user,
		Password: password,
		Database: database,
    })
    defer d.db.Close()

	return d.db
}

func (d *Database) Create(model ...interface{}) orm.Result {
	var err error
	var response orm.Result
	response, err = d.db.Model(model).Insert()
    if err != nil {
        panic(err)
    }

	return response
}