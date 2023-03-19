package invoice

import "database/sql/driver"

type Status string

const (
	StatusDraft         Status = "draft"
	StatusOpen          Status = "open"
	StatusPaid          Status = "paid"
	StatusUncollectible Status = "uncollectible"
	StatusVoid          Status = "void"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type CollectionMethod string

const (
	CollectionMethodChargeAutomatically CollectionMethod = "charge_automatically"
	CollectionMethodSendInvoiced        CollectionMethod = "send_invoice"
)

func (ct *CollectionMethod) Scan(value interface{}) error {
	*ct = CollectionMethod(value.(string))
	return nil
}

func (ct CollectionMethod) Value() (driver.Value, error) {
	return string(ct), nil
}

type BillingReason string

const (
	BillingReasonSubscriptionCycle     BillingReason = "subscription_cycle"
	BillingReasonSubscriptionCreate    BillingReason = "subscription_create"
	BillingReasonSubscriptionUpdate    BillingReason = "subscription_update"
	BillingReasonSubscriptionThreshold BillingReason = "subscription_threshold"
	BillingReasonSubscription          BillingReason = "subscription"
	BillingReasonManual                BillingReason = "manual"
	BillingReasonUpcoming              BillingReason = "upcoming"
)

func (ct *BillingReason) Scan(value interface{}) error {
	*ct = BillingReason(value.(string))
	return nil
}

func (ct BillingReason) Value() (driver.Value, error) {
	return string(ct), nil
}
