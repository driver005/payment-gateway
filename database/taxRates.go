package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetTaxRate(ctx context.Context, config types.FilterableTaxRateProps, model models.TaxRate) (*models.TaxRate, error) {
	var m models.TaxRate

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetTaxRates(ctx context.Context, filters models.Filter, config types.FilterableTaxRateProps, model models.TaxRate) ([]models.TaxRate, *int64, error) {
	var m = make([]models.TaxRate, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.TaxRate{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.TaxRate{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateTaxRate(ctx context.Context, m *models.TaxRate) (*models.TaxRate, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateTaxRate(ctx context.Context, m *models.TaxRate) (*models.TaxRate, error) {
	q := models.TaxRate{}
	q.Id = m.Id

	o, err := h.GetTaxRate(ctx, types.FilterableTaxRateProps{}, q)
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

func (h *Handler) DeleteTaxRate(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.TaxRate{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
