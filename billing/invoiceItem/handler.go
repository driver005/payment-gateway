package invoiceItem

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/invoiceitems", h.RouteList)
	r.Post("/invoiceitems", h.RouteCreate)
	r.Get("/invoiceitems/:id", h.RouteGet)
	r.Post("/invoiceitems/:id", h.RouteUpdate)
	r.Delete("/invoiceitems/:id", h.RouteDelete)
}

// RouteGet func gets Invoiceitem by given ID or 404 error.
// @Description Get Invoiceitem by given ID or 404 error.
// @Summary get Invoiceitem by given ID or 404 error.
// @Tags Invoiceitem
// @Accept json
// @Produce json
// @Param id path string true "Invoiceitem ID"
// @Success 200 {object} Invoiceitem
// @Router /v1/invoiceitems/{id} [get]
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

// RouteList func gets all existing Invoiceitems.
// @Description Get all existing Invoiceitems.
// @Summary get all existing Invoiceitems
// @Tags Invoiceitem
// @Accept json
// @Produce json
// @Success 200 {array} Invoiceitem
// @Router /v1/invoiceitems [get]
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

// RouteCreate func for creates a new Invoiceitem.
// @Description Create a new Invoiceitem.
// @Summary create a new Invoiceitem
// @Tags Invoiceitem
// @Accept json
// @Produce json
// @Param model body invoiceItem.Bind.request true "Request Data"
// @Success 200 {object} Invoiceitem
// @Router /v1/invoiceitems [post]
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

// RouteUpdate func for updates Invoiceitem by given ID.
// @Description Update Invoiceitem.
// @Summary update Invoiceitem
// @Tags Invoiceitem
// @Accept json
// @Produce json
// @Param model body invoiceItem.Bind.request true "Request Data"
// @Param id path string true "Invoiceitem ID"
// @Success 200 {object} Invoiceitem
// @Router /v1/invoiceitems/{id} [post]
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

// RouteDelete func for deletes Invoiceitem by given ID.
// @Description Delete Invoiceitem by given ID.
// @Summary delete Invoiceitem by given ID
// @Tags Invoiceitem
// @Accept json
// @Produce json
// @Param id path string true "Invoiceitem ID"
// @Success 204 {string} status "ok"
// @Router /v1/invoiceitems/{id} [delete]
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
