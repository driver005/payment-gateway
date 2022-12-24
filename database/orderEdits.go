package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

func (h *Handler) GetOrderEdit(ctx context.Context, config core.Filter, model models.OrderEdit) (*models.OrderEdit, error) {
	var m models.OrderEdit

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetOrderEdits(ctx context.Context, filters models.Filter, config core.Filter, model models.OrderEdit) ([]models.OrderEdit, *int64, error) {
	var m = make([]models.OrderEdit, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.OrderEdit{})
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
	q := models.OrderEdit{}
	q.Id = m.Id

	o, err := h.GetOrderEdit(ctx, core.Filter{}, q)
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
