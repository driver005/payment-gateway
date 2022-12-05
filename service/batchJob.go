package service

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

type BatchJob struct {
	h   Handler
	ctx context.Context
}

func NewBatchJob(r Registry) *BatchJob {
	return &BatchJob{
		h:   Handler{r: r},
		ctx: context.Background(),
	}
}

func (b *BatchJob) setContext(ctx context.Context) *BatchJob {
	b.ctx = ctx
	return b
}

func (b *BatchJob) retrieve(Id uuid.UUID, config types.FilterableBatchJobProps) (*models.BatchJob, error) {
	m := models.BatchJob{}
	m.Id = Id

	r, err := b.h.r.Repository().Get(m, config, m)
	if err != nil {
		return nil, err
	}

	return r.(*models.BatchJob), nil
}

func (b *BatchJob) listAndCount(config types.FilterableBatchJobProps, model models.BatchJob) ([]models.BatchJob, *int64, error) {
	r, count, err := b.h.r.ClientManager().GetBatchJobCount(b.ctx, config, model)
	if err != nil {
		return nil, nil, err
	}

	return r, count, nil
}

func (b *BatchJob) create(model *models.BatchJob) (*models.BatchJob, error) {
	r, err := b.h.r.ClientManager().CreateBatchJob(b.ctx, model)
	if err != nil {
		return nil, err
	}

	return r, nil
}
