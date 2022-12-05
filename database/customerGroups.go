package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetCustomerGroup(ctx context.Context, config types.FilterableCustomerGroupProps, model models.CustomerGroup) (*models.CustomerGroup, error) {
	var m models.CustomerGroup

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetCustomerGroups(ctx context.Context, filters models.Filter, config types.FilterableCustomerGroupProps, model models.CustomerGroup) ([]models.CustomerGroup, *int64, error) {
	var m = make([]models.CustomerGroup, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.CustomerGroup{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.CustomerGroup{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateCustomerGroup(ctx context.Context, m *models.CustomerGroup) (*models.CustomerGroup, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateCustomerGroup(ctx context.Context, m *models.CustomerGroup) (*models.CustomerGroup, error) {
	q := models.CustomerGroup{}
	q.Id = m.Id

	o, err := h.GetCustomerGroup(ctx, types.FilterableCustomerGroupProps{}, q)
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

func (h *Handler) DeleteCustomerGroup(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.CustomerGroup{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetCustomerGroupsByIds(ctx context.Context, Ids []string) ([]models.CustomerGroup, error) {
	var m = make([]models.CustomerGroup, 0)

	for _, p := range Ids {
		Id, err := uuid.FromString(p)
		if err != nil {
			return nil, helper.ParseError(err)
		}

		q := models.CustomerGroup{}
		q.Id = Id

		r, err := h.GetCustomerGroup(ctx, types.FilterableCustomerGroupProps{}, q)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
