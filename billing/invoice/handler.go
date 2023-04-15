package invoice

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/invoices", h.RouteList)
	r.Post("/invoices", h.RouteCreate)
	r.Get("/invoices/:id", h.RouteGet)
	r.Post("/invoices/:id", h.RouteUpdate)
	r.Delete("/invoices/:id", h.RouteDelete)
	r.Post("/invoices/:id/finalize", h.RouteFinalize)
	r.Post("/invoices/:id/pay", h.RoutePay)
	r.Post("/invoices/:id/send", h.RouteSend)
	r.Post("/invoices/:id/void", h.RouteVoid)
	r.Post("/invoices/:id/mark_uncollectible", h.RouteMarkUncollectible)
	r.Post("/invoices/:id/upcoming", h.RouteUpcoming)
}

// RouteGet func gets Invoice by given ID or 404 error.
// @Description Get Invoice by given ID or 404 error.
// @Summary get Invoice by given ID or 404 error.
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id} [get]
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

// RouteList func gets all existing Invoices.
// @Description Get all existing Invoices.
// @Summary get all existing Invoices
// @Tags Invoice
// @Accept json
// @Produce json
// @Success 200 {array} Invoice
// @Router /v1/invoices [get]
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

// RouteCreate func for creates a new Invoice.
// @Description Create a new Invoice.
// @Summary create a new Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param model body invoice.Bind.request true "Request Data"
// @Success 200 {object} Invoice
// @Router /v1/invoices [post]
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

// RouteUpdate func for updates Invoice by given ID.
// @Description Update Invoice.
// @Summary update Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param model body invoice.Bind.request true "Request Data"
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id} [post]
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

// RouteDelete func for deletes Invoice by given ID.
// @Description Delete Invoice by given ID.
// @Summary delete Invoice by given ID
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 204 {string} status "ok"
// @Router /v1/invoices/{id} [delete]
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

// RouteCapture func for capturing Invoices by given ID.
// @Description Capture Invoice.
// @Summary capture Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/capture [post]
func (h *Handler) RouteCapture(context *fiber.Ctx) error {
	var m Invoice

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

// RouteCancel func for canceling Invoices by given ID.
// @Description Cancel Invoice.
// @Summary cancel Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/cancel [post]
func (h *Handler) RouteCancel(context *fiber.Ctx) error {
	var m Invoice

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

// RouteFinalize func for finalizing Invoices by given ID.
// @Description Finalize Invoice.
// @Summary finalize Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/finalize [post]
func (h *Handler) RouteFinalize(context *fiber.Ctx) error {
	var m Invoice

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "draft" {
		m.Status = "open"
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

// RoutePay func for paying Invoices by given ID.
// @Description Pay Invoice.
// @Summary pay Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/pay [post]
func (h *Handler) RoutePay(context *fiber.Ctx) error {
	var m Invoice

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "open" || m.Status == "draft" {
		m.Status = "paid"
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

// RouteSend func for sending Invoices by given ID.
// @Description Send Invoice.
// @Summary send Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/send [post]
func (h *Handler) RouteSend(context *fiber.Ctx) error {
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
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

// RouteVoid func for voiding Invoices by given ID.
// @Description Void Invoice.
// @Summary void Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/void [post]
func (h *Handler) RouteVoid(context *fiber.Ctx) error {
	var m Invoice

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "open" {
		m.Status = "void"
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

// RouteMarkUncollectible func for marking Invoices as uncollectible by given ID.
// @Description Mark Invoice as uncollectible.
// @Summary mark Invoice as uncollectible
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/mark_uncollectible [post]
func (h *Handler) RouteMarkUncollectible(context *fiber.Ctx) error {
	var m Invoice

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m.Id = Id
	if m.Status == "open" || m.Status == "draft" {
		m.Status = "uncollectible"
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

// RouteUpcoming func for upcoming Invoices by given ID.
// @Description Upcoming Invoice.
// @Summary upcoming Invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} Invoice
// @Router /v1/invoices/{id}/upcoming [post]
func (h *Handler) RouteUpcoming(context *fiber.Ctx) error {
	// Id, err := uuid.Parse(context.Params("id"))
	// if err != nil {
	// 	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 		"type":    "parse_error",
	// 	})
	// }

	// r, err := h.Retrive(context.Context(), Id)
	// if err != nil {
	// 	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 		"type":    "database_error",
	// 	})
	// }

	// return context.Status(fiber.StatusOK).JSON(&r)
	return context.Status(fiber.StatusNoContent).JSON(nil)
}
