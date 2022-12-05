package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetProductOption(ctx context.Context, config core.Filter, model models.ProductOption) (*models.ProductOption, error) {
	var m models.ProductOption

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetProductOptions(ctx context.Context, filters models.Filter, config core.Filter, model models.ProductOption) ([]models.ProductOption, *int64, error) {
	var m = make([]models.ProductOption, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ProductOption{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ProductOption{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateProductOption(ctx context.Context, Id uuid.UUID, m *models.ProductOption) (*models.ProductOption, error) {
	m.ProductId = uuid.NullUUID{UUID: Id, Valid: true}
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateProductOption(ctx context.Context, OptionId uuid.UUID, ProductId uuid.UUID, m *models.ProductOption) (*models.ProductOption, error) {
	var o models.ProductOption

	if err := h.r.Manager(ctx).Where("id = ? AND product_id = ?", OptionId, ProductId).First(&o).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := h.r.Manager(ctx).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteProductOption(ctx context.Context, OptionId uuid.UUID, ProductId uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ? and product_id = ?", OptionId, ProductId).Delete(&models.ProductOption{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
