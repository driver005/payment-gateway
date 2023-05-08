package usageRecord

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	//TODO: Add UsageRecord
	// r.Get("/subscription_items/:id/usage_record_summaries", h.RouteList)
	// r.Post("subscription_items/:id/usage_records", h.RouteCreate)
}

// RouteGet func gets UsageRecord by given ID or 404 error.
// @Description Get UsageRecord by given ID or 404 error.
// @Summary get UsageRecord by given ID or 404 error.
// @Tags UsageRecord
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionItem ID"
// @Success 200 {object} UsageRecord
// @Router /v1/subscription_items/{id}/usage_records [get]
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

// RouteList func gets all existing UsageRecords.
// @Description Get all existing UsageRecords.
// @Summary get all existing UsageRecords
// @Tags UsageRecord
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionItem ID"
// @Success 200 {array} UsageRecord
// @Router /v1/subscription_items/{id}/usage_record_summaries [get]
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

// RouteCreate func for creates a new UsageRecord.
// @Description Create a new UsageRecord.
// @Summary create a new UsageRecord
// @Tags UsageRecord
// @Accept json
// @Produce json
// @Param model body UsageRecord true "Request Data"
// @Success 200 {object} UsageRecord
// @Router /v1/usage_record [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	var m UsageRecord

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := h.Create(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteUpdate func for updates UsageRecord by given ID.
// @Description Update UsageRecord.
// @Summary update UsageRecord
// @Tags UsageRecord
// @Accept json
// @Produce json
// @Param model body UsageRecord true "Request Data"
// @Param id path string true "SubscriptionItem ID"
// @Success 200 {object} UsageRecord
// @Router /v1/usage_record/{id} [post]
func (h *Handler) RouteUpdate(context *fiber.Ctx) error {
	var m UsageRecord

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

	r, err := h.Update(context, &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDelete func for deletes UsageRecord by given ID.
// @Description Delete UsageRecord by given ID.
// @Summary delete UsageRecord by given ID
// @Tags UsageRecord
// @Accept json
// @Produce json
// @Param id path string true "UsageRecord ID"
// @Success 204 {string} status "ok"
// @Router /v1/usage_record/{id} [delete]
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
