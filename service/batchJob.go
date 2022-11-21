package service

import (
	"context"

	"github.com/driver005/gateway/database"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

type BatchJob struct {
	db  database.Handler
	ctx context.Context
}

func (b *BatchJob) setContext(ctx context.Context) *BatchJob {
	b.ctx = ctx
	return b
}

func (b *BatchJob) retrieve(batchJobId uuid.UUID) (*models.BatchJob, error) {
	bj, err := b.db.GetBatchJob(b.ctx, batchJobId)
	if err != nil {
		return nil, err
	}

	return bj, nil
}

func (b *BatchJob) listAndCount(batchJobId uuid.UUID) (*models.BatchJob, error) {
	bj, err := b.db.GetBatchJob(b.ctx, batchJobId)
	if err != nil {
		return nil, err
	}

	return bj, nil
}
