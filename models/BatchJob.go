package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// BatchJob - A Batch Job.
type BatchJob struct {
	core.Model

	// The type of batch job.
	Type string `json:"type"`

	// The status of the batch job.
	Status BatchJobStatus `json:"status" database:"default:created"`

	// The unique identifier of the user that created the batch job.
	CreatedBy string `json:"created_by" database:"default:null"`

	// A user object. Available if the relation `created_by_user` is expanded.
	CreatedByUser *User `json:"created_by_user" database:"foreignKey:id"`

	// The context of the batch job, the type of the batch job determines what the context should contain.
	Context JSONB `json:"context" database:"default:null"`

	// Specify if the job must apply the modifications or not.
	DryRun bool `json:"dry_run" database:"default:null"`

	Result BatchJobResult `json:"result" database:"default:null"`

	// The date from which the job has been pre processed.
	PreProcessedAt time.Time `json:"pre_processed_at" database:"default:null"`

	// The date the job is processing at.
	ProcessingAt time.Time `json:"processing_at" database:"default:null"`

	// The date when the confirmation has been done.
	ConfirmedAt time.Time `json:"confirmed_at" database:"default:null"`

	// The date of the completion.
	CompletedAt time.Time `json:"completed_at" database:"default:null"`

	// The date of the concellation.
	CanceledAt time.Time `json:"canceled_at" database:"default:null"`

	// The date when the job failed.
	FailedAt time.Time `json:"failed_at" database:"default:null"`
}
