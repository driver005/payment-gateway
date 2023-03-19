package customer

import "database/sql/driver"

type Type string

const (
	TypeAdjustment            Type = "adjustment"
	TypeAppliedToInvoice      Type = "applied_to_invoice"
	TypeCreditNote            Type = "credit_note "
	TypeInitial               Type = "initial"
	TypeInvoiceOverpaid       Type = "invoice_overpaid"
	TypeInvoiceTooLarge       Type = "invoice_too_large"
	TypeInvoiceTooSmall       Type = "invoice_too_small"
	TypeunSpentReceiverCredit Type = "unspent_receiver_credit"
	TypeunAppliedFromInvoice  Type = "unapplied_from_invoice"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}
