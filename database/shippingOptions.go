package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetShippingOption(ctx context.Context, id uuid.UUID) (*models.ShippingOption, error) {
	var m models.ShippingOption

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetShippingOptions(ctx context.Context, filters models.Filter) ([]models.ShippingOption, *int64, error) {
	var m = make([]models.ShippingOption, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.ShippingOption{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ShippingOption{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateShippingOption(ctx context.Context, m *models.ShippingOption) (*models.ShippingOption, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateShippingOption(ctx context.Context, m *models.ShippingOption) (*models.ShippingOption, error) {
	o, err := h.GetShippingOption(ctx, m.Id)
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

func (h *Handler) DeleteShippingOption(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ShippingOption{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
