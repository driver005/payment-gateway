package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
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

func (b *BatchJob) SetContext(ctx context.Context) *BatchJob {
	b.ctx = ctx
	return b
}

func (b *BatchJob) Retrieve(Id uuid.UUID, config repository.FindOption) (*models.BatchJob, error) {
	m := models.BatchJob{}
	m.Id = Id

	r, err := b.h.r.Repository().FindOne(m, config)
	if err != nil {
		return nil, err
	}

	return r.(*models.BatchJob), nil
}

func (b *BatchJob) ListAndCount(config repository.FindOption) ([]models.BatchJob, *int64, error) {
	m := models.BatchJob{}

	r, count, err := b.h.r.Repository().FindAndCount(m, config)
	if err != nil {
		return nil, nil, err
	}

	return r.([]models.BatchJob), count, nil
}

func (b *BatchJob) Create(model *models.BatchJob) (*models.BatchJob, error) {
	r, err := b.h.r.Repository().Create(model)
	if err != nil {
		return nil, err
	}

	bus := EventBus.New()
	bus.SubscribeAsync("batch.created", map[string]interface{}{"id": r.(*models.BatchJob).Id}, false)

	return r.(*models.BatchJob), nil
}

func (b *BatchJob) Update(model *models.BatchJob) (*models.BatchJob, error) {
	r, err := b.h.r.Repository().Update(model)
	if err != nil {
		return nil, err
	}

	bus := EventBus.New()
	bus.SubscribeAsync("batch.updated", map[string]interface{}{"id": r.(*models.BatchJob).Id}, false)

	return r.(*models.BatchJob), nil
}

func (b *BatchJob) UpdateStatus(model *models.BatchJob) (*models.BatchJob, error) {
	r, err := b.h.r.Repository().Update(model)
	if err != nil {
		return nil, err
	}

	bus := EventBus.New()
	bus.SubscribeAsync(fmt.Sprintf("batch.%+v", r.(*models.BatchJob).Status), map[string]interface{}{"id": r.(*models.BatchJob).Id}, false)

	return r.(*models.BatchJob), nil
}

func (b *BatchJob) Confirm(Id uuid.UUID) (*models.BatchJob, error) {
	m, err := b.Retrieve(Id, repository.FindOption{})
	if err != nil {
		return nil, err
	}

	if m.Status != models.BatchJobStatusPreProcessed {
		return nil, errors.New(fmt.Sprintf(`Cannot complete a batch job with status "%+v The batch job must be processing`, string(m.Status)))
	}

	m.Status = models.BatchJobStatusConfirmed

	r, err := b.UpdateStatus(m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (b *BatchJob) Complete(Id uuid.UUID) (*models.BatchJob, error) {
	m, err := b.Retrieve(Id, repository.FindOption{})
	if err != nil {
		return nil, err
	}

	if m.Status != models.BatchJobStatusPreProcessed {
		return nil, fmt.Errorf(`Cannot complete a batch job with status "%+v The batch job must be processing`, m.Status)
	}

	m.Status = models.BatchJobStatusCompleted

	r, err := b.UpdateStatus(m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (b *BatchJob) Cancel(Id uuid.UUID) (*models.BatchJob, error) {
	m, err := b.Retrieve(Id, repository.FindOption{})
	if err != nil {
		return nil, err
	}

	if m.Status == models.BatchJobStatusCompleted {
		return nil, errors.New("cannot cancel completed batch job")
	}

	m.Status = models.BatchJobStatusCanceled

	r, err := b.UpdateStatus(m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (b *BatchJob) SetPreProcessingDone(Id uuid.UUID) (*models.BatchJob, error) {
	m, err := b.Retrieve(Id, repository.FindOption{})
	if err != nil {
		return nil, err
	}

	if m.Status == models.BatchJobStatusPreProcessed {
		return m, nil
	}

	if m.Status != models.BatchJobStatusCreated {
		return nil, errors.New("cannot mark a batch job as pre processed if it is not in created status")
	}

	m.Status = models.BatchJobStatusPreProcessed

	r, err := b.UpdateStatus(m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (b *BatchJob) SetProcessing(Id uuid.UUID) (*models.BatchJob, error) {
	m, err := b.Retrieve(Id, repository.FindOption{})
	if err != nil {
		return nil, err
	}

	if m.Status != models.BatchJobStatusConfirmed {
		return nil, errors.New("cannot mark a batch job as processing if the status is different that confirmed")
	}

	m.Status = models.BatchJobStatusConfirmed

	r, err := b.UpdateStatus(m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (b *BatchJob) SetFailed(Id uuid.UUID) (*models.BatchJob, error) {
	m, err := b.Retrieve(Id, repository.FindOption{})
	if err != nil {
		return nil, err
	}

	m.Status = models.BatchJobStatusFailed

	r, err := b.UpdateStatus(m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// func (b *BatchJob) prepareBatchJobForProcessing() (*models.BatchJob, error) {

// }
