package subscriptionSchedule

import (
	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Migrate() {
	err := h.r.Context().AutoMigrate(
		&SubscriptionSchedule{},
		&SubscriptionSchedulesResourceDefaultSettings{},
		&SubscriptionScheduleSubscriptionItem{},
		&SubscriptionSchedulePhaseConfiguration{},
	)
	if err != nil {
		panic(err)
	}
}

func (s *Handler) Retrive(ctx *fiber.Ctx, Id uuid.UUID) (*SubscriptionSchedule, error) {
	var m SubscriptionSchedule

	if err := s.r.Manager(ctx.Context()).Model(&SubscriptionSchedule{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) List(ctx *fiber.Ctx, Offset int, Size int) ([]SubscriptionSchedule, *int64, error) {
	var m = make([]SubscriptionSchedule, 0)

	r := s.r.Manager(ctx.Context()).Model(&SubscriptionSchedule{}).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) Create(ctx *fiber.Ctx, m *SubscriptionSchedule) (*SubscriptionSchedule, error) {
	if err := s.r.Manager(ctx.Context()).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) Update(ctx *fiber.Ctx, m *SubscriptionSchedule) (*SubscriptionSchedule, error) {
	var o *SubscriptionSchedule

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

func (r *Handler) Upsert(ctx *fiber.Ctx, m *SubscriptionSchedule) (*SubscriptionSchedule, error) {
	res := r.r.Manager(ctx.Context()).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) Delete(ctx *fiber.Ctx, Id uuid.UUID) error {
	if err := s.r.Manager(ctx.Context()).Where("id = ?", Id).Delete(&SubscriptionSchedule{}).Error; err != nil {
		return err
	}

	return nil
}
