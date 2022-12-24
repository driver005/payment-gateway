package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
)

func (h *Handler) GetPriceList(ctx context.Context, config types.FilterablePriceListProps, model models.PriceList) (*models.PriceList, error) {
	var m models.PriceList

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetPriceLists(ctx context.Context, filters models.Filter, config types.FilterablePriceListProps, model models.PriceList) ([]models.PriceList, *int64, error) {
	var m = make([]models.PriceList, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.PriceList{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.PriceList{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreatePriceList(ctx context.Context, m *models.PriceList) (*models.PriceList, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdatePriceList(ctx context.Context, m *models.PriceList) (*models.PriceList, error) {
	q := models.PriceList{}
	q.Id = m.Id

	o, err := h.GetPriceList(ctx, types.FilterablePriceListProps{}, q)
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

func (h *Handler) DeletePriceList(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.PriceList{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
