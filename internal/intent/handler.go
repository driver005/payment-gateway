package intent

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/payment_intents", h.RouteList)
	r.Post("/payment_intents", h.RouteCreate)
	r.Get("/payment_intents/:id", h.RouteGet)
	r.Post("/payment_intents/:id", h.RouteUpdate)
}

// RouteGet func gets PaymentIntent by given ID or 404 error.
// @Description Get PaymentIntent by given ID or 404 error.
// @Summary get PaymentIntent by given ID or 404 error.
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param id path string true "PaymentIntent ID"
// @Success 200 {object} PaymentIntent
// @Router /v1/payment_intents/{id} [get]
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

// RouteList func gets all existing PaymentIntents.
// @Description Get all existing PaymentIntents.
// @Summary get all existing PaymentIntents
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Success 200 {array} PaymentIntent
// @Router /v1/payment_intents [get]
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

// RouteCreate func for creates a new PaymentIntent.
// @Description Create a new PaymentIntent.
// @Summary create a new PaymentIntent
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param model body intent.Bind.request true "Request Data"
// @Success 200 {object} PaymentIntent
// @Router /v1/payment_intents [post]
func (h *Handler) RouteCreate(context *fiber.Ctx) error {
	m, err := h.Bind(context, true)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteUpdate func for updates PaymentIntent by given ID.
// @Description Update PaymentIntent.
// @Summary update PaymentIntent
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param model body intent.Bind.request true "Request Data"
// @Param id path string true "PaymentIntent ID"
// @Success 200 {object} PaymentIntent
// @Router /v1/payment_intents/{id} [post]
func (h *Handler) RouteUpdate(context *fiber.Ctx) error {
	m, err := h.Bind(context, false)
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

	r, err = h.Pay(context.Context(), *m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteDelete func for deletes PaymentIntent by given ID.
// @Description Delete PaymentIntent by given ID.
// @Summary delete PaymentIntent by given ID
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param id path string true "PaymentIntent ID"
// @Success 204 {string} status "ok"
// @Router /v1/payment_intents/{id} [delete]
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

//TODO: ADD all

// RouteConfirm func for confirming PaymentIntents by given ID.
// @Description Confirm PaymentIntent.
// @Summary confirm PaymentIntent
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param id path string true "PaymentIntent ID"
// @Success 200 {object} PaymentIntent
// @Router /v1/payment_intents/{id}/confirm [post]
func (h *Handler) RouteConfirm(context *fiber.Ctx) error {
	m, err := h.Bind(context, false)
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
	if m.Status == "requires_action" {
		m.Status = "succeeded"
	}

	r, err := h.Update(context.Context(), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err = h.Pay(context.Context(), *m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteCapture func for capturing PaymentIntents by given ID.
// @Description Capture PaymentIntent.
// @Summary capture PaymentIntent
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param id path string true "PaymentIntent ID"
// @Success 200 {object} PaymentIntent
// @Router /v1/payment_intents/{id}/capture [post]
func (h *Handler) RouteCapture(context *fiber.Ctx) error {
	var m PaymentIntent

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "requires_capture" {
		m.Status = "succeeded"
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

// RouteCancel func for canceling PaymentIntents by given ID.
// @Description Cancel PaymentIntent.
// @Summary cancel PaymentIntent
// @Tags PaymentIntent
// @Accept json
// @Produce json
// @Param id path string true "PaymentIntent ID"
// @Success 200 {object} PaymentIntent
// @Router /v1/payment_intents/{id}/cancel [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m PaymentIntent

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status != "succeeded" {
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
