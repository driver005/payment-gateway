package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetInvite(ctx context.Context, config core.Filter, model models.Invite) (*models.Invite, error) {
	var m models.Invite

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetInvites(ctx context.Context, filters models.Filter, config core.Filter, model models.Invite) ([]models.Invite, *int64, error) {
	var m = make([]models.Invite, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.Invite{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Invite{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateInvite(ctx context.Context, m *models.Invite) (*models.Invite, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateInvite(ctx context.Context, m *models.Invite) (*models.Invite, error) {
	q := models.Invite{}
	q.Id = m.Id

	o, err := h.GetInvite(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteInvite(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.Invite{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
