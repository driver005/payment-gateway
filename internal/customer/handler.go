package customer

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/customers", h.RouteList)
	r.Post("/customers", h.RouteCreate)
	r.Get("/customers/:id", h.RouteGet)
	r.Post("/customers/:id", h.RouteUpdate)
	r.Get("/customers/:id/tax_ids", h.RouteListTax)
	r.Post("/customers/:id/tax_ids", h.RouteCreateTax)
	r.Get("/customers/:id/tax_ids/:tax_id", h.RouteGetTax)
	r.Delete("/customers/:id/tax_ids/:tax_id", h.RouteDeleteTax)
}

// RouteGet func gets Customer by given ID or 404 error.
// @Description Get Customer by given ID or 404 error.
// @Summary get Customer by given ID or 404 error.
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /v1/customers/{id} [get]
func (h *Handler) RouteGet(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m, err := h.Retrive(context, Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteList func gets all existing Customers.
// @Description Get all existing Customers.
// @Summary get all existing Customers
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {array} Customer
// @Router /v1/customers [get]
func (h *Handler) RouteList(context *fiber.Ctx) error {
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

	m, n, err := h.List(context, offset, pageSize)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteCreate func for creates a new Customer.
// @Description Create a new Customer.
// @Summary create a new Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param model body Customer true "Request Data"
// @Success 200 {object} Customer
// @Router /v1/customers [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	m, err := h.Bind(context)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	r, err := h.Create(context, m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteUpdate func for updates Customer by given ID.
// @Description Update Customer.
// @Summary update Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param model body Customer true "Request Data"
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /v1/customers/{id} [post]
func (h *Handler) RouteUpdate(context *fiber.Ctx) error {
	m, err := h.Bind(context)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id

	r, err := h.Update(context, m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDelete func for deletes Customer by given ID.
// @Description Delete Customer by given ID.
// @Summary delete Customer by given ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 204 {string} status "ok"
// @Router /v1/customers/{id} [delete]
func (h *Handler) RouteDelete(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	err = h.Delete(context, Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

// RouteGetTax func gets Customer by given ID or 404 error.
// @Description Get Customer by given ID or 404 error.
// @Summary get Customer by given ID or 404 error.
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param tax_id path string true "Tax ID"
// @Success 200 {object} Customer
// @Router /v1/customers/{id}/tax_ids/{tax_id} [get]
func (h *Handler) RouteGetTax(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	_, err = uuid.Parse(context.Params("tax_id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m, err := h.Retrive(context, Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	// idx := slices.IndexFunc(m.Tax, func(c CustomerTax) bool { return c.Id == TaxId })

	// return context.Status(fiber.StatusOK).JSON(&m.Tax[idx])

	return context.Status(fiber.StatusOK).JSON(&m.Tax)
}

// RouteListTax func gets all existing Customers.
// @Description Get all existing Customers.
// @Summary get all existing Customers
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {array} Customer
// @Router /v1/customers/{id}/tax_ids [get]
func (h *Handler) RouteListTax(context *fiber.Ctx) error {
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

	m, n, err := h.List(context, offset, pageSize)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteCreateTax func for creates a new Customer.
// @Description Create a new Customer.
// @Summary create a new Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param model body Customer true "Request Data"
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /v1/customers/{id}/tax_ids [post]
func (h *Handler) RouteCreateTax(context *fiber.Ctx) error {
	var m Customer

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := h.Create(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteUpdateTax func for updates Customer by given ID.
// @Description Update Customer.
// @Summary update Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param model body Customer true "Request Data"
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /v1/customers/{id} [post]
func (h *Handler) RouteUpdateTax(context *fiber.Ctx) error {
	var m Customer

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id

	r, err := h.Update(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDeleteTax func for deletes Customer by given ID.
// @Description Delete Customer by given ID.
// @Summary delete Customer by given ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param tax_id path string true "Tax ID"
// @Success 204 {string} status "ok"
// @Router /v1/customers/{id}/tax_ids/{tax_id} [delete]
func (h *Handler) RouteDeleteTax(context *fiber.Ctx) error {
	// Id, err := uuid.Parse(context.Params("id"))
	// if err != nil {
	// 	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 		"type":    "parse_error",
	// 	})
	// }

	TaxId, err := uuid.Parse(context.Params("tax_id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	err = h.DeleteTax(context, TaxId)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}
