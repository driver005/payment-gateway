package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetProductTag(ctx context.Context, config core.Filter, model models.ProductTag) (*models.ProductTag, error) {
	var m models.ProductTag

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetProductTags(ctx context.Context, filters models.Filter, config core.Filter, model models.ProductTag) ([]models.ProductTag, *int64, error) {
	var m = make([]models.ProductTag, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ProductTag{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ProductTag{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateProductTag(ctx context.Context, m *models.ProductTag) (*models.ProductTag, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateProductTag(ctx context.Context, m *models.ProductTag) (*models.ProductTag, error) {
	q := models.ProductTag{}
	q.Id = m.Id

	o, err := h.GetProductTag(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteProductTag(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ProductTag{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetProductTagCount(ctx context.Context) ([]models.ProductTag, *int64, error) {
	var m = make([]models.ProductTag, 0)
	var count int64

	if err := h.r.Manager(ctx).Order("id").Find(&m).Count(&count).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	return m, &count, nil
}

func (h *Handler) GetProductTagsByIds(ctx context.Context, Ids []string) ([]models.ProductTag, error) {
	var m = make([]models.ProductTag, 0)

	for _, p := range Ids {
		Id, err := uuid.FromString(p)
		if err != nil {
			return nil, helper.ParseError(err)
		}

		q := models.ProductTag{}
		q.Id = Id

		r, err := h.GetProductTag(ctx, core.Filter{}, q)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
