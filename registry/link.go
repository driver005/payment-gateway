package registry

import "github.com/driver005/gateway/link"

func (m *Base) Link() *link.Handler {
	if m.link == nil {
		m.link = link.NewHandler(m.r)
	}
	return m.link
}
