package admin

import (
	"github.com/driver005/gateway/handler"
	"github.com/gofiber/fiber/v2"
)

const (
	Currencies       = "/currencies"
	BtcpayPath       = "/btcpay"
	Products         = "/products"
	PriceLists       = "/price-lists"
	ShippingProfiles = "/shipping-profiles"
	BatchJobs        = "/batch-jobs"
	Claims           = "/orders"
	Collections      = "/collections"
	Customers        = "/customers"
	CustomerGroups   = "/customer-groups"
	Discounts        = "/discounts"
)

type Handler struct {
	a Admin
}

func NewHandler(r handler.Registry) *Handler {
	return &Handler{
		a: Admin{r: r},
	}
}

func (h *Handler) Routes(public fiber.Router) {
	h.CurrencyRoutes(public)

	h.PriceListRoutes(public)

	h.ProductRoutes(public)
	h.ProductOptionRoutes(public)
	h.ProductTagRoutes(public)
	h.ProductTypeRoutes(public)
	h.ProductVariantRoutes(public)

	h.ShippingProfileRoutes(public)
}

func (h *Handler) BatchJobRoutes(public fiber.Router) {
	public.Get(BatchJobs, h.a.ListBatchJob)
	public.Post(BatchJobs, h.a.CreateBatchJob)
	public.Get(BatchJobs+"/:id", h.a.GetBatchJob)
	public.Get(BatchJobs+"/:id/confirm", h.a.ConfirmBatchJob)
	public.Get(BatchJobs+"/:id/cancel", h.a.CancelBatchJob)
}

func (h *Handler) ClaimRoutes(public fiber.Router) {
	public.Post(Claims+"/:id/claims", h.a.CreateClaim)
	public.Post(Claims+"/:id/claims/:claim_id", h.a.UpdateClaim)
	public.Post(Claims+"/:id/claims/:claim_id/cancel", h.a.CancelClaim)
	public.Post(Claims+"/:id/claims/:claim_id/shipments", h.a.UpdateClaimShipment)
}

func (h *Handler) CollectionRoutes(public fiber.Router) {
	public.Get(Collections, h.a.ListCollection)
	public.Post(Collections, h.a.CreateCollection)
	public.Get(Collections+"/:id", h.a.GetCollection)
	public.Delete(Collections+"/:id", h.a.DeleteCollection)
	public.Post(Collections+"/:id", h.a.UpdateCollection)
	public.Post(Collections+"/:id/products/batch", h.a.UpdateCollectionProducts)
	public.Delete(Collections+"/:id/products/batch", h.a.DeleteCollectionProducts)
}

func (h *Handler) CurrencyRoutes(public fiber.Router) {
	public.Get(Currencies, h.a.ListCurrency)
	public.Get(Currencies+"/:code", h.a.GetCurrency)
	public.Post(Currencies+"/:code", h.a.UpdateCurrency)
}

func (h *Handler) CustomerRoutes(public fiber.Router) {
	public.Get(Customers, h.a.ListCustomer)
	public.Post(Customers, h.a.CreateCustomer)
	public.Delete(Customers+"/:id", h.a.DeleteCustomer)
	public.Get(Customers+"/:id", h.a.GetCustomer)
	public.Post(Customers+"/:id", h.a.UpdateCustomer)
}

func (h *Handler) CustomerGroupRoutes(public fiber.Router) {
	public.Get(CustomerGroups, h.a.ListCustomerGroup)
	public.Post(CustomerGroups, h.a.CreateCustomerGroup)
	public.Get(CustomerGroups+"/:id", h.a.GetCustomerGroup)
	public.Delete(CustomerGroups+"/:id", h.a.DeleteCustomerGroup)
	public.Post(CustomerGroups+"/:id", h.a.UpdateCustomerGroup)
	public.Post(CustomerGroups+"/:id/customers/batch", h.a.UpdateCustomerGroupCustomers)
	public.Delete(CustomerGroups+"/:id/customers/batch", h.a.DeleteCustomerGroupCustomers)
}

func (h *Handler) DiscountRoutes(public fiber.Router) {
	public.Get(Discounts, h.a.ListDiscount)
	public.Post(Discounts, h.a.CreateDiscount)
	public.Get(Discounts+"/:id", h.a.GetDiscount)
	public.Get(Discounts+"/code/:code", h.a.GetDiscountByCode)
	public.Delete(Discounts+"/:id", h.a.DeleteDiscount)
	public.Post(Discounts+"/:id", h.a.UpdateDiscount)
	public.Post(Discounts+"/:id/regions/:region_id", h.a.UpdateRegion)
	public.Delete(Discounts+"/:id/regions/:region_id", h.a.DeleteRegion)
	public.Post(Discounts+"/:id/dynamic-codes", h.a.UpdateDynamicCode)
	public.Delete(Discounts+"/:id/dynamic-codes/:code", h.a.DeleteDynamicCode)
	public.Post(Discounts+"/:id/conditions/:condition_id", h.a.UpdateDiscountsCondition)
}

func (h *Handler) DiscountConditionRoutes(public fiber.Router) {
	// public.Get(Discounts, h.a.ListDiscountCondition)
	public.Post(Discounts, h.a.CreateDiscountCondition)
	public.Get(Discounts+"/:id", h.a.GetDiscountCondition)
	public.Delete(Discounts+"/:id", h.a.DeleteDiscountCondition)
}

func (h *Handler) PriceListRoutes(public fiber.Router) {
	public.Get(PriceLists, h.a.ListPriceList)
	public.Post(PriceLists, h.a.CreatePriceList)
	public.Delete(PriceLists+"/:id", h.a.DeletePriceList)
	public.Get(PriceLists+"/:id", h.a.GetPriceList)
	public.Post(PriceLists+"/:id", h.a.UpdatePriceList)
	// public.Delete(PriceLists+"/:id/prices/batch", h.a.UpdatePrices)
	// public.Post(PriceLists+"/:id/prices/batch", h.UpdatePriceList)
	// public.Get(PriceLists+"/:id/products", h.a.ListProducts)
	// public.Delete(PriceLists+"/:id/products/:product_id/prices", wrapper.DeletePriceListsPriceListProductsProductPrices)
	// public.Delete(PriceLists+"/:id/variants/:variant_id/prices", wrapper.DeletePriceListsPriceListVariantsVariantPrices)
}

func (h *Handler) ProductRoutes(public fiber.Router) {
	public.Post(Products, h.a.CreateProduct)
	public.Get(Products, h.a.ListProduct)
	public.Delete(Products+"/:id", h.a.DeleteProduct)
	public.Get(Products+"/:id", h.a.GetProduct)
	public.Post(Products+"/:id", h.a.UpdateProduct)
	public.Post(Products+"/:id/metadata", h.a.CreateProductMetadata)
}

func (h *Handler) ProductOptionRoutes(public fiber.Router) {
	public.Post(Products+"/:id/options", h.a.CreateProductOption)
	public.Delete(Products+"/:id/options/:option_id", h.a.DeleteProductOption)
	public.Post(Products+"/:id/options/:option_id", h.a.UpdateProductOption)
}

func (h *Handler) ProductTagRoutes(public fiber.Router) {
	public.Get(Products+"/product-tags", h.a.ListProductTag)
	public.Get(Products+"/tag-usage", h.a.GetProductTagUsageCount)
}

func (h *Handler) ProductTypeRoutes(public fiber.Router) {
	public.Get(Products+"/product-types", h.a.ListProductType)
	public.Get(Products+"/types", h.a.ListProductType)
}

func (h *Handler) ProductVariantRoutes(public fiber.Router) {
	public.Get(Products+"/:id/variants", h.a.ListProductVariant)
	public.Post(Products+"/:id/variants", h.a.CreateProductVariant)
	public.Delete(Products+"/:id/variants/:variant_id", h.a.DeleteProductVariant)
	public.Post(Products+"/:id/variants/:variant_id", h.a.UpdateProductVariant)
	public.Get("/variants", h.a.ListProductVariants)
}

func (h *Handler) ShippingProfileRoutes(public fiber.Router) {
	public.Post(ShippingProfiles, h.a.CreateShippingProfile)
	public.Get(ShippingProfiles, h.a.ListShippingProfile)
	public.Delete(ShippingProfiles+"/:id", h.a.DeleteShippingProfile)
	public.Get(ShippingProfiles+"/:id", h.a.GetShippingProfile)
	public.Post(ShippingProfiles+"/:id", h.a.UpdateShippingProfile)
}
