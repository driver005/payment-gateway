package attempt

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/setup_attempts", h.RouteList)
	// r.Post("/setup_attempts", h.RouteCreate)
	// r.Get("/setup_attempts/:id", h.RouteGet)
	// r.Post("/setup_attempts/:id", h.RouteUpdate)
}

// RouteGet func gets SetupAttempt by given ID or 404 error.
// @Description Get SetupAttempt by given ID or 404 error.
// @Summary get SetupAttempt by given ID or 404 error.
// @Tags SetupAttempt
// @Accept json
// @Produce json
// @Param id path string true "SetupAttempt ID"
// @Success 200 {object} SetupAttempt
// @Router /v1/setup_attempts/{id} [get]
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

// RouteList func gets all existing SetupAttempts.
// @Description Get all existing SetupAttempts.
// @Summary get all existing SetupAttempts
// @Tags SetupAttempt
// @Accept json
// @Produce json
// @Success 200 {array} SetupAttempt
// @Router /v1/setup_attempts [get]
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

// RouteCreate func for creates a new SetupAttempt.
// @Description Create a new SetupAttempt.
// @Summary create a new SetupAttempt
// @Tags SetupAttempt
// @Accept json
// @Produce json
// @Param model body attempt.Bind.request true "Request Data"
// @Success 200 {object} SetupAttempt
// @Router /v1/setup_attempts [post]
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

// RouteUpdate func for updates SetupAttempt by given ID.
// @Description Update SetupAttempt.
// @Summary update SetupAttempt
// @Tags SetupAttempt
// @Accept json
// @Produce json
// @Param model body attempt.Bind.request true "Request Data"
// @Param id path string true "SetupAttempt ID"
// @Success 200 {object} SetupAttempt
// @Router /v1/setup_attempts/{id} [post]
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

// RouteDelete func for deletes SetupAttempt by given ID.
// @Description Delete SetupAttempt by given ID.
// @Summary delete SetupAttempt by given ID
// @Tags SetupAttempt
// @Accept json
// @Produce json
// @Param id path string true "SetupAttempt ID"
// @Success 204 {string} status "ok"
// @Router /v1/setup_attempts/{id} [delete]
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
