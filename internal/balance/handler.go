package balance

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Get("/balance", h.RouteGetBalance)
	r.Get("/balance_transactions", h.RouteListBalanceTransaction)
	r.Get("/balance_transactions/:id", h.RouteGetBalanceTransaction)
}

// RouteGet func gets Balance by given ID or 404 error.
// @Description Get Balance by given ID or 404 error.
// @Summary get Balance by given ID or 404 error.
// @Tags Balance
// @Accept json
// @Produce json
// @Param id path string true "Balance ID"
// @Success 200 {object} Balance
// @Router /v1/balance/{id} [get]
func (h *Handler) RouteGetBalance(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m, err := h.RetriveBalance(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteGet func gets BalanceTransaction by given ID or 404 error.
// @Description Get BalanceTransaction by given ID or 404 error.
// @Summary get BalanceTransaction by given ID or 404 error.
// @Tags BalanceTransaction
// @Accept json
// @Produce json
// @Param id path string true "BalanceTransaction ID"
// @Success 200 {object} BalanceTransaction
// @Router /v1/balance_transactions/{id} [get]
func (h *Handler) RouteGetBalanceTransaction(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	m, err := h.RetriveBalanceTransaction(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

// RouteList func gets all existing BalanceTransaction.
// @Description Get all existing BalanceTransaction.
// @Summary get all existing BalanceTransaction
// @Tags BalanceTransaction
// @Accept json
// @Produce json
// @Success 200 {array} BalanceTransaction
// @Router /v1/balance_transactions [get]
func (h *Handler) RouteListBalanceTransaction(context *fiber.Ctx) error {
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

	m, n, err := h.ListBalanceTransaction(context.Context(), offset, pageSize)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}
