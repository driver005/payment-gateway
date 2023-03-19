package dispute

import "database/sql/driver"

type Status string

const (
	StatusWarningNeedsResponse Status = "warning_needs_response"
	StatusWarningUnderReview   Status = "warning_under_review"
	StatusWarningClosed        Status = "warning_closed"
	StatusNeedsResponse        Status = "needs_response"
	StatusUnderReview          Status = "under_review"
	StatusChargeRefunded       Status = "charge_refunded"
	StatusWon                  Status = "won"
	StatusLost                 Status = "lost"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type Reason string

const (
	ReasonBankCannotProcess       Reason = "bank_cannot_process"
	ReasonCheckReturned           Reason = "check_returned"
	ReasonCreditNotProcessed      Reason = "credit_not_processed"
	ReasonCustomerInitiated       Reason = "customer_initiated"
	ReasonDebitNotAuthorized      Reason = "debit_not_authorized"
	ReasonDuplicate               Reason = "duplicate"
	ReasonFraudulent              Reason = "fraudulent"
	ReasonGeneral                 Reason = "general"
	ReasonIncorrectAccountDetails Reason = "incorrect_account_details"
	ReasonInsufficientFunds       Reason = "insufficient_funds"
	ReasonProductNotReceived      Reason = "product_not_received"
	ReasonProductUnacceptable     Reason = "product_unacceptable"
	ReasonSubscriptionCanceled    Reason = "subscription_canceled"
	ReasonUnrecognized            Reason = "unrecognized"
)

func (ct *Reason) Scan(value interface{}) error {
	*ct = Reason(value.(string))
	return nil
}

func (ct Reason) Value() (driver.Value, error) {
	return string(ct), nil
}
