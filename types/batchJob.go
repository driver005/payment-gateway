package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterableBatchJobProps struct {
	core.Filter
	Status    models.BatchJobStatus `json:"status,omitempty"`
	CreatedBy string                `json:"created_by,omitempty"`
}
