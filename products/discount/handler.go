package discount

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	// r.Get("/prices", h.RouteList)
	// r.Post("/prices", h.RouteCreate)
	// r.Get("/prices/:id", h.RouteGet)
	// r.Post("/prices/:id", h.RouteUpdate)
}

// RouteGet func gets Discount by given ID or 404 error.
// @Description Get Discount by given ID or 404 error.
// @Summary get Discount by given ID or 404 error.
// @Tags Discount
// @Accept json
// @Produce json
// @Param id path string true "Discount ID"
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Discount
// @Router /v1/customers/{customer_id}/discount/{id} [get]
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

// RouteList func gets all existing Discounts.
// @Description Get all existing Discounts.
// @Summary get all existing Discounts
// @Tags Discount
// @Accept json
// @Produce json
// @Param customer_id path string true "Customer ID"
// @Success 200 {array} Discount
// @Router /v1/customers/{customer_id}/discount [get]
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

// RouteCreate func for creates a new Discount.
// @Description Create a new Discount.
// @Summary create a new Discount
// @Tags Discount
// @Accept json
// @Produce json
// @Param model body discount.Bind.request true "Request Data"
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Discount
// @Router /v1/customers/{customer_id}/discount [post]
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

// RouteUpdate func for updates Discount by given ID.
// @Description Update Discount.
// @Summary update Discount
// @Tags Discount
// @Accept json
// @Produce json
// @Param model body discount.Bind.request true "Request Data"
// @Param id path string true "Discount ID"
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Discount
// @Router /v1/customers/{customer_id}/discount/{id} [post]
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

// RouteDelete func for deletes Discount by given ID.
// @Description Delete Discount by given ID.
// @Summary delete Discount by given ID
// @Tags Discount
// @Accept json
// @Produce json
// @Param id path string true "Discount ID"
// @Param customer_id path string true "Customer ID"
// @Success 204 {string} status "ok"
// @Router /v1/customers/{customer_id}/discount/{id} [delete]
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
