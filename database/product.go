package database

import (
	"context"
	"strings"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetProduct(ctx context.Context, config types.FilterableProductProps, model models.Product) (*models.Product, error) {
	var m models.Product

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetProducts(ctx context.Context, filters models.Filter, config types.FilterableProductProps, model models.Product) ([]models.Product, *int64, error) {
	var m = make([]models.Product, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.Product{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Product{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateProduct(ctx context.Context, m *models.Product) (*models.Product, error) {

	if m.Handle == "" {
		m.Handle = strings.ToLower(m.Title)
	}

	if m.Thumbnail == "" && len(m.Images) > 0 {
		m.Thumbnail = m.Images[0].Url
	}

	if m.IsGiftcard == true {
		s, err := h.GetByTypeShippingProfile(ctx, "gift_card")
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m.ProfileId = uuid.NullUUID{UUID: s.Id, Valid: true}
	} else {
		s, err := h.GetByTypeShippingProfile(ctx, "default")
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m.ProfileId = uuid.NullUUID{UUID: s.Id, Valid: true}
	}

	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt

	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateProduct(ctx context.Context, m *models.Product) (*models.Product, error) {
	q := models.Product{}
	q.Id = m.Id

	o, err := h.GetProduct(ctx, types.FilterableProductProps{}, q)
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

func (h *Handler) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetProductsByIds(ctx context.Context, ProductIds []string) ([]models.Product, error) {
	var m = make([]models.Product, 0)

	for _, p := range ProductIds {
		Id, err := uuid.FromString(p)
		if err != nil {
			return nil, helper.ParseError(err)
		}

		q := models.Product{}
		q.Id = Id

		r, err := h.GetProduct(ctx, types.FilterableProductProps{}, q)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
