package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetSalesChannel(ctx context.Context, id uuid.UUID) (*models.SalesChannel, error) {
	var m models.SalesChannel

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetSalesChannels(ctx context.Context, filters models.Filter) ([]models.SalesChannel, *int64, error) {
	var m = make([]models.SalesChannel, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.SalesChannel{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.SalesChannel{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateSalesChannel(ctx context.Context, m *models.SalesChannel) (*models.SalesChannel, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateSalesChannel(ctx context.Context, m *models.SalesChannel) (*models.SalesChannel, error) {
	o, err := h.GetSalesChannel(ctx, m.Id)
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

func (h *Handler) DeleteSalesChannel(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.SalesChannel{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
