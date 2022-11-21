package database

import (
	"context"
	"time"

	"github.com/driver005/database"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetBatchJob(ctx context.Context, id uuid.UUID) (*models.BatchJob, error) {
	var m models.BatchJob

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetBatchJobs(ctx context.Context, filters models.Filter) ([]models.BatchJob, *int64, error) {
	var m = make([]models.BatchJob, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.BatchJob{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.BatchJob{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateBatchJob(ctx context.Context, m *models.BatchJob) (*models.BatchJob, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Session(&database.Session{DryRun: m.DryRun}).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateBatchJob(ctx context.Context, m *models.BatchJob) (*models.BatchJob, error) {
	o, err := h.GetBatchJob(ctx, m.Id)
	if err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := h.r.Manager(ctx).Session(&database.Session{DryRun: m.DryRun}).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteBatchJob(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.BatchJob{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
