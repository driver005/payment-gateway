package registry

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/database/driver/postgres"
	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/helper"
	_ "github.com/lib/pq"
)

type SQL struct {
	*Base
}

var _ Registry = new(SQL)

func init() {
	RegisterDriver(func() Driver {
		return NewRegistrySQL()
	})
}

func NewRegistrySQL() *SQL {
	r := &SQL{
		Base: new(Base),
	}
	r.Base.with(r)
	return r
}

func (m *SQL) Init(ctx context.Context) error {
	if m.database == nil {
		// new db connection
		c := postgres.New(postgres.Config{
			DriverName: "postgres",
			DSN: "postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqtc",
		})
		
		database, err := database.Open(c, &database.Config{})
		
		if err != nil {
			return helper.WithStack(err)
		}


		m.database, err = db.NewManager(database, m.l)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *SQL) CanHandle(dsn string) bool {
	return m.alwaysCanHandle(dsn)
}

func (m *SQL) alwaysCanHandle(dsn string) bool {
	scheme := helper.Split(dsn, "://")[0]
	s := helper.Canonicalize(scheme)
	return s == helper.DriverMySQL || s == helper.DriverPostgreSQL
}

func (m *SQL) Context() *database.DB {
	return m.database.DB(context.Background())
}

func (m *SQL) Manager(ctx context.Context) *database.DB {
	return m.database.DB(ctx)
}
