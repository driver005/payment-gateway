package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetClaim(ctx context.Context, id uuid.UUID) (*models.ClaimOrder, error) {
	var m models.ClaimOrder

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetClaims(ctx context.Context, filters models.Filter) ([]models.ClaimOrder, *int64, error) {
	var m = make([]models.ClaimOrder, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.ClaimOrder{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ClaimOrder{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateClaim(ctx context.Context, m *models.ClaimOrder) (*models.ClaimOrder, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateClaim(ctx context.Context, m *models.ClaimOrder) (*models.ClaimOrder, error) {
	o, err := h.GetClaim(ctx, m.Id)
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

func (h *Handler) UpdateClaimWithOrderId(ctx context.Context, OrderId uuid.UUID, m *models.ClaimOrder) (*models.ClaimOrder, error) {
	var o models.ClaimOrder

	if err := h.r.Manager(ctx).Where("id = ? AND product_id = ?", m.Id, OrderId).First(&o).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := h.r.Manager(ctx).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteClaim(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ClaimOrder{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
