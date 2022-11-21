package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetImage(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	var m models.Image

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetImages(ctx context.Context, filters models.Filter) ([]models.Image, *int64, error) {
	var m = make([]models.Image, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.Image{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Image{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateImage(ctx context.Context, m *models.Image) (*models.Image, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateImage(ctx context.Context, m *models.Image) (*models.Image, error) {
	o, err := h.GetImage(ctx, m.Id)
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

func (h *Handler) DeleteImage(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.Image{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
