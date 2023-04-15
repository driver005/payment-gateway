package subscriptionSchedule

import (
	"strconv"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/subscription_schedules", h.RouteList)
	r.Post("/subscription_schedules", h.RouteCreate)
	r.Get("/subscription_schedules/:id", h.RouteGet)
	r.Post("/subscription_schedules/:id", h.RouteUpdate)
	r.Post("/subscription_schedules/:id/cancel", h.RouteCancel)
	r.Post("/subscription_schedules/:id/release", h.RouteRelease)
}

// RouteGet func gets SubscriptionSchedule by given ID or 404 error.
// @Description Get SubscriptionSchedule by given ID or 404 error.
// @Summary get SubscriptionSchedule by given ID or 404 error.
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionSchedule ID"
// @Success 200 {object} SubscriptionSchedule
// @Router /v1/subscription_schedules/{id} [get]
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

// RouteList func gets all existing SubscriptionSchedules.
// @Description Get all existing SubscriptionSchedules.
// @Summary get all existing SubscriptionSchedules
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Success 200 {array} SubscriptionSchedule
// @Router /v1/subscription_schedules [get]
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

// RouteCreate func for creates a new SubscriptionSchedule.
// @Description Create a new SubscriptionSchedule.
// @Summary create a new SubscriptionSchedule
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Param model body subscriptionSchedule.Bind.request true "Request Data"
// @Success 200 {object} SubscriptionSchedule
// @Router /v1/subscription_schedules [post]
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

// RouteUpdate func for updates SubscriptionSchedule by given ID.
// @Description Update SubscriptionSchedule.
// @Summary update SubscriptionSchedule
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Param model body subscriptionSchedule.Bind.request true "Request Data"
// @Param id path string true "SubscriptionSchedule ID"
// @Success 200 {object} SubscriptionSchedule
// @Router /v1/subscription_schedules/{id} [post]
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

// RouteDelete func for deletes SubscriptionSchedule by given ID.
// @Description Delete SubscriptionSchedule by given ID.
// @Summary delete SubscriptionSchedule by given ID
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionSchedule ID"
// @Success 204 {string} status "ok"
// @Router /v1/subscription_schedules/{id} [delete]
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

// RouteCancel func for canceling SubscriptionSchedules by given ID.
// @Description Cancel SubscriptionSchedule.
// @Summary cancel SubscriptionSchedule
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionSchedule ID"
// @Success 200 {object} SubscriptionSchedule
// @Router /v1/subscription_schedules/{id}/cancele [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m SubscriptionSchedule

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "not_started" || m.Status == "active" {
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

// RouteRelease func for releasing SubscriptionSchedules by given ID.
// @Description Release SubscriptionSchedule.
// @Summary release SubscriptionSchedule
// @Tags SubscriptionSchedule
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionSchedule ID"
// @Success 200 {object} SubscriptionSchedule
// @Router /v1/subscription_schedules/{id}/release [post]
func (h *Handler) RouteRelease(context *fiber.Ctx) error {
	var m SubscriptionSchedule

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	m.ReleasedAt = time.Now().UTC().Round(time.Second)
	if m.Status == "not_started" || m.Status == "active" {
		m.Status = "released"
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
