package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetProductVariant(ctx context.Context, config types.FilterableProductVariantProps, model models.ProductVariant) (*models.ProductVariant, error) {
	var m models.ProductVariant

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetProductVariants(ctx context.Context, filters models.Filter, config types.FilterableProductVariantProps, model models.ProductVariant) ([]models.ProductVariant, *int64, error) {
	var m = make([]models.ProductVariant, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ProductVariant{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ProductVariant{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) GetProductVariantsWithProductID(ctx context.Context, Id uuid.UUID, filters models.Filter, config types.FilterableProductVariantProps, model models.ProductVariant) ([]models.ProductVariant, *int64, error) {
	var m = make([]models.ProductVariant, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ProductVariant{ProductId: uuid.NullUUID{UUID: Id, Valid: true}})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ProductVariant{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateProductVariant(ctx context.Context, Id uuid.UUID, m *models.ProductVariant) (*models.ProductVariant, error) {

	m.ProductId = uuid.NullUUID{UUID: Id, Valid: true}
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateProductVariant(ctx context.Context, VariantId uuid.UUID, ProductId uuid.UUID, m *models.ProductVariant) (*models.ProductVariant, error) {
	var o models.ProductVariant

	if err := h.r.Manager(ctx).Where("id = ? AND product_id = ?", VariantId, ProductId).First(&o).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := h.r.Manager(ctx).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteProductVariant(ctx context.Context, VariantId uuid.UUID, ProductId uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ? and product_id = ?", VariantId, ProductId).Delete(&models.ProductVariant{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
