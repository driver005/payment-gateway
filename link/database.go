package link

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/link/options"
	"github.com/google/uuid"
)

func (h *Handler) Migrate() {
	options.Migrate(h.r)

	err := h.r.Context().AutoMigrate(
		&PaymentLink{},
	)
	if err != nil {
		panic(err)
	}
}

func (s *Handler) Retrive(ctx context.Context, Id uuid.UUID) (*PaymentLink, error) {
	var m PaymentLink

	if err := s.r.Manager(ctx).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) List(ctx context.Context, Offset int, Size int) ([]PaymentLink, *int64, error) {
	var m = make([]PaymentLink, 0)

	r := s.r.Manager(ctx).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) Create(ctx context.Context, m *PaymentLink) (*PaymentLink, error) {
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) Update(ctx context.Context, m *PaymentLink) (*PaymentLink, error) {
	var o *PaymentLink

	o, err := s.Retrive(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) Upsert(ctx context.Context, m *PaymentLink) (*PaymentLink, error) {
	res := r.r.Manager(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) Delete(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&PaymentLink{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *Handler) DeleteTax(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&PaymentLink{}).Error; err != nil {
		return err
	}

	return nil
}