package registry

import (
	"context"
	"fmt"
	"os"

	"github.com/driver005/database"
	"github.com/driver005/database/driver/postgres"
	"github.com/driver005/database/logger"
	"github.com/driver005/database/schema"
	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/sql"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type SQL struct {
	*Base

	migrate bool
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

func (m *SQL) Init(ctx context.Context, migrate bool) error {
	m.migrate = migrate

	if m.database == nil {
		// new db connection
		c := postgres.New(postgres.Config{
			DriverName: "postgres",
			// DSN:        "postgres://postgres:postgres@127.0.0.1:5432/medusa-docker?sslmode=disable",
			DSN: "postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqtc",
		})

		database, err := database.Open(c, &database.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			return helper.WithStack(err)
		}

		m.database, err = sql.NewManager(database, m.l)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *SQL) DeleteTabels(db *database.DB) error {
	tabels, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}

	w := wow.New(os.Stdout, spin.Get(spin.Earth), fmt.Sprintf("    Drop %+v tabels", len(tabels)))
	w.Start()
	for _, v := range tabels {
		err := db.Migrator().DropTable(v)
		if err != nil {
			return err
		}
	}

	w.PersistWith(spin.Spinner{Frames: pq.StringArray{"✅"}}, " Finished")

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

func (m *SQL) Migrate() bool {
	return m.migrate
}

func (m *SQL) Context() *database.DB {
	return m.database.DB(context.Background())
}

func (m *SQL) Manager(ctx context.Context) *database.DB {
	return m.database.DB(ctx)
}

func (m *SQL) ClientManager() *db.Handler {
	return m.Db()
}
