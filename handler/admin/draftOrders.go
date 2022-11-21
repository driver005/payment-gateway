package admin

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func (a *Admin) GetDraftOrder(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDraftOrder(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (p *Admin) ListDraftOrder(context *fiber.Ctx) error {
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

	m, n, err := p.r.ClientManager().GetDraftOrders(context.Context(), models.Filter{
		Size:   pageSize,
		Offset: offset,
	})
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) CreateDraftOrder(context *fiber.Ctx) error {
	var m models.DraftOrder

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := a.r.ClientManager().CreateDraftOrder(context.Context(), &m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateDraftOrder(context *fiber.Ctx) error {
	var m models.DraftOrder

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Id = Id

	r, err := a.r.ClientManager().UpdateDraftOrder(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

func (a *Admin) DeleteDraftOrder(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteDraftOrder(context.Context(), Id)
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}

type DraftOrderLineItemsRequest struct {

	// The ID of the Product Variant to generate the Line Item from.
	VariantId uuid.NullUUID `json:"variant_id,omitempty"`

	// The potential custom price of the item.
	UnitPrice int `json:"unit_price,omitempty"`

	// The potential custom title of the item.
	Title string `json:"title,omitempty" default:"Custom item"`

	// The quantity of the Line Item.
	Quantity int32 `json:"quantity"`

	// The optional key-value map with additional details about the Line Item.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (a *Admin) CreateDraftOrderLineItem(context *fiber.Ctx) error {
	var b DraftOrderLineItemsRequest
	var l models.LineItem

	if err := context.BodyParser(&b); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDraftOrder(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	l.Quantity = b.Quantity
	l.VariantId = b.VariantId
	l.UnitPrice = b.UnitPrice
	l.Title = b.Title
	l.Metadata = b.Metadata

	_, err = a.r.ClientManager().CreateLineItem(context.Context(), &l)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Status = models.DraftOrderStatusCompleted

	r, err := a.r.ClientManager().CreateDraftOrder(context.Context(), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateDraftOrderLineItem(context *fiber.Ctx) error {
	var b DraftOrderLineItemsRequest
	if err := context.BodyParser(&b); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	LineId, err := uuid.FromString(context.Params("line_id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDraftOrder(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	l, err := a.r.ClientManager().GetLineItem(context.Context(), LineId)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	l.Quantity = b.Quantity
	l.VariantId = b.VariantId
	l.UnitPrice = b.UnitPrice
	l.Title = b.Title
	l.Metadata = b.Metadata

	_, err = a.r.ClientManager().CreateLineItem(context.Context(), l)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Status = models.DraftOrderStatusCompleted

	r, err := a.r.ClientManager().CreateDraftOrder(context.Context(), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateDraftOrderPayment(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDraftOrder(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Status = models.DraftOrderStatusCompleted

	r, err := a.r.ClientManager().CreateDraftOrder(context.Context(), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

