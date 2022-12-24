package admin

import (
	"strconv"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (a *Admin) GetProduct(context *fiber.Ctx) error {
	f := models.Product{}

	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}
	f.Id = Id

	m, err := a.r.ClientManager().GetProduct(context.Context(), types.FilterableProductProps{}, f)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) ListProduct(context *fiber.Ctx) error {
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

	m, n, err := a.r.ClientManager().GetProducts(context.Context(), models.Filter{
		Size:   pageSize,
		Offset: offset,
	}, types.FilterableProductProps{}, models.Product{})
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	helper.Header(context, context.Request().URI(), *n, pageSize, offset)

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) CreateProduct(context *fiber.Ctx) error {
	var m models.Product

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := a.r.ClientManager().CreateProduct(context.Context(), &m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateProduct(context *fiber.Ctx) error {
	var m models.Product

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
			"type":    "database_error",
		})
	}

	m.Id = Id

	r, err := a.r.ClientManager().UpdateProduct(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

func (a *Admin) DeleteProduct(context *fiber.Ctx) error {
	Id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteProduct(context.Context(), Id)
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}

type PostProductsProductMetadataJSONBody struct {
	// The metadata key
	Key string `json:"key"`

	// The metadata value
	Value string `json:"value"`
}

func (a *Admin) CreateProductMetadata(context *fiber.Ctx) error {
	var m models.Product
	var o models.Product
	var b PostProductsProductMetadataJSONBody

	if err := context.BodyParser(&b); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	ID, err := uuid.Parse(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	if err = a.r.Manager(context.Context()).Where("id = ?", ID).First(&o).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Id = o.Id
	m.Metadata = models.JSONB{
		b.Key: b.Value,
	}
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := a.r.Manager(context.Context()).Model(&o).Updates(m).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}
