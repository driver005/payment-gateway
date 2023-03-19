package source

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/sources", h.RouteList)
	r.Post("/sources", h.RouteCreate)
	r.Get("/sources/:id", h.RouteGet)
	r.Post("/sources/:id", h.RouteUpdate)
	r.Post("/customers/:id/sources", h.RouteAttach)
	r.Delete("/customers/:id/sources/:id", h.RouteDettach)
}

// RouteGet func gets Source by given ID or 404 error.
// @Description Get Source by given ID or 404 error.
// @Summary get Source by given ID or 404 error.
// @Tags Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} Source
// @Router /v1/sources/{id} [get]
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

// RouteList func gets all existing Sources.
// @Description Get all existing Sources.
// @Summary get all existing Sources
// @Tags Source
// @Accept json
// @Produce json
// @Success 200 {array} Source
// @Router /v1/sources [get]
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

// RouteCreate func for creates a new Source.
// @Description Create a new Source.
// @Summary create a new Source
// @Tags Source
// @Accept json
// @Produce json
// @Success 200 {object} Source
// @Router /v1/sources [post]
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

// RouteUpdate func for updates Source by given ID.
// @Description Update Source.
// @Summary update Source
// @Tags Source
// @Accept json
// @Produce json
// @Param id body string true "Source ID"
// @Success 200 {object} Source
// @Router /v1/sources/{id} [post]
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

// RouteDelete func for deletes Source by given ID.
// @Description Delete Source by given ID.
// @Summary delete Source by given ID
// @Tags Source
// @Accept json
// @Produce json
// @Param id body string true "Source ID"
// @Success 204 {string} status "ok"
// @Router /v1/sources/{id} [delete]
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

// RouteAttach func for attaching Source to Customer by given ID.
// @Description Attaching Source to Customer by given ID.
// @Summary attaching Source to Customer by given ID
// @Tags Source
// @Accept json
// @Produce json
// @Param id body string true "Source ID"
// @Success 204 {string} status "ok"
// @Router /v1/customers/{id}/sources [post]
func (h *Handler) RouteAttach(context *fiber.Ctx) error {
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

// RouteDettach func for dettaching Source from Customer by given ID.
// @Description dettaching Source from Customer by given ID.
// @Summary dettaching Source from Customer by given ID
// @Tags Source
// @Accept json
// @Produce json
// @Param id body string true "Source ID"
// @Success 204 {string} status "ok"
// @Router /v1/customers/{id}/sources/{id} [delete]
func (h *Handler) RouteDettach(context *fiber.Ctx) error {
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
