package registry

import (
	"context"

	"github.com/driver005/gateway/database"
	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/hydra/x"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Base struct {
	l            *logger.Logger
	al           *logger.Logger
	h            *handler.Handler
	buildVersion string
	buildHash    string
	buildDate    string
	r            Registry
	trc          trace.Tracer
	database     database.Database
	writer       herodot.Writer
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

func (m *Base) Writer() herodot.Writer {
	if m.writer == nil {
		h := herodot.NewJSONWriter(m.Logger())
		h.ErrorEnhancer = x.ErrorEnhancer
		m.writer = h
	}
	return m.writer
}

func (m *Base) RegisterRoutes(public *httprouter.Router) {
	group := helper.NewRouteGroup(public, "/api")
	m.Handler().NbxplorerRoutes(group)
	m.Handler().BtcpayRoutes(group)
}

func (m *Base) Handler() *handler.Handler {
	if m.h == nil {
		m.h = handler.NewHandler(m.r, "localhost:32838")
	}
	return m.h
}

func (m *Base) Database() database.Database {
	return m.database
}
