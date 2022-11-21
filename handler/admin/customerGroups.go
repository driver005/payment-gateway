package admin

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func (a *Admin) removeCustomer(slice []models.Customer, s int) []models.Customer {
	return append(slice[:s], slice[s+1:]...)
}

func (a *Admin) GetCustomerGroup(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetCustomerGroup(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) ListCustomerGroup(context *fiber.Ctx) error {
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

	m, n, err := a.r.ClientManager().GetCustomerGroups(context.Context(), models.Filter{
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

func (a *Admin) CreateCustomerGroup(context *fiber.Ctx) error {
	var m models.CustomerGroup

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := a.r.ClientManager().CreateCustomerGroup(context.Context(), &m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateCustomerGroup(context *fiber.Ctx) error {
	var m models.CustomerGroup

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

	r, err := a.r.ClientManager().UpdateCustomerGroup(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

func (a *Admin) DeleteCustomerGroup(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteCustomerGroup(context.Context(), Id)
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}

type UpdateCustomerGroupCustomerBody struct {
	CustomerIds []string `json:"customer_ids"`
}

func (a *Admin) UpdateCustomerGroupCustomers(context *fiber.Ctx) error {
	var m models.CustomerGroup
	var b UpdateCustomerGroupCustomerBody

	if err := context.BodyParser(&b); err != nil {
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

	p, err := a.r.ClientManager().GetCustomerByIds(context.Context(), b.CustomerIds)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Id = Id
	m.Customers = append(m.Customers, p...)

	r, err := a.r.ClientManager().UpdateCustomerGroup(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

type DeleteCustomerGroupCustomerBody struct {
	CustomerIds []string `json:"customer_ids"`
}

func (a *Admin) DeleteCustomerGroupCustomers(context *fiber.Ctx) error {
	var m models.CustomerGroup
	var b DeleteCustomerGroupCustomerBody

	if err := context.BodyParser(&b); err != nil {
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

	p, err := a.r.ClientManager().GetCustomerByIds(context.Context(), b.CustomerIds)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	for i, pd := range p {
		if m.Customers[i].Id == pd.Id {
			m.Customers = a.removeCustomer(m.Customers, i)
		}
	}
	m.Id = Id

	r, err := a.r.ClientManager().UpdateCustomerGroup(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}
