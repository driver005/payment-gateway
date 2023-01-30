package source

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/payment/source/mandate"
	"github.com/driver005/gateway/payment/source/transaction"
	"github.com/driver005/gateway/payment/source/types"
	"github.com/google/uuid"
)

func (h *Handler) Migrate() {
	mandate.Migrate(h.r)
	transaction.Migrate(h.r)
	types.Migrate(h.r)

	err := h.r.Context().AutoMigrate(
		&Source{},
		&SourceRedirectFlow{},
		&SourceReceiverFlow{},
		&SourceOwner{},
		&SourceOrderItem{},
		&SourceOrder{},
		&SourceCodeVerificationFlow{},
	)
	if err != nil {
		panic(err)
	}
}

func (s *Handler) Retrive(ctx context.Context, Id uuid.UUID) (*Source, error) {
	var m Source

	if err := s.r.Manager(ctx).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) List(ctx context.Context, Offset int, Size int) ([]Source, *int64, error) {
	var m = make([]Source, 0)

	r := s.r.Manager(ctx).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) Create(ctx context.Context, m *Source) (*Source, error) {
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) Update(ctx context.Context, m *Source) (*Source, error) {
	var o *Source

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

func (r *Handler) Upsert(ctx context.Context, m *Source) (*Source, error) {
	res := r.r.Manager(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) Delete(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&Source{}).Error; err != nil {
		return err
	}

	return nil
}