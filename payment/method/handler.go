package method

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/payment_methods", h.RouteList)
	r.Post("/payment_methods", h.RouteCreate)
	r.Get("/payment_methods/:id", h.RouteGet)
	r.Post("/payment_methods/:id", h.RouteUpdate)
	r.Post("/payment_methods/:id", h.RouteUpdate)
	r.Post("/payment_methods/:id/attach", h.RouteAttach)
	r.Post("/payment_methods/:id/dettach", h.RouteDettach)
}

// RouteGet func gets PaymentMethod by given ID or 404 error.
// @Description Get PaymentMethod by given ID or 404 error.
// @Summary get PaymentMethod by given ID or 404 error.
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod ID"
// @Success 200 {object} PaymentMethod
// @Router /v1/payment_methods/{id} [get]
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

// RouteList func gets all existing PaymentMethods.
// @Description Get all existing PaymentMethods.
// @Summary get all existing PaymentMethods
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Success 200 {array} PaymentMethod
// @Router /v1/payment_methods [get]
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

// RouteCreate func for creates a new PaymentMethod.
// @Description Create a new PaymentMethod.
// @Summary create a new PaymentMethod
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param model body method.Bind.request true "Request Data"
// @Success 200 {object} PaymentMethod
// @Router /v1/payment_methods [post]
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

// RouteUpdate func for updates PaymentMethod by given ID.
// @Description Update PaymentMethod.
// @Summary update PaymentMethod
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param model body method.Bind.request true "Request Data"
// @Param id path string true "PaymentMethod ID"
// @Success 200 {object} PaymentMethod
// @Router /v1/payment_methods/{id} [post]
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

// RouteDelete func for deletes PaymentMethod by given ID.
// @Description Delete PaymentMethod by given ID.
// @Summary delete PaymentMethod by given ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod ID"
// @Success 204 {string} status "ok"
// @Router /v1/payment_methods/{id} [delete]
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

// RouteAttach func for attaching PaymentMethod to Customer by given ID.
// @Description Attaching PaymentMethod to Customer by given ID.
// @Summary attaching PaymentMethod to Customer by given ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod ID"
// @Success 204 {string} status "ok"
// @Router /v1/payment_methods/{id}/attach [post]
func (h *Handler) RouteAttach(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	r, err := h.Retrive(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDettach func for dettaching PaymentMethod from Customer by given ID.
// @Description dettaching PaymentMethod from Customer by given ID.
// @Summary dettaching PaymentMethod from Customer by given ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod ID"
// @Success 204 {string} status "ok"
// @Router /v1/payment_methods/{id}/dettach [post]
func (h *Handler) RouteDettach(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	r, err := h.Retrive(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}
