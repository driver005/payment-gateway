package balance

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/google/uuid"
)

func (h *Handler) Migrate() {
	err := h.r.Context().AutoMigrate(
		&Balance{},
		&BalanceAmount{},
		&BalanceAmountBySourceType{},
		&BalanceDetail{},
		&BalanceTransaction{},
	)
	if err != nil {
		panic(err)
	}
}

func (s *Handler) RetriveBalance(ctx context.Context, Id uuid.UUID) (*Balance, error) {
	var m Balance

	if err := s.r.Manager(ctx).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) ListBalance(ctx context.Context, Offset int, Size int) ([]Balance, *int64, error) {
	var m = make([]Balance, 0)

	r := s.r.Manager(ctx).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) CreateBalance(ctx context.Context, m *Balance) (*Balance, error) {
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) UpdateBalance(ctx context.Context, m *Balance) (*Balance, error) {
	var o *Balance

	o, err := s.RetriveBalance(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) UpsertBalance(ctx context.Context, m *Balance) (*Balance, error) {
	res := r.r.Manager(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) DeleteBalance(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&Balance{}).Error; err != nil {
		return err
	}

	return nil
}

// BalanceTransaction

func (s *Handler) RetriveBalanceTransaction(ctx context.Context, Id uuid.UUID) (*BalanceTransaction, error) {
	var m BalanceTransaction

	if err := s.r.Manager(ctx).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) ListBalanceTransaction(ctx context.Context, Offset int, Size int) ([]BalanceTransaction, *int64, error) {
	var m = make([]BalanceTransaction, 0)

	r := s.r.Manager(ctx).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) CreateBalanceTransaction(ctx context.Context, m *BalanceTransaction) (*BalanceTransaction, error) {
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) UpdateBalanceTransaction(ctx context.Context, m *BalanceTransaction) (*BalanceTransaction, error) {
	var o *BalanceTransaction

	o, err := s.RetriveBalanceTransaction(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) UpsertBalanceTransaction(ctx context.Context, m *BalanceTransaction) (*BalanceTransaction, error) {
	res := r.r.Manager(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) DeleteBalanceTransaction(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&BalanceTransaction{}).Error; err != nil {
		return err
	}

	return nil
}
