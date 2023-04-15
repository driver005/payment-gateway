package coupon

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/coupons", h.RouteList)
	r.Post("/coupons", h.RouteCreate)
	r.Get("/coupons/:id", h.RouteGet)
	r.Post("/coupons/:id", h.RouteUpdate)
	r.Delete("/coupons/:id", h.RouteDelete)
}

// RouteGet func gets Coupon by given ID or 404 error.
// @Description Get Coupon by given ID or 404 error.
// @Summary get Coupon by given ID or 404 error.
// @Tags Coupon
// @Accept json
// @Produce json
// @Param id path string true "Coupon ID"
// @Success 200 {object} Coupon
// @Router /v1/coupons/{id} [get]
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

// RouteList func gets all existing Coupons.
// @Description Get all existing Coupons.
// @Summary get all existing Coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Success 200 {array} Coupon
// @Router /v1/coupons [get]
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

// RouteCreate func for creates a new Coupon.
// @Description Create a new Coupon.
// @Summary create a new Coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param model body coupon.Bind.request true "Request Data"
// @Success 200 {object} Coupon
// @Router /v1/coupons [post]
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

// RouteUpdate func for updates Coupon by given ID.
// @Description Update Coupon.
// @Summary update Coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param model body coupon.Bind.request true "Request Data"
// @Param id path string true "Coupon ID"
// @Success 200 {object} Coupon
// @Router /v1/coupons/{id} [post]
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

// RouteDelete func for deletes Coupon by given ID.
// @Description Delete Coupon by given ID.
// @Summary delete Coupon by given ID
// @Tags Coupon
// @Accept json
// @Produce json
// @Param id path string true "Coupon ID"
// @Success 204 {string} status "ok"
// @Router /v1/coupons/{id} [delete]
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
