package admin

import (
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func (a *Admin) removeRegion(slice []models.Region, s int) []models.Region {
	return append(slice[:s], slice[s+1:]...)
}

func (a *Admin) GetDiscount(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDiscount(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) GetDiscountByCode(context *fiber.Ctx) error {
	m, err := a.r.ClientManager().GetDiscountByCode(context.Context(), context.Params("code"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}

func (a *Admin) ListDiscount(context *fiber.Ctx) error {
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

	m, n, err := a.r.ClientManager().GetDiscounts(context.Context(), models.Filter{
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

func (a *Admin) CreateDiscount(context *fiber.Ctx) error {
	var m models.Discount

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r, err := a.r.ClientManager().CreateDiscount(context.Context(), &m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) UpdateDiscount(context *fiber.Ctx) error {
	var m models.Discount

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

	r, err := a.r.ClientManager().UpdateDiscount(context.Context(), &m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(r)
}

func (a *Admin) DeleteDiscount(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteDiscount(context.Context(), Id)
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}

type DynamicCodesRequest struct {

	// The unique code that will be used to redeem the Discount.
	Code string `json:"code"`

	// amount of times the discount can be applied.
	UsageLimit int32 `json:"usage_limit,omitempty"`

	// An optional set of key-value paris to hold additional information.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (a *Admin) UpdateDynamicCode(context *fiber.Ctx) error {
	var b DynamicCodesRequest

	if err := context.BodyParser(&b); err != nil {
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

	m, err := a.r.ClientManager().GetDiscount(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Code = b.Code
	m.UsageLimit = b.UsageLimit
	m.Metadata = b.Metadata
	m.ParentDiscountId = uuid.NullUUID{UUID: Id, Valid: true}

	r, err := a.r.ClientManager().CreateDiscount(context.Context(), m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) DeleteDynamicCode(context *fiber.Ctx) error {
	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	r := a.r.ClientManager().DeleteDynamicCode(context.Context(), Id, context.Params("code"))
	if r != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": helper.WithStack(r),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusNoContent).JSON(r)
}

func (a *Admin) UpdateRegion(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	RegionId, err := uuid.FromString(context.Params("region_id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDiscount(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	l, err := a.r.ClientManager().GetRegion(context.Context(), RegionId)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Id = Id
	m.Regions = append(m.Regions, *l)

	r, err := a.r.ClientManager().UpdateDiscount(context.Context(), m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

func (a *Admin) DeleteRegion(context *fiber.Ctx) error {

	Id, err := uuid.FromString(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	RegionId, err := uuid.FromString(context.Params("region_id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDiscount(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	l, err := a.r.ClientManager().GetRegion(context.Context(), RegionId)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m.Id = Id
	for i, pd := range m.Regions {
		if pd.Id == l.Id {
			m.Regions = a.removeRegion(m.Regions, i)
		}
	}

	r, err := a.r.ClientManager().UpdateDiscount(context.Context(), m)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&r)
}

type DiscountConditionsRequest struct {

	// list of product IDs if the condition is applied on products.
	Products []string `json:"products,omitempty"`

	// list of product type IDs if the condition is applied on product types.
	ProductTypes []string `json:"product_types,omitempty"`

	// list of product collection IDs if the condition is applied on product collections.
	ProductCollections []string `json:"product_collections,omitempty"`

	// list of product tag IDs if the condition is applied on product tags.
	ProductTags []string `json:"product_tags,omitempty"`

	// list of customer group IDs if the condition is applied on customer groups.
	CustomerGroups []string `json:"customer_groups,omitempty"`
}

func (a *Admin) UpdateDiscountsCondition(context *fiber.Ctx) error {
	var b DiscountConditionsRequest

	if err := context.BodyParser(&b); err != nil {
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

	ConditionId, err := uuid.FromString(context.Params("condition_id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	m, err := a.r.ClientManager().GetDiscount(context.Context(), Id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	c, err := a.r.ClientManager().GetDiscountCondition(context.Context(), ConditionId)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	c.Id = ConditionId
	c.DiscountRuleId = m.RuleId

	if b.Products != nil {
		r, err := a.r.ClientManager().GetProductsByIds(context.Context(), b.Products)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "database_error",
			})
		}
		c.Products = append(c.Products, r...)
		c.Type = models.DiscountConditionTypeRroducts
	}
	if b.ProductTypes != nil {
		r, err := a.r.ClientManager().GetProductTypesByIds(context.Context(), b.ProductTypes)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "database_error",
			})
		}
		c.ProductTypes = append(c.ProductTypes, r...)
		c.Type = models.DiscountConditionTypeProductTypes
	}
	if b.ProductCollections != nil {
		r, err := a.r.ClientManager().GetProductCollectionsByIds(context.Context(), b.ProductCollections)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "database_error",
			})
		}
		c.ProductCollections = append(c.ProductCollections, r...)
		c.Type = models.DiscountConditionTypeProductCollections
	}
	if b.ProductTags != nil {
		r, err := a.r.ClientManager().GetProductTagsByIds(context.Context(), b.ProductTags)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "database_error",
			})
		}
		c.ProductTags = append(c.ProductTags, r...)
		c.Type = models.DiscountConditionTypeProductTags
	}
	if b.CustomerGroups != nil {
		r, err := a.r.ClientManager().GetCustomerGroupsByIds(context.Context(), b.CustomerGroups)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "database_error",
			})
		}
		c.CustomerGroups = append(c.CustomerGroups, r...)
		c.Type = models.DiscountConditionTypeCustomerGroups
	}

	_, err = a.r.ClientManager().UpdateDiscountCondition(context.Context(), c)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&m)
}
