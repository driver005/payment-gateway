package method

import (
	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/payment/method/details"
	"github.com/driver005/gateway/payment/method/methods"
	"github.com/driver005/gateway/payment/method/options"
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Migrate() {
	details.Migrate(h.r)
	methods.Migrate(h.r)
	options.Migrate(h.r)

	err := h.r.Context().AutoMigrate(&PaymentMethod{})
	if err != nil {
		panic(err)
	}
}

func (s *Handler) Retrive(ctx *fiber.Ctx, Id uuid.UUID) (*PaymentMethod, error) {
	var m PaymentMethod

	if err := s.r.Manager(ctx.Context()).Model(&PaymentMethod{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) List(ctx *fiber.Ctx, Offset int, Size int) ([]PaymentMethod, *int64, error) {
	var m = make([]PaymentMethod, 0)

	r := s.r.Manager(ctx.Context()).Model(&PaymentMethod{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) Create(ctx *fiber.Ctx, m *PaymentMethod) (*PaymentMethod, error) {
	if err := s.r.Manager(ctx.Context()).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) Update(ctx *fiber.Ctx, m *PaymentMethod) (*PaymentMethod, error) {
	var o *PaymentMethod

	o, err := s.Retrive(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx.Context()).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) Upsert(ctx *fiber.Ctx, m *PaymentMethod) (*PaymentMethod, error) {
	res := r.r.Manager(ctx.Context()).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) Delete(ctx *fiber.Ctx, Id uuid.UUID) error {
	if err := s.r.Manager(ctx.Context()).Where("id = ?", Id).Delete(&PaymentMethod{}).Error; err != nil {
		return err
	}

	return nil
}
