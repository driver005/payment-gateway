package dispute

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/disputes", h.RouteList)
	r.Post("/disputes", h.RouteCreate)
	r.Get("/disputes/:id", h.RouteGet)
	r.Post("/disputes/:id", h.RouteUpdate)
	r.Post("/disputes/:id/close", h.RouteClose)
}

// RouteGet func gets Dispute by given ID or 404 error.
// @Description Get Dispute by given ID or 404 error.
// @Summary get Dispute by given ID or 404 error.
// @Tags Dispute
// @Accept json
// @Produce json
// @Param id path string true "Dispute ID"
// @Success 200 {object} Dispute
// @Router /v1/disputes/{id} [get]
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

// RouteList func gets all existing Disputes.
// @Description Get all existing Disputes.
// @Summary get all existing Disputes
// @Tags Dispute
// @Accept json
// @Produce json
// @Success 200 {array} Dispute
// @Router /v1/disputes [get]
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

// RouteCreate func for creates a new Dispute.
// @Description Create a new Dispute.
// @Summary create a new Dispute
// @Tags Dispute
// @Accept json
// @Produce json
// @Param model body dispute.Bind.request true "Request Data"
// @Success 200 {object} Dispute
// @Router /v1/disputes [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	m, err := h.Bind(context)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	m.Status = "needs_response"

	r, err := h.Create(context, m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteUpdate func for updates Dispute by given ID.
// @Description Update Dispute.
// @Summary update Dispute
// @Tags Dispute
// @Accept json
// @Produce json
// @Param model body dispute.Bind.request true "Request Data"
// @Param id path string true "Dispute ID"
// @Success 200 {object} Dispute
// @Router /v1/disputes/{id} [post]
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

// RouteDelete func for deletes Dispute by given ID.
// @Description Delete Dispute by given ID.
// @Summary delete Dispute by given ID
// @Tags Dispute
// @Accept json
// @Produce json
// @Param id path string true "Dispute ID"
// @Success 204 {string} status "ok"
// @Router /v1/disputes/{id} [delete]
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

// RouteClose func for closing Disputes by given ID.
// @Description Close Dispute.
// @Summary close Dispute
// @Tags Dispute
// @Accept json
// @Produce json
// @Param id path string true "Dispute ID"
// @Success 200 {object} Dispute
// @Router /v1/disputes/{id}/close [post]
func (h *Handler) RouteClose(context *fiber.Ctx) error {
	var m Dispute

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "needs_response" {
		m.Status = "lost"
	}

	r, err := h.Update(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}
