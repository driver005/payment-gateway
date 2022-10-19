package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

type Charge struct {
	ID                   uuid.UUID `json:"id"`
	Object               string    `json:"object"`
	Amount               int       `json:"amount"`
	AmountCaptured       int       `json:"amount_captured"`
	AmountRefunded       int       `json:"amount_refunded"`
	Application          uuid.UUID `json:"application"`
	ApplicationFee       int       `json:"application_fee"`
	ApplicationFeeAmount int       `json:"application_fee_amount"`
	BalanceTransaction   uuid.UUID `json:"balance_transaction"`
	BillingDetails       struct {
		Address Address `json:"address"`
		Email   string  `json:"email"`
		Name    string  `json:"name"`
		Phone   string  `json:"phone"`
	} `json:"billing_details"`
	CalculatedStatementDescriptor string    `json:"calculated_statement_descriptor"`
	Captured                      bool      `json:"captured"`
	Created                       int       `json:"created"`
	Currency                      string    `json:"currency"`
	Customer                      uuid.UUID `json:"customer"`
	Description                   string    `json:"description"`
	Disputed                      bool      `json:"disputed"`
	FailureBalanceTransaction     uuid.UUID `json:"failure_balance_transaction"`
	FailureCode                   string    `json:"failure_code"`
	FailureMessage                string    `json:"failure_message"`
	FraudDetails                  struct {
		StripeReport string `json:"stripe_report"`
		UserReport   string `json:"user_report"`
	} `json:"fraud_details"`
	Invoice  string `json:"invoice"`
	Livemode bool   `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	OnBehalfOf string    `json:"on_behalf_of"`
	Order      uuid.UUID `json:"order"`
	Outcome    struct {
		NetworkStatus bool `json:"network_status"`
		Reason        bool `json:"reason"`
		RiskLevel     bool `json:"risk_level"`
		RiskScore     bool `json:"risk_score"`
		Rule          bool `json:"rule"`
		SellerMessage bool `json:"sellerMessage"`
		Type          bool `json:"type"`
	} `json:"outcome"`
	Paid                 bool      `json:"paid"`
	PaymentIntent        uuid.UUID `json:"payment_intent"`
	PaymentMethod        uuid.UUID `json:"payment_method"`
	PaymentMethodDetails struct {
		Account           Account           `json:"account"`
		AchCreditTransfer AchCreditTransfer `json:"ach_credit_transfer"`
		AchDebit          AchDebit          `json:"ach_debit"`
		AcssDebit         AcssDebit         `json:"acss_debit"`
		AfterpayClearpay  AfterpayClearpay  `json:"afterpay_clearpay"`
		Alipay            Alipay            `json:"alipay"`
		AuBecsDebit       AuBecsDebit       `json:"au_becs_debit"`
		BacsDebit         BacsDebit         `json:"bacs_debit"`
		Bankcontact       Bankcontact       `json:"bancontact"`
		Boleto            Boleto            `json:"boleto"`
		Card              Card              `json:"card"`
		CardPresent       CardPresent       `json:"card_present"`
		CustomerBalance   CustomerBalance   `json:"customer_balance"`
		Eps               Eps               `json:"eps"`
		Fpx               Fpx               `json:"fpx"`
		Giropay           Giropay           `json:"giropay"`
		Grabpay           Grabpay           `json:"grabpay"`
		Ideal             Ideal             `json:"ideal"`
		InteracPresent    InteracPresent    `json:"interac_present"`
		Klarna            Klarna            `json:"klarna"`
		Konbini           Konbini           `json:"konbini"`
		Multibanco        Multibanco        `json:"multibanco"`
		Oxxo              Oxxo              `json:"oxxo"`
		P24               P24               `json:"p24"`
		Paynow            Paynow            `json:"paynow"`
		SepaDebit         SepaDebit         `json:"sepa_debit"`
		Sofort            Sofort            `json:"sofort"`
		UsBankAccount     UsBankAccount     `json:"us_bank_account"`
		Wechat            Wechat            `json:"wechat"`
		WechatPay         WechatPay         `json:"wechat_pay"`
		Type              string            `json:"type"`
	} `json:"payment_method_details"`
	ReceiptEmail  string `json:"receipt_email"`
	ReceiptNumber string `json:"receipt_number"`
	ReceiptURL    string `json:"receipt_url"`
	Redaction     string `json:"redaction"`
	Refunded      bool   `json:"refunded"`
	Refunds       struct {
		Object  string      `json:"object"`
		Id      []uuid.UUID `json:"id"`
		HasMore bool        `json:"has_more"`
		URL     string      `json:"url"`
	} `json:"refunds"`
	Review   uuid.UUID `json:"review"`
	Shipping struct {
		Address struct {
			City    string `json:"city" db:"city"`
			State   string `json:"state" db:"state"`
			Country string `json:"country" db:"country"`
			Line1   string `json:"line1" db:"line1"`
			Line2   string `json:"line2" db:"line2"`
			Zip     string `json:"zip" db:"zip"`
		} `json:"address"`
		Carrier        string `json:"carrier"`
		Name           string `json:"name"`
		Phone          string `json:"phone"`
		TrackingNumber string `json:"tracking_number"`
	} `json:"shipping"`
	SourceTransfer            string `json:"source_transfer"`
	StatementDescriptor       string `json:"statement_descriptor"`
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix"`
	Status                    string `json:"status"`
	TransferData              struct {
		Amount      int       `json:"amount"`
		Destination uuid.UUID `json:"destination"`
	} `json:"transfer_data"`
	TransferGroup string `json:"transfer_group"`
}

func (c Charge) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required, is.UUID),
		validation.Field(&c.Object, validation.Required, validation.In("charge")),
		validation.Field(&c.Amount, validation.Required),
		validation.Field(&c.AmountCaptured, validation.Required),
		validation.Field(&c.AmountRefunded, validation.Required),
		validation.Field(&c.Application, is.UUID),
		validation.Field(&c.ApplicationFee),
		validation.Field(&c.ApplicationFeeAmount),
		validation.Field(&c.BalanceTransaction, validation.Required, is.UUID),
		validation.Field(&c.BillingDetails.Address.City, validation.Length(5, 50)),
		validation.Field(&c.BillingDetails.Address.Country, validation.Length(5, 50), is.CountryCode2),
		validation.Field(&c.BillingDetails.Address.Line1, validation.Length(5, 50)),
		validation.Field(&c.BillingDetails.Address.Line2, validation.Length(5, 50)),
		validation.Field(&c.BillingDetails.Address.State, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&c.BillingDetails.Address.Zip, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&c.BillingDetails.Name, validation.Required),
		validation.Field(&c.BillingDetails.Email, is.Email),
		validation.Field(&c.BillingDetails.Phone),
		validation.Field(&c.CalculatedStatementDescriptor),
		validation.Field(&c.Captured, validation.Required),
		validation.Field(&c.Captured, validation.Required),
		validation.Field(&c.Currency, validation.Required, is.CurrencyCode),
		validation.Field(&c.Customer, is.UUID),
		validation.Field(&c.Description, validation.Required),
		validation.Field(&c.Disputed, validation.Required),
		validation.Field(&c.FailureBalanceTransaction, is.UUID),
		validation.Field(&c.FailureCode),
		validation.Field(&c.FailureMessage),
		validation.Field(&c.FailureMessage),
		validation.Field(&c.FraudDetails.StripeReport, validation.In("fraudulent")),
		validation.Field(&c.FraudDetails.UserReport, validation.In("safe", "fraudulent")),
		validation.Field(&c.Livemode, validation.Required),
		validation.Field(&c.OnBehalfOf),
		validation.Field(&c.Order, is.UUID),
		validation.Field(&c.Outcome.NetworkStatus, validation.In("approved_by_network", "declined_by_network", "not_sent_to_network", "reversed_after_approval", "reversed_after_approval ")),
		validation.Field(&c.Outcome.Reason, validation.In("type", "highest_risk_level", "elevated_risk_level", "rule")),
		validation.Field(&c.Outcome.RiskLevel, validation.In("normal", "elevated", "highest", "not_assessed", "unknown")),
		validation.Field(&c.Outcome.RiskScore, validation.Min(0), validation.Max(100)),
		validation.Field(&c.Outcome.SellerMessage),
		validation.Field(&c.Outcome.Type, validation.In("authorized", "manual_review", "issuer_declined", "blocked", "invalid")),

		validation.Field(&c.Paid, validation.Required),
		validation.Field(&c.PaymentIntent, is.UUID),
		validation.Field(&c.PaymentMethod, validation.Required, is.UUID),

		validation.Field(&c.ReceiptEmail, is.Email),
		validation.Field(&c.ReceiptNumber),
		validation.Field(&c.ReceiptURL, is.URL),
		validation.Field(&c.Redaction),
		validation.Field(&c.Refunded),
		validation.Field(&c.Refunds.Id, is.UUID),
		validation.Field(&c.Refunds.HasMore),
		validation.Field(&c.Refunds.Object),
		validation.Field(&c.Refunds.URL),
		validation.Field(&c.Review, is.UUID),
		validation.Field(&c.Shipping),
		validation.Field(&c.Shipping.Address.City, validation.Length(5, 50)),
		validation.Field(&c.Shipping.Address.Country, validation.Length(5, 50), is.CountryCode2),
		validation.Field(&c.Shipping.Address.Line1, validation.Length(5, 50)),
		validation.Field(&c.Shipping.Address.Line2, validation.Length(5, 50)),
		validation.Field(&c.Shipping.Address.State, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&c.Shipping.Address.Zip, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&c.Shipping.Carrier),
		validation.Field(&c.Shipping.Name),
		validation.Field(&c.Shipping.Phone),
		validation.Field(&c.Shipping.TrackingNumber),
		validation.Field(&c.SourceTransfer),
		validation.Field(&c.StatementDescriptor),
		validation.Field(&c.StatementDescriptorSuffix),
		validation.Field(&c.Status, validation.In("succeeded", "pending", "failed")),
		validation.Field(&c.TransferData.Amount),
		validation.Field(&c.TransferData.Destination, is.UUID),
		validation.Field(&c.TransferGroup),
	)
}
