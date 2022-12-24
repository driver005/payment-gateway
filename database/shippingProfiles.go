package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

func (h *Handler) GetShippingProfile(ctx context.Context, config core.Filter, model models.ShippingProfile) (*models.ShippingProfile, error) {
	var m models.ShippingProfile

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetShippingProfiles(ctx context.Context, filters models.Filter, config core.Filter, model models.ShippingProfile) ([]models.ShippingProfile, *int64, error) {
	var m = make([]models.ShippingProfile, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ShippingProfile{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ShippingProfile{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateShippingProfile(ctx context.Context, m *models.ShippingProfile) (*models.ShippingProfile, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateShippingProfile(ctx context.Context, m *models.ShippingProfile) (*models.ShippingProfile, error) {
	q := models.ShippingProfile{}
	q.Id = m.Id

	o, err := h.GetShippingProfile(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteShippingProfile(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ShippingProfile{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetByTypeShippingProfile(ctx context.Context, shippingProfileType string) (*models.ShippingProfile, error) {
	var m models.ShippingProfile

	if err := h.r.Manager(ctx).Where("type = ?", shippingProfileType).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GenerateDefaultShippingProfile(ctx context.Context) ([]models.ShippingProfile, error) {
	var m = make([]models.ShippingProfile, 0)
	var d models.ShippingProfile
	var g models.ShippingProfile

	d.Name = "Default Shipping Profile"
	d.Type = "default"

	g.Name = "Gift Card Profile"
	g.Type = "gift_card"

	r, err := h.CreateShippingProfile(ctx, &d)
	if err != nil {
		return nil, err
	}
	m = append(m, *r)

	r, err = h.CreateShippingProfile(ctx, &g)
	if err != nil {
		return nil, err
	}

	m = append(m, *r)

	return m, nil
}
