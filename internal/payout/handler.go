package payout

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/payouts", h.RouteList)
	r.Post("/payouts", h.RouteCreate)
	r.Get("/payouts/:id", h.RouteGet)
	r.Post("/payouts/:id", h.RouteUpdate)
	r.Post("/payouts/:id/cancele", h.RouteCancel)
	r.Post("/payouts/:id/reverse", h.RouteReverse)
}

// RouteGet func gets Payout by given ID or 404 error.
// @Description Get Payout by given ID or 404 error.
// @Summary get Payout by given ID or 404 error.
// @Tags Payout
// @Accept json
// @Produce json
// @Param id path string true "Payout ID"
// @Success 200 {object} Payout
// @Router /v1/payouts/{id} [get]
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

// RouteList func gets all existing Payouts.
// @Description Get all existing Payouts.
// @Summary get all existing Payouts
// @Tags Payout
// @Accept json
// @Produce json
// @Success 200 {array} Payout
// @Router /v1/payouts [get]
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

// RouteCreate func for creates a new Payout.
// @Description Create a new Payout.
// @Summary create a new Payout
// @Tags Payout
// @Accept json
// @Produce json
// @Param model body payout.Bind.request true "Request Data"
// @Success 200 {object} Payout
// @Router /v1/payouts [post]
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

// RouteUpdate func for updates Payout by given ID.
// @Description Update Payout.
// @Summary update Payout
// @Tags Payout
// @Accept json
// @Produce json
// @Param model body payout.Bind.request true "Request Data"
// @Param id path string true "Payout ID"
// @Success 200 {object} Payout
// @Router /v1/payouts/{id} [post]
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

// RouteDelete func for deletes Payout by given ID.
// @Description Delete Payout by given ID.
// @Summary delete Payout by given ID
// @Tags Payout
// @Accept json
// @Produce json
// @Param id path string true "Payout ID"
// @Success 204 {string} status "ok"
// @Router /v1/payouts/{id} [delete]
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

// RouteCancel func for canceling Payouts by given ID.
// @Description Cancel Payout.
// @Summary cancel Payout
// @Tags Payout
// @Accept json
// @Produce json
// @Param id path string true "Payout ID"
// @Success 200 {object} Payout
// @Router /v1/payouts/{id}/cancele [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m Payout

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	m.Status = "canceled"

	r, err := h.Update(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteReverse func for reversing Payouts by given ID.
// @Description Reverse Payout.
// @Summary reverse Payout
// @Tags Payout
// @Accept json
// @Produce json
// @Param id path string true "Payout ID"
// @Success 200 {object} Payout
// @Router /v1/payouts/{id}/reverse [post]
func (h *Handler) RouteReverse(context *fiber.Ctx) error {
	var m Payout

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	m.Status = "succeeded"

	r, err := h.Update(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}
