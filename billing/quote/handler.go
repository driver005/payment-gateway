package quote

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/quotes", h.RouteList)
	r.Post("/quotes", h.RouteCreate)
	r.Get("/quotes/:id", h.RouteGet)
	r.Post("/quotes/:id", h.RouteUpdate)
	r.Post("/quotes/:id/finalize", h.RouteFinalize)
	r.Post("/quotes/:id/accept", h.RouteAccept)
	r.Post("/quotes/:id/cancel", h.RouteCancel)
}

// RouteGet func gets Quote by given ID or 404 error.
// @Description Get Quote by given ID or 404 error.
// @Summary get Quote by given ID or 404 error.
// @Tags Quote
// @Accept json
// @Produce json
// @Param id path string true "Quote ID"
// @Success 200 {object} Quote
// @Router /v1/quotes/{id} [get]
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

// RouteList func gets all existing Quotes.
// @Description Get all existing Quotes.
// @Summary get all existing Quotes
// @Tags Quote
// @Accept json
// @Produce json
// @Success 200 {array} Quote
// @Router /v1/quotes [get]
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

// RouteCreate func for creates a new Quote.
// @Description Create a new Quote.
// @Summary create a new Quote
// @Tags Quote
// @Accept json
// @Produce json
// @Success 200 {object} Quote
// @Router /v1/quotes [post]
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

// RouteUpdate func for updates Quote by given ID.
// @Description Update Quote.
// @Summary update Quote
// @Tags Quote
// @Accept json
// @Produce json
// @Param id body string true "Quote ID"
// @Success 200 {object} Quote
// @Router /v1/quotes/{id} [post]
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

// RouteDelete func for deletes Quote by given ID.
// @Description Delete Quote by given ID.
// @Summary delete Quote by given ID
// @Tags Quote
// @Accept json
// @Produce json
// @Param id body string true "Quote ID"
// @Success 204 {string} status "ok"
// @Router /v1/quotes/{id} [delete]
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

// RouteFinalize func for finalizing Quote by given ID.
// @Description Finalize Quote.
// @Summary finalize Quote
// @Tags Quote
// @Accept json
// @Produce json
// @Param id body string true "Quote ID"
// @Success 200 {object} Quote
// @Router /v1/quotes/{id}/finalize [post]
func (h *Handler) RouteFinalize(context *fiber.Ctx) error {
	var m Quote

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "draft" {
		m.Status = "open"
	}

	r, err := h.Update(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteAccept func for accepting Quote by given ID.
// @Description Accept Quote.
// @Summary accept Quote
// @Tags Quote
// @Accept json
// @Produce json
// @Param id body string true "Quote ID"
// @Success 200 {object} Quote
// @Router /v1/quotes/{id}/accept [post]
func (h *Handler) RouteAccept(context *fiber.Ctx) error {
	var m Quote

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "open" {
		m.Status = "accepted"
	}

	r, err := h.Update(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteCancel func for canceling Quote by given ID.
// @Description Cancel Quote.
// @Summary cancel Quote
// @Tags Quote
// @Accept json
// @Produce json
// @Param id body string true "Quote ID"
// @Success 200 {object} Quote
// @Router /v1/quotes/{id}/cancel [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m Quote

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status != "draft" {
		m.Status = "canceled"
	}

	r, err := h.Update(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}
