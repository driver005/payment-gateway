package admin

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func (a *Admin) GetCustomer(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetCustomer(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) ListCustomer(context *fiber.Ctx) error {
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

	m, n, err := a.r.ClientManager().GetCustomers(context.Context(), models.Filter{
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

func (a *Admin) CreateCustomer(context *fiber.Ctx) error {
	var m models.Customer

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := a.r.ClientManager().CreateCustomer(context.Context(), &m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateCustomer(context *fiber.Ctx) error {
	var m models.Customer

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

	r, err := a.r.ClientManager().UpdateCustomer(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

func (a *Admin) DeleteCustomer(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteCustomer(context.Context(), Id)
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}
