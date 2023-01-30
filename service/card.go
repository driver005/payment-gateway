package service

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type Card struct {
	r Registry
}

func NewCard(r Registry) *Card {
	return &Card{
		r: r,
	}
}

func (s *Card) Retrive(ctx context.Context, Id uuid.UUID) (*models.Card, error) {
	var m models.Card

	if err := s.r.Manager(ctx).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (s *Card) List(ctx context.Context, Offset int, Size int) ([]models.Card, *int64, error) {
	var m = make([]models.Card, 0)

	r := s.r.Manager(ctx).Scopes(Paginate(Offset, Size)).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, helper.ParseError(r.Error)
	}

	return m, &r.RowsAffected, nil
}

func (s *Card) Create(ctx context.Context, m *models.Card) (*models.Card, error) {
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (s *Card) Update(ctx context.Context, Id uuid.UUID) (*models.Card, error) {
	var m models.Card

	o, err := s.Retrive(ctx, Id)
	if err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Model(&o).Updates(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (s *Card) Delete(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&models.Card{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
