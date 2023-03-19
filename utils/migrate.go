package utils

import (
	"github.com/driver005/gateway/utils/account"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/country"
	"github.com/driver005/gateway/utils/currency"
	"github.com/driver005/gateway/utils/email"
	"github.com/driver005/gateway/utils/errors"
	"github.com/driver005/gateway/utils/fee"
	"github.com/driver005/gateway/utils/fundingInstruction"
	"github.com/driver005/gateway/utils/paymentFlow"
	"github.com/driver005/gateway/utils/radar"
	"github.com/driver005/gateway/utils/region"
	"github.com/driver005/gateway/utils/review"
	"github.com/driver005/gateway/utils/rule"
	"github.com/driver005/gateway/utils/shipping"
	"github.com/driver005/gateway/utils/tax"
)

func (h *Handler) Migrate() {
	account.Migrate(h.r)
	address.Migrate(h.r)
	country.Migrate(h.r)
	currency.Migrate(h.r)
	email.Migrate(h.r)
	errors.Migrate(h.r)
	fee.Migrate(h.r)
	fundingInstruction.Migrate(h.r)
	paymentFlow.Migrate(h.r)
	radar.Migrate(h.r)
	region.Migrate(h.r)
	review.Migrate(h.r)
	rule.Migrate(h.r)
	shipping.Migrate(h.r)
	tax.Migrate(h.r)
}
