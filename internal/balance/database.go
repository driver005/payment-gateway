package balance

import (
	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v2"
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

func (s *Handler) RetriveBalance(ctx *fiber.Ctx, Id uuid.UUID) (*Balance, error) {
	var m Balance

	if err := s.r.Manager(ctx.Context()).Model(&Balance{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) ListBalance(ctx *fiber.Ctx, Offset int, Size int) ([]Balance, *int64, error) {
	var m = make([]Balance, 0)

	r := s.r.Manager(ctx.Context()).Model(&Balance{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) CreateBalance(ctx *fiber.Ctx, m *Balance) (*Balance, error) {
	if err := s.r.Manager(ctx.Context()).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) UpdateBalance(ctx *fiber.Ctx, m *Balance) (*Balance, error) {
	var o *Balance

	o, err := s.RetriveBalance(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx.Context()).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) UpsertBalance(ctx *fiber.Ctx, m *Balance) (*Balance, error) {
	res := r.r.Manager(ctx.Context()).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) DeleteBalance(ctx *fiber.Ctx, Id uuid.UUID) error {
	if err := s.r.Manager(ctx.Context()).Where("id = ?", Id).Delete(&Balance{}).Error; err != nil {
		return err
	}

	return nil
}

// BalanceTransaction

func (s *Handler) RetriveBalanceTransaction(ctx *fiber.Ctx, Id uuid.UUID) (*BalanceTransaction, error) {
	var m BalanceTransaction

	if err := s.r.Manager(ctx.Context()).Model(&BalanceTransaction{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) ListBalanceTransaction(ctx *fiber.Ctx, Offset int, Size int) ([]BalanceTransaction, *int64, error) {
	var m = make([]BalanceTransaction, 0)

	r := s.r.Manager(ctx.Context()).Model(&BalanceTransaction{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) CreateBalanceTransaction(ctx *fiber.Ctx, m *BalanceTransaction) (*BalanceTransaction, error) {
	if err := s.r.Manager(ctx.Context()).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) UpdateBalanceTransaction(ctx *fiber.Ctx, m *BalanceTransaction) (*BalanceTransaction, error) {
	var o *BalanceTransaction

	o, err := s.RetriveBalanceTransaction(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx.Context()).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) UpsertBalanceTransaction(ctx *fiber.Ctx, m *BalanceTransaction) (*BalanceTransaction, error) {
	res := r.r.Manager(ctx.Context()).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) DeleteBalanceTransaction(ctx *fiber.Ctx, Id uuid.UUID) error {
	if err := s.r.Manager(ctx.Context()).Where("id = ?", Id).Delete(&BalanceTransaction{}).Error; err != nil {
		return err
	}

	return nil
}
