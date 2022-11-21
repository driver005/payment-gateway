package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetOrderEdit(ctx context.Context, id uuid.UUID) (*models.OrderEdit, error) {
	var m models.OrderEdit

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetOrderEdits(ctx context.Context, filters models.Filter) ([]models.OrderEdit, *int64, error) {
	var m = make([]models.OrderEdit, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.OrderEdit{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.OrderEdit{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateOrderEdit(ctx context.Context, m *models.OrderEdit) (*models.OrderEdit, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateOrderEdit(ctx context.Context, m *models.OrderEdit) (*models.OrderEdit, error) {
	o, err := h.GetOrderEdit(ctx, m.Id)
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

func (h *Handler) DeleteOrderEdit(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.OrderEdit{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
