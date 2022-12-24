package service

import (
	"context"

	"github.com/google/uuid"
)

type TaxProvider struct {
	h   Handler
	ctx context.Context
}

func NewTaxProvider(r Registry) *TaxProvider {
	return &TaxProvider{
		h:   Handler{r: r},
		ctx: context.Background(),
	}
}

func (t *TaxProvider) ClearLineItemsTaxLines(Ids []uuid.UUID) error {
	for _, values := range Ids {
		err := t.h.r.Repository().Delete(values)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TaxProvider) ClearTaxLines(Ids []uuid.UUID) error {
	for _, values := range Ids {
		err := t.h.r.Repository().Delete(values)
		if err != nil {
			return err
		}
	}

	return nil
}
