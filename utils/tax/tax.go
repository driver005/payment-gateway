package tax

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/google/uuid"
)

type AutomaticTax struct {
	core.Model

	// Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://stripe.com/docs/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices.
	Enabled bool `json:"enabled,omitempty"`
	// The status of the most recent automated tax calculation for this invoice.
	Status string `json:"status,omitempty"`
}

// TaxDeductedAtSource
type TaxDeductedAtSource struct {
	core.Model

	// The end of the invoicing period. This TDS applies to Stripe fees collected during this invoicing period.
	PeriodEnd int `json:"period_end,omitempty"`
	// The start of the invoicing period. This TDS applies to Stripe fees collected during this invoicing period.
	PeriodStart int `json:"period_start,omitempty"`
	// The TAN that was supplied to Stripe when TDS was assessed
	TaxDeductionAccountNumber string `json:"tax_deduction_account_number,omitempty"`
}

// TaxCode [Tax codes](https://stripe.com/docs/tax/tax-categories) classify goods and services for tax purposes.
type TaxCode struct {
	core.Model

	// A detailed description of which types of products the tax code represents.
	Description string `json:"description,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// A short name for the tax code.
	Name string `json:"name,omitempty"`
}

// TaxRate Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.  Related guide: [Tax Rates](https://stripe.com/docs/billing/taxes/tax-rates).
type TaxRate struct {
	core.Model

	// The amount, in %s, of the tax.
	Amount int `json:"amount,omitempty"`
	// Defaults to `true`. When set to `false`, this tax rate cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
	Active bool `json:"active,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
	Description string `json:"description,omitempty"`
	// The display name of the tax rates as it will appear to your customer on their receipt email, PDF, and the hosted invoice page.
	DisplayName string `json:"display_name,omitempty"`
	// This specifies if the tax rate is inclusive or exclusive.
	Inclusive bool `json:"inclusive,omitempty"`
	// The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
	Jurisdiction string `json:"jurisdiction,omitempty"`
	// This represents the tax rate percent out of 100.
	Percentage float32 `json:"percentage,omitempty"`
	// [ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2:US), without country prefix. For example, \"NY\" for New York, United States.
	State string `json:"state,omitempty"`
	// The high-level tax type, such as `vat` or `sales_tax`.
	TaxType string `json:"tax_type,omitempty"`
}

// TaxId You can add one or multiple tax IDs to a [customer](https://stripe.com/docs/api/customers). A customer's tax IDs are displayed on invoices and credit notes issued for the customer.  Related guide: [Customer Tax Identification Numbers](https://stripe.com/docs/billing/taxes/tax-ids).
type TaxId struct {
	core.Model

	// Two-letter ISO code representing the country of the tax ID.
	Country string `json:"country,omitempty"`
	// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	// Verification status, one of `pending`, `verified`, `unverified`, or `unavailable`.
	Status string `json:"status,omitempty"`
	// Verified address.
	VerifiedAddress string `json:"verified_address,omitempty"`
	// Verified name.
	VerifiedName string `json:"verified_name,omitempty"`

	CustomerId uuid.UUID          `json:"customer_id" database:"default:null"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
