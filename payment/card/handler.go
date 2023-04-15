package card

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/customers/:id/sources", h.RouteList)
	r.Post("/customers/:id/sources", h.RouteCreate)
	r.Get("/customers/:id/sources/:id", h.RouteGet)
	r.Post("/customers/:id/sources/:id", h.RouteUpdate)
	r.Delete("/customers/:id/sources/:id", h.RouteUpdate)
}

// RouteGet func gets Card by given ID or 404 error.
// @Description Get Card by given ID or 404 error.
// @Summary get Card by given ID or 404 error.
// @Tags Card
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Card
// @Router /v1/customers/{customer_id}/sources/{id} [get]
func (h *Handler) RouteGet(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m, err := h.Retrive(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteList func gets all existing Cards.
// @Description Get all existing Cards.
// @Summary get all existing Cards
// @Tags Card
// @Accept json
// @Produce json
// @Param customer_id path string true "Customer ID"
// @Success 200 {array} Card
// @Router /v1/customers/{customer_id}/sources [get]
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

	m, n, err := h.List(context.Context(), offset, pageSize)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteCreate func for creates a new Card.
// @Description Create a new Card.
// @Summary create a new Card
// @Tags Card
// @Accept json
// @Produce json
// @Param model body card.Bind.request true "Request Data"
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Card
// @Router /v1/customers/{customer_id}/sources [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	m, err := h.Bind(context)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	r, err := h.Create(context.Context(), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteUpdate func for updates Card by given ID.
// @Description Update Card.
// @Summary update Card
// @Tags Card
// @Accept json
// @Produce json
// @Param model body card.Bind.request true "Request Data"
// @Param id path string true "Card ID"
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Card
// @Router /v1/customers/{customer_id}/sources/{id} [post]
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

	r, err := h.Update(context.Context(), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDelete func for deletes Card by given ID.
// @Description Delete Card by given ID.
// @Summary delete Card by given ID
// @Tags Card
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Param customer_id path string true "Customer ID"
// @Success 204 {string} status "ok"
// @Router /v1/customers/{customer_id}/sources/{id} [delete]
func (h *Handler) RouteDelete(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	err = h.Delete(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}
