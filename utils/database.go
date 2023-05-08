package utils

import (
	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/sql"
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Retrive[T any](s *Handler, ctx *fiber.Ctx, Id uuid.UUID) (*T, error) {
	var m T
	var t T

	if err := s.r.Manager(ctx.Context()).Model(&t).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func List[T any](s *Handler, ctx *fiber.Ctx, Offset int, Size int) ([]T, *int64, error) {
	var m = make([]T, 0)
	var t T

	r := s.r.Manager(ctx.Context()).Model(&t).Scopes(sql.FilterByQuery(ctx, sql.ALL)).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func Create[T any](s *Handler, ctx *fiber.Ctx, m *T) (*T, error) {
	if err := s.r.Manager(ctx.Context()).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func Update[T any](s *Handler, ctx *fiber.Ctx, m T) (*T, error) {
	var o *T

	new := structs.New(m)
	old := structs.New(o)

	newIdName := new.Field("Id")
	oldIdName := old.Field("Id")

	// newIdValue := newIdName.Value().(uuid.UUID)
	oldIdValue := oldIdName.Value().(uuid.UUID)

	o, err := Retrive[T](s, ctx, oldIdValue)
	if err != nil {
		return nil, err
	}

	newIdName.Set(oldIdValue)
	if err := s.r.Manager(ctx.Context()).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func Upsert[T any](s *Handler, ctx *fiber.Ctx, m *T) (*T, error) {
	res := s.r.Manager(ctx.Context()).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func Delete[T any](s *Handler, ctx *fiber.Ctx, Id uuid.UUID) error {
	var m T
	if err := s.r.Manager(ctx.Context()).Where("id = ?", Id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
