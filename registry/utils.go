package registry

import (
	"github.com/driver005/gateway/utils"
)

func (m *Base) Utils() *utils.Handler {
	if m.utils == nil {
		m.utils = utils.NewHandler(m.r)
	}
	return m.utils
}
