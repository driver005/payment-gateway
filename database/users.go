package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetUser(ctx context.Context, config types.FilterableUserProps, model models.User) (*models.User, error) {
	var m models.User

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetUsers(ctx context.Context, filters models.Filter, config types.FilterableUserProps, model models.User) ([]models.User, *int64, error) {
	var m = make([]models.User, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.User{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.User{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateUser(ctx context.Context, m *models.User) (*models.User, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateUser(ctx context.Context, m *models.User) (*models.User, error) {
	q := models.User{}
	q.Id = m.Id

	o, err := h.GetUser(ctx, types.FilterableUserProps{}, q)
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

func (h *Handler) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
