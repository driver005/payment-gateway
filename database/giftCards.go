package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetGiftCard(ctx context.Context, id uuid.UUID) (*models.GiftCard, error) {
	var m models.GiftCard

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetGiftCards(ctx context.Context, filters models.Filter) ([]models.GiftCard, *int64, error) {
	var m = make([]models.GiftCard, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.GiftCard{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.GiftCard{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateGiftCard(ctx context.Context, m *models.GiftCard) (*models.GiftCard, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateGiftCard(ctx context.Context, m *models.GiftCard) (*models.GiftCard, error) {
	o, err := h.GetGiftCard(ctx, m.Id)
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

func (h *Handler) DeleteGiftCard(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.GiftCard{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
