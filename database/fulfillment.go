package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetFulfillment(ctx context.Context, id uuid.UUID) (*models.Fulfillment, error) {
	var m models.Fulfillment

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetFulfillments(ctx context.Context, filters models.Filter) ([]models.Fulfillment, *int64, error) {
	var m = make([]models.Fulfillment, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.Fulfillment{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Fulfillment{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateFulfillment(ctx context.Context, m *models.Fulfillment) (*models.Fulfillment, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateFulfillment(ctx context.Context, m *models.Fulfillment) (*models.Fulfillment, error) {
	o, err := h.GetFulfillment(ctx, m.Id)
	if err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := h.r.Manager(ctx).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteFulfillment(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.Fulfillment{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
