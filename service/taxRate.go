package service

import (
	"context"
	"errors"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
)

type TaxRate struct {
	h   Handler
	ctx context.Context
}

func NewTaxRate(r Registry) *TaxRate {
	return &TaxRate{
		h:   Handler{r: r},
		ctx: context.Background(),
	}
}

func (t *TaxRate) List(config repository.FindOption) ([]models.TaxRate, error) {
	m := models.TaxRate{}

	r, err := t.h.r.Repository().Find(m, config)
	if err != nil {
		return nil, err
	}

	return r.([]models.TaxRate), nil
}

func (t *TaxRate) Retrieve(Id uuid.UUID, config repository.FindOption) (*models.TaxRate, error) {
	m := models.TaxRate{}
	m.Id = Id

	r, err := t.h.r.Repository().FindOne(m, config)
	if err != nil {
		return nil, err
	}

	return r.(*models.TaxRate), nil
}

func (t *TaxRate) ListAndCount(config repository.FindOption) ([]models.TaxRate, *int64, error) {
	m := models.TaxRate{}

	r, count, err := t.h.r.Repository().FindAndCount(m, config)
	if err != nil {
		return nil, nil, err
	}

	return r.([]models.TaxRate), count, nil
}

func (t *TaxRate) Create(model models.TaxRate) (*models.TaxRate, error) {

	if model.RegionId.UUID == uuid.Nil {
		return nil, errors.New("taxRates must belong to a Region")
	}

	r, err := t.h.r.Repository().Create(model)
	if err != nil {
		return nil, err
	}
	return r.(*models.TaxRate), nil
}

func (b *TaxRate) Update(model models.TaxRate) (*models.TaxRate, error) {
	r, err := b.h.r.Repository().Update(model)
	if err != nil {
		return nil, err
	}

	return r.(*models.TaxRate), nil
}

func (b *TaxRate) Delete(model models.TaxRate) error {
	err := b.h.r.Repository().Delete(model)
	if err != nil {
		return err
	}

	return nil
}

func (b *TaxRate) RemoveFromProduct(Id uuid.UUID, productIds []uuid.UUID) error {
	for _, value := range productIds {
		err := b.h.r.Repository().Delete(models.ProductTaxRate{
			ProductId: uuid.NullUUID{UUID: value, Valid: true},
			RateId:    uuid.NullUUID{UUID: Id, Valid: true},
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func (b *TaxRate) RemoveFromProductType(Id uuid.UUID, typeIds []uuid.UUID) error {
	for _, value := range typeIds {
		err := b.h.r.Repository().Delete(models.ProductTypeTaxRate{
			ProductTypeId: uuid.NullUUID{UUID: value, Valid: true},
			RateId:        uuid.NullUUID{UUID: Id, Valid: true},
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func (b *TaxRate) RemoveFromShippingOption(Id uuid.UUID, optionIds []uuid.UUID) error {
	for _, value := range optionIds {
		err := b.h.r.Repository().Delete(models.ShippingTaxRate{
			ShippingOptionId: uuid.NullUUID{UUID: value, Valid: true},
			RateId:           uuid.NullUUID{UUID: Id, Valid: true},
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TaxRate) AddToProduct(Id uuid.UUID, productIds []uuid.UUID, replace bool) ([]models.ProductTaxRate, error) {
	// for _, value := range productIds {
	// 	if replace {
	// 		t.h.r.Repository().
	// 	} else {

	// 	}
	// }
	return nil, nil
}
