package admin

import (
	"strconv"

	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

type Admin struct {
	r handler.Registry
}

func (a *Admin) GetDiscountCondition(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDiscountCondition(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (p *Admin) ListDiscountCondition(context *fiber.Ctx) error {
	page, _ := strconv.Atoi(context.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	m, n, err := p.r.ClientManager().GetDiscountConditions(context.Context(), models.Filter{
		Size:   pageSize,
		Offset: offset,
	})
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) CreateDiscountCondition(context *fiber.Ctx) error {
	var m models.DiscountCondition

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := a.r.ClientManager().CreateDiscountCondition(context.Context(), &m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateDiscountCondition(context *fiber.Ctx) error {
	var m models.DiscountCondition

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Id = Id

	r, err := a.r.ClientManager().UpdateDiscountCondition(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

func (a *Admin) DeleteDiscountCondition(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteDiscountCondition(context.Context(), Id)
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}

// type AddBatchDiscountConditionsRequest struct {
// 	Resources map[string]interface{} `json:"resources"`
// }

// func (a *Admin) UpdateBatchofResourcesDiscountCondition(context *fiber.Ctx) error {
// 	var b AddBatchDiscountConditionsRequest

// 	if err := context.BodyParser(&b); err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	Id, err := uuid.FromString(context.Params("id"))
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	ConditionId, err := uuid.FromString(context.Params("condition_id"))
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	r, err := a.r.ClientManager().UpdateDiscountCondition(context.Context(), &m)
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	return context.Status(fiber.StatusOK).JSON(r)
// }

// func (a *Admin) DeleteBatchofResourcesDiscountCondition(context *fiber.Ctx) error {
// 	var b AddBatchDiscountConditionsRequest

// 	if err := context.BodyParser(&b); err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	Id, err := uuid.FromString(context.Params("id"))
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	ConditionId, err := uuid.FromString(context.Params("condition_id"))
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	r, err := a.r.ClientManager().UpdateDiscountCondition(context.Context(), &m)
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "database_error",
// 		})
// 	}

// 	return context.Status(fiber.StatusOK).JSON(r)
// }
