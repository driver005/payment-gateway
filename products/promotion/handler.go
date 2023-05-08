package promotion

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/promotion_codes", h.RouteList)
	r.Post("/promotion_codes", h.RouteCreate)
	r.Get("/promotion_codes/:id", h.RouteGet)
	r.Post("/promotion_codes/:id", h.RouteUpdate)
}

// RouteGet func gets PromotionCode by given ID or 404 error.
// @Description Get PromotionCode by given ID or 404 error.
// @Summary get PromotionCode by given ID or 404 error.
// @Tags PromotionCode
// @Accept json
// @Produce json
// @Param id path string true "PromotionCode ID"
// @Success 200 {object} PromotionCode
// @Router /v1/promotion_codes/{id} [get]
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

// RouteList func gets all existing PromotionCodes.
// @Description Get all existing PromotionCodes.
// @Summary get all existing PromotionCodes
// @Tags PromotionCode
// @Accept json
// @Produce json
// @Success 200 {array} PromotionCode
// @Router /v1/promotion_codes [get]
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

// RouteCreate func for creates a new PromotionCode.
// @Description Create a new PromotionCode.
// @Summary create a new PromotionCode
// @Tags PromotionCode
// @Accept json
// @Produce json
// @Param model body promotion.Bind.request true "Request Data"
// @Success 200 {object} PromotionCode
// @Router /v1/promotion_codes [post]
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

// RouteUpdate func for updates PromotionCode by given ID.
// @Description Update PromotionCode.
// @Summary update PromotionCode
// @Tags PromotionCode
// @Accept json
// @Produce json
// @Param model body promotion.Bind.request true "Request Data"
// @Param id path string true "PromotionCode ID"
// @Success 200 {object} PromotionCode
// @Router /v1/promotion_codes/{id} [post]
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

// RouteDelete func for deletes PromotionCode by given ID.
// @Description Delete PromotionCode by given ID.
// @Summary delete PromotionCode by given ID
// @Tags PromotionCode
// @Accept json
// @Produce json
// @Param id path string true "PromotionCode ID"
// @Success 204 {string} status "ok"
// @Router /v1/promotion_codes/{id} [delete]
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
