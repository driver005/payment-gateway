package refund

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/refunds", h.RouteList)
	r.Post("/refunds", h.RouteCreate)
	r.Get("/refunds/:id", h.RouteGet)
	r.Post("/refunds/:id", h.RouteUpdate)
	r.Post("/refunds/:id/cancele", h.RouteCancel)
}

// RouteGet func gets Refund by given ID or 404 error.
// @Description Get Refund by given ID or 404 error.
// @Summary get Refund by given ID or 404 error.
// @Tags Refund
// @Accept json
// @Produce json
// @Param id path string true "Refund ID"
// @Success 200 {object} Refund
// @Router /v1/refunds/{id} [get]
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

// RouteList func gets all existing Refunds.
// @Description Get all existing Refunds.
// @Summary get all existing Refunds
// @Tags Refund
// @Accept json
// @Produce json
// @Success 200 {array} Refund
// @Router /v1/refunds [get]
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

// RouteCreate func for creates a new Refund.
// @Description Create a new Refund.
// @Summary create a new Refund
// @Tags Refund
// @Accept json
// @Produce json
// @Param model body refund.Bind.request true "Request Data"
// @Success 200 {object} Refund
// @Router /v1/refunds [post]
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

// RouteUpdate func for updates Refund by given ID.
// @Description Update Refund.
// @Summary update Refund
// @Tags Refund
// @Accept json
// @Produce json
// @Param model body refund.Bind.request true "Request Data"
// @Param id path string true "Refund ID"
// @Success 200 {object} Refund
// @Router /v1/refunds/{id} [post]
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

// RouteDelete func for deletes Refund by given ID.
// @Description Delete Refund by given ID.
// @Summary delete Refund by given ID
// @Tags Refund
// @Accept json
// @Produce json
// @Param id path string true "Refund ID"
// @Success 204 {string} status "ok"
// @Router /v1/refunds/{id} [delete]
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

// RouteCancel func for canceling Refunds by given ID.
// @Description Cancel Refund.
// @Summary cancel Refund
// @Tags Refund
// @Accept json
// @Produce json
// @Param id path string true "Refund ID"
// @Success 200 {object} Refund
// @Router /v1/refunds/{id}/cancele [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m Refund

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "requires_action" {
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
