package balance

import "database/sql/driver"

type Type string

const (
	TypeAdjustment                  Type = "adjustment"
	TypeAdvance                     Type = "advance"
	TypeAdvanceFunding              Type = "advance_funding"
	TypeAnticipationRepayment       Type = "anticipation_repayment"
	TypeApplicationFee              Type = "application_fee"
	TypeApplicationFeeRefund        Type = "application_fee_refund"
	TypeCharge                      Type = "charge"
	TypeConnectCollectionTransfer   Type = "connect_collection_transfer"
	TypeContribution                Type = "contribution"
	TypeIssuingAuthorizationHold    Type = "issuing_authorization_hold"
	TypeIssuingAuthorizationRelease Type = "issuing_authorization_release"
	TypeIssuingDispute              Type = "issuing_dispute"
	TypeIssuingTransaction          Type = "issuing_transaction"
	TypePayment                     Type = "payment"
	TypePaymentFailureRefund        Type = "payment_failure_refund"
	TypePaymentRefund               Type = "payment_refund"
	TypePayout                      Type = "payout"
	TypePayoutCancel                Type = "payout_cancel"
	TypePayoutFailure               Type = "payout_failure"
	TypeRefund                      Type = "refund"
	TypeRefundFailure               Type = "refund_failure"
	TypeReserveTransaction          Type = "reserve_transaction"
	TypeReservedFunds               Type = "reserved_funds"
	TypeReportingCategory           Type = "reporting_category"
	TypeStripeFee                   Type = "stripe_fee"
	TypeStripeFxFee                 Type = "stripe_fx_fee"
	TypeTaxFee                      Type = "tax_fee"
	TypeTopup                       Type = "topup"
	TypeTopupReversal               Type = "topup_reversal"
	TypeTransfer                    Type = "transfer"
	TypeTransferCancel              Type = "transfer_cancel"
	TypeTransferFailure             Type = "transfer_failure"
	TypeTransferRefund              Type = "transfer_refund"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type Status string

const (
	StatusAvailable Status = "available"
	StatusPending   Status = "pending"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}
