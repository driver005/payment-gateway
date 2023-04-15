package checkout

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/checkout/sessions", h.RouteList)
	r.Post("/checkout/sessions", h.RouteCreate)
	r.Get("/checkout/sessions/:id", h.RouteGet)
	r.Post("/checkout/sessions/:id", h.RouteUpdate)
	r.Post("/checkout/sessions/:id/expire", h.RouteExpire)
}

// RouteGet func gets CheckoutSession by given ID or 404 error.
// @Description Get CheckoutSession by given ID or 404 error.
// @Summary get CheckoutSession by given ID or 404 error.
// @Tags CheckoutSession
// @Accept json
// @Produce json
// @Param id path string true "CheckoutSession ID"
// @Success 200 {object} CheckoutSession
// @Router /v1/checkout/sessions/{id} [get]
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

// RouteList func gets all existing CheckoutSessions.
// @Description Get all existing CheckoutSessions.
// @Summary get all existing CheckoutSessions
// @Tags CheckoutSession
// @Accept json
// @Produce json
// @Success 200 {array} CheckoutSession
// @Router /v1/checkout/sessions [get]
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

// RouteCreate func for creates a new CheckoutSession.
// @Description Create a new CheckoutSession.
// @Summary create a new CheckoutSession
// @Tags CheckoutSession
// @Accept json
// @Produce json
// @Param model body CheckoutSession true "Request Data"
// @Success 200 {object} CheckoutSession
// @Router /v1/checkout/sessions [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	var m CheckoutSession

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := h.Create(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteUpdate func for updates CheckoutSession by given ID.
// @Description Update CheckoutSession.
// @Summary update CheckoutSession
// @Tags CheckoutSession
// @Accept json
// @Produce json
// @Param model body CheckoutSession true "Request Data"
// @Param id path string true "CheckoutSession ID"
// @Success 200 {object} CheckoutSession
// @Router /v1/checkout/sessions/{id} [post]
func (h *Handler) RouteUpdate(context *fiber.Ctx) error {
	var m CheckoutSession

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

	r, err := h.Update(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDelete func for deletes CheckoutSession by given ID.
// @Description Delete CheckoutSession by given ID.
// @Summary delete CheckoutSession by given ID
// @Tags CheckoutSession
// @Accept json
// @Produce json
// @Param id path string true "CheckoutSession ID"
// @Success 204 {string} status "ok"
// @Router /v1/checkout/sessions/{id} [delete]
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

// RouteExpire func for expireing CheckoutSessions by given ID.
// @Description Expire CheckoutSession.
// @Summary expire CheckoutSession
// @Tags CheckoutSession
// @Accept json
// @Produce json
// @Param id path string true "CheckoutSession ID"
// @Success 200 {object} CheckoutSession
// @Router /v1/checkout/sessions/{id}/expire [post]
func (h *Handler) RouteExpire(context *fiber.Ctx) error {
	var m CheckoutSession

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "open" {
		m.Status = "expired"
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
