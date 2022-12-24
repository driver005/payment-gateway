package service

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
)

type Cart struct {
	h   Handler
	ctx context.Context
}

func NewCart(r Registry) *Cart {
	return &Cart{
		h:   Handler{r: r},
		ctx: context.Background(),
	}
}

func (c *Cart) SetContext(ctx context.Context) *Cart {
	c.ctx = ctx
	return c
}

func (c *Cart) List(Id uuid.UUID, config repository.FindOption) ([]models.Cart, error) {
	m := models.Cart{}

	r, err := c.h.r.Repository().Find(m, config)
	if err != nil {
		return nil, err
	}

	return r.([]models.Cart), nil
}

func (c *Cart) Retrieve(Id uuid.UUID, config repository.FindOption) (*models.Cart, error) {
	m := models.Cart{}
	m.Id = Id

	r, err := c.h.r.Repository().FindOne(m, config)
	if err != nil {
		return nil, err
	}

	return r.(*models.Cart), nil
}

func (c *Cart) RetrieveNew(Id uuid.UUID, config repository.FindOption) (*models.Cart, error) {
	m := models.Cart{}
	m.Id = Id

	r, err := c.h.r.Repository().FindOne(m, config)
	if err != nil {
		return nil, err
	}

	return r.(*models.Cart), nil
}

// func (c *Cart) prepareCartForProcessing() (*models.Cart, error) {

// }
