package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetCustomer(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	var m models.Customer

	if err := h.r.Manager(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetCustomers(ctx context.Context, filters models.Filter) ([]models.Customer, *int64, error) {
	var m = make([]models.Customer, 0)

	if err := h.r.Manager(ctx).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.r.Manager(ctx).Find(&models.Customer{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Customer{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateCustomer(ctx context.Context, m *models.Customer) (*models.Customer, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateCustomer(ctx context.Context, m *models.Customer) (*models.Customer, error) {
	o, err := h.GetCustomer(ctx, m.Id)
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

func (h *Handler) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.Customer{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetCustomerByIds(ctx context.Context, CustomerIds []string) ([]models.Customer, error) {
	var m = make([]models.Customer, 0)

	for _, p := range CustomerIds {
		Id, err := uuid.FromString(p)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		r, err := h.GetCustomer(ctx, Id)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
