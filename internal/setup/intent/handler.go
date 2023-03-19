package intent

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/setup_intents", h.RouteList)
	r.Post("/setup_intents", h.RouteCreate)
	r.Get("/setup_intents/:id", h.RouteGet)
	r.Post("/setup_intents/:id", h.RouteUpdate)
	r.Post("/setup_intents/:id/cancele", h.RouteCancel)
	r.Post("/setup_intents/:id/confirm", h.RouteConfirm)
}

// RouteGet func gets SetupIntent by given ID or 404 error.
// @Description Get SetupIntent by given ID or 404 error.
// @Summary get SetupIntent by given ID or 404 error.
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Param id path string true "SetupIntent ID"
// @Success 200 {object} SetupIntent
// @Router /v1/setup_intents/{id} [get]
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

// RouteList func gets all existing SetupIntents.
// @Description Get all existing SetupIntents.
// @Summary get all existing SetupIntents
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Success 200 {array} SetupIntent
// @Router /v1/setup_intents [get]
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

// RouteCreate func for creates a new SetupIntent.
// @Description Create a new SetupIntent.
// @Summary create a new SetupIntent
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Success 200 {object} SetupIntent
// @Router /v1/setup_intents [post]
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

// RouteUpdate func for updates SetupIntent by given ID.
// @Description Update SetupIntent.
// @Summary update SetupIntent
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Param id body string true "SetupIntent ID"
// @Success 200 {object} SetupIntent
// @Router /v1/setup_intents/{id} [post]
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

// RouteDelete func for deletes SetupIntent by given ID.
// @Description Delete SetupIntent by given ID.
// @Summary delete SetupIntent by given ID
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Param id body string true "SetupIntent ID"
// @Success 204 {string} status "ok"
// @Router /v1/setup_intents/{id} [delete]
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

// RouteCancel func for canceling SetupIntents by given ID.
// @Description Cancel SetupIntent.
// @Summary cancel SetupIntent
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Param id body string true "SetupIntent ID"
// @Success 200 {object} SetupIntent
// @Router /v1/setup_intents/{id}/cancele [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m SetupIntent

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

// RouteConfirm func for confirming SetupIntents by given ID.
// @Description Confirm SetupIntent.
// @Summary confirm SetupIntent
// @Tags SetupIntent
// @Accept json
// @Produce json
// @Param id body string true "SetupIntent ID"
// @Success 200 {object} SetupIntent
// @Router /v1/setup_intents/{id}/confirm [post]
func (h *Handler) RouteConfirm(context *fiber.Ctx) error {
	var m SetupIntent

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "requires_action" {
		m.Status = "next_action"
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
