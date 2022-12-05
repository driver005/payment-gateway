package registry

import (
	"context"

	"github.com/driver005/gateway/database"
	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/repository"

	// "github.com/driver005/gateway/service"
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Base struct {
	l  *logger.Logger
	al *logger.Logger
	h  *handler.Handler
	// a  *admin.Handler
	// s            *service.Handler
	buildVersion string
	buildHash    string
	buildDate    string
	r            Registry
	trc          trace.Tracer
	database     sql.Database
	db           *database.Handler
	rp           repository.TransactionRepository
}

func (m *Base) with(r Registry) *Base {
	m.r = r
	return m
}

func (m *Base) WithBuildInfo(version, hash, date string) Registry {
	m.buildVersion = version
	m.buildHash = hash
	m.buildDate = date
	return m.r
}

func (m *Base) BuildVersion() string {
	return m.buildVersion
}

func (m *Base) BuildDate() string {
	return m.buildDate
}

func (m *Base) BuildHash() string {
	return m.buildHash
}

func (m *Base) WithLogger(l *logger.Logger) Registry {
	m.l = l
	return m.r
}

func (m *Base) Logger() *logger.Logger {
	if m.l == nil {
		m.l = logger.New("ORY Hydra", m.BuildVersion())
	}
	return m.l
}

func (m *Base) Tracer(ctx context.Context) trace.Tracer {
	if m.trc == nil {
		tp := otel.GetTracerProvider()
		m.trc = tp.Tracer("github.com/driver005/gateway", trace.WithInstrumentationVersion(m.BuildVersion()))
	}

	return m.trc
}

func (m *Base) AuditLogger() *logger.Logger {
	if m.al == nil {
		m.al = logger.NewAudit("ORY Hydra", m.BuildVersion())
	}
	return m.al
}

func (m *Base) RegisterRoutes(router *fiber.App) {
	group := router.Group("/api")
	m.Handler().NbxplorerRoutes(group)
	m.Handler().BtcpayRoutes(group)
	// m.Admin().Routes(group)
}

func (m *Base) Handler() *handler.Handler {
	if m.h == nil {
		m.h = handler.NewHandler(m.r, "localhost:32838")
	}
	return m.h
}

func (m *Base) Db() *database.Handler {
	if m.db == nil {
		m.db = database.NewHandler(m.r)
	}
	return m.db
}

func (m *Base) Repository() repository.TransactionRepository {
	if m.rp == nil {
		m.rp = repository.NewRepositories(m.Database().DB(context.Background()))
	}
	return m.rp
}

// func (m *Base) Admin() *admin.Handler {
// 	if m.a == nil {
// 		m.a = admin.NewHandler(m.r)
// 	}
// 	return m.a
// }

// func (m *Base) Service() *service.Handler {
// 	if m.s == nil {
// 		m.s = service.NewHandler(m.r)
// 	}
// 	return m.s
// }

func (m *Base) Database() sql.Database {
	return m.database
}

func (m *Base) Setup() {
	m.r.ClientManager().GenerateDefaultShippingProfile(context.Background())
}
