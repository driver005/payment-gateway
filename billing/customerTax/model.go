package customerTax

import "github.com/driver005/gateway/internal/customer"

// CustomerTaxLocation
type CustomerTaxLocation struct {
	// The customer's country as identified by Stripe Tax.
	Country string `json:"country"`
	// The data source used to infer the customer's location.
	Source string `json:"source"`
	// The customer's state, county, province, or region as identified by Stripe Tax.
	State string `json:"state,omitempty"`
}

// CustomerTax
type CustomerTax struct {
	// Surfaces if automatic tax computation is possible given the current customer location information.
	AutomaticTax string `json:"automatic_tax"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IpAddress string `json:"ip_address,omitempty"`
	Location CustomerTaxLocation `json:"location,omitempty"`
}

// TaxIdVerification
type TaxIdVerification struct {
	// Verification status, one of `pending`, `verified`, `unverified`, or `unavailable`.
	Status string `json:"status"`
	// Verified address.
	VerifiedAddress string `json:"verified_address,omitempty"`
	// Verified name.
	VerifiedName string `json:"verified_name,omitempty"`
}

// TaxId You can add one or multiple tax IDs to a [customer](https://stripe.com/docs/api/customers). A customer's tax IDs are displayed on invoices and credit notes issued for the customer.  Related guide: [Customer Tax Identification Numbers](https://stripe.com/docs/billing/taxes/tax-ids).
type TaxId struct {
	// Two-letter ISO code representing the country of the tax ID.
	Country string `json:"country,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int32 `json:"created"`
	Customer customer.Customer `json:"customer,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `eg_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ph_tin`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tr_tin`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`. Note that some legacy tax IDs have type `unknown`
	Type string `json:"type"`
	// Value of the tax ID.
	Value string `json:"value"`
	Verification TaxIdVerification `json:"verification,omitempty"`
}

