package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetLineItem(ctx context.Context, config types.FilterableLineItemProps, model models.LineItem) (*models.LineItem, error) {
	var m models.LineItem

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetLineItems(ctx context.Context, filters models.Filter, config types.FilterableLineItemProps, model models.LineItem) ([]models.LineItem, *int64, error) {
	var m = make([]models.LineItem, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.LineItem{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.LineItem{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateLineItem(ctx context.Context, m *models.LineItem) (*models.LineItem, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateLineItem(ctx context.Context, m *models.LineItem) (*models.LineItem, error) {
	q := models.LineItem{}
	q.Id = m.Id

	o, err := h.GetLineItem(ctx, types.FilterableLineItemProps{}, q)
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

func (h *Handler) DeleteLineItem(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.LineItem{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
