package subscriptionItem

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/subscription_items", h.RouteList)
	r.Post("/subscription_items", h.RouteCreate)
	r.Get("/subscription_items/:id", h.RouteGet)
	r.Post("/subscription_items/:id", h.RouteUpdate)
	r.Delete("/subscription_items/:id", h.RouteDelete)
}

// RouteGet func gets SubscriptionItem by given ID or 404 error.
// @Description Get SubscriptionItem by given ID or 404 error.
// @Summary get SubscriptionItem by given ID or 404 error.
// @Tags SubscriptionItem
// @Accept json
// @Produce json
// @Param id path string true "SubscriptionItem ID"
// @Success 200 {object} SubscriptionItem
// @Router /v1/subscription_items/{id} [get]
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

// RouteList func gets all existing SubscriptionItems.
// @Description Get all existing SubscriptionItems.
// @Summary get all existing SubscriptionItems
// @Tags SubscriptionItem
// @Accept json
// @Produce json
// @Success 200 {array} SubscriptionItem
// @Router /v1/subscription_items [get]
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

// RouteCreate func for creates a new SubscriptionItem.
// @Description Create a new SubscriptionItem.
// @Summary create a new SubscriptionItem
// @Tags SubscriptionItem
// @Accept json
// @Produce json
// @Success 200 {object} SubscriptionItem
// @Router /v1/subscription_items [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	var m SubscriptionItem

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

// RouteUpdate func for updates SubscriptionItem by given ID.
// @Description Update SubscriptionItem.
// @Summary update SubscriptionItem
// @Tags SubscriptionItem
// @Accept json
// @Produce json
// @Param id body string true "SubscriptionItem ID"
// @Success 200 {object} SubscriptionItem
// @Router /v1/subscription_items/{id} [post]
func (h *Handler) RouteUpdate(context *fiber.Ctx) error {
	var m SubscriptionItem

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

// RouteDelete func for deletes SubscriptionItem by given ID.
// @Description Delete SubscriptionItem by given ID.
// @Summary delete SubscriptionItem by given ID
// @Tags SubscriptionItem
// @Accept json
// @Produce json
// @Param id body string true "SubscriptionItem ID"
// @Success 204 {string} status "ok"
// @Router /v1/subscription_items/{id} [delete]
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
