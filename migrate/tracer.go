package migrate

import (
	"go.opentelemetry.io/otel/trace"
)

// IsLoaded returns true if the tracer has been loaded.
func (m *Migrator) IsLoaded() bool {
	if m.tracer == nil || m.tracer == nil {
		return false
	}
	return true
}

// Returns the wrapped tracer.
func (m *Migrator) Tracer() trace.Tracer {
	return m.tracer
}