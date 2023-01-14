package models

import (
	"github.com/driver005/gateway/core"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type Charge struct {
	core.Model

	Amount                        int       `json:"amount"`
	AmountCaptured                int       `json:"amount_captured"`
	AmountRefunded                int       `json:"amount_refunded"`
	BalanceTransaction            uuid.UUID `json:"balance_transaction"`
	BillingDetails                *Address  `json:"billing_details"  database:"foreignKey:id"`
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
	// Invoice  string `json:"invoice"`
	// Order      uuid.UUID `json:"order"`
	// Outcome    struct {
	// 	NetworkStatus bool `json:"network_status"`
	// 	Reason        bool `json:"reason"`
	// 	RiskLevel     bool `json:"risk_level"`
	// 	RiskScore     bool `json:"risk_score"`
	// 	Rule          bool `json:"rule"`
	// 	SellerMessage bool `json:"sellerMessage"`
	// 	Type          bool `json:"type"`
	// } `json:"outcome"`
	Paid                 bool           `json:"paid"`
	PaymentIntent        *PaymentIntent `json:"payment_intent" database:"foreignKey:id"`
	PaymentMethod        *PaymentMethod `json:"payment_method" database:"foreignKey:id"`
	PaymentMethodDetails struct {
		Account           *Account           `json:"account" database:"foreignKey:id"`
		AchCreditTransfer *AchCreditTransfer `json:"ach_credit_transfer" database:"foreignKey:id"`
		AchDebit          *AchDebit          `json:"ach_debit" database:"foreignKey:id"`
		AcssDebit         *AcssDebit         `json:"acss_debit" database:"foreignKey:id"`
		AfterpayClearpay  *AfterpayClearpay  `json:"afterpay_clearpay" database:"foreignKey:id"`
		Alipay            *Alipay            `json:"alipay" database:"foreignKey:id"`
		AuBecsDebit       *AuBecsDebit       `json:"au_becs_debit" database:"foreignKey:id"`
		BacsDebit         *BacsDebit         `json:"bacs_debit" database:"foreignKey:id"`
		Bankcontact       *Bankcontact       `json:"bancontact" database:"foreignKey:id"`
		Boleto            *Boleto            `json:"boleto" database:"foreignKey:id"`
		Card              *Card              `json:"card" database:"foreignKey:id"`
		CardPresent       *CardPresent       `json:"card_present" database:"foreignKey:id"`
		CustomerBalance   *CustomerBalance   `json:"customer_balance" database:"foreignKey:id"`
		Eps               *Eps               `json:"eps" database:"foreignKey:id"`
		Fpx               *Fpx               `json:"fpx" database:"foreignKey:id"`
		Giropay           *Giropay           `json:"giropay" database:"foreignKey:id"`
		Grabpay           *Grabpay           `json:"grabpay" database:"foreignKey:id"`
		Ideal             *Ideal             `json:"ideal" database:"foreignKey:id"`
		InteracPresent    *InteracPresent    `json:"interac_present" database:"foreignKey:id"`
		Klarna            *Klarna            `json:"klarna" database:"foreignKey:id"`
		Konbini           *Konbini           `json:"konbini" database:"foreignKey:id"`
		Multibanco        *Multibanco        `json:"multibanco" database:"foreignKey:id"`
		Oxxo              *Oxxo              `json:"oxxo" database:"foreignKey:id"`
		P24               *P24               `json:"p24" database:"foreignKey:id"`
		Paynow            *Paynow            `json:"paynow" database:"foreignKey:id"`
		SepaDebit         *SepaDebit         `json:"sepa_debit" database:"foreignKey:id"`
		Sofort            *Sofort            `json:"sofort" database:"foreignKey:id"`
		UsBankAccount     *UsBankAccount     `json:"us_bank_account" database:"foreignKey:id"`
		Wechat            *Wechat            `json:"wechat" database:"foreignKey:id"`
		WechatPay         *WechatPay         `json:"wechat_pay" database:"foreignKey:id"`
		Type              string             `json:"type"`
	} `json:"payment_method_details"`
	ReceiptEmail  string `json:"receipt_email"`
	ReceiptNumber string `json:"receipt_number"`
	ReceiptURL    string `json:"receipt_url"`
	Redaction     string `json:"redaction"`
	Refunded      bool   `json:"refunded"`
	// Refunds       struct {
	// 	Object  string      `json:"object"`
	// 	Id      []uuid.UUID `json:"id"`
	// 	HasMore bool        `json:"has_more"`
	// 	URL     string      `json:"url"`
	// } `json:"refunds"`
	Review                    uuid.UUID `json:"review"`
	Shipping                  *Address  `json:"shipping"`
	StatementDescriptor       string    `json:"statement_descriptor"`
	StatementDescriptorSuffix string    `json:"statement_descriptor_suffix"`
	Status                    string    `json:"status"`
}

func (c Charge) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Id, validation.Required, is.UUID),
		validation.Field(&c.Object, validation.Required, validation.In("charge")),
		validation.Field(&c.Amount, validation.Required),
		validation.Field(&c.AmountCaptured, validation.Required),
		validation.Field(&c.AmountRefunded, validation.Required),
		// validation.Field(&c.Application, is.UUID),
		// validation.Field(&c.ApplicationFee),
		// validation.Field(&c.ApplicationFeeAmount),
		validation.Field(&c.BalanceTransaction, validation.Required, is.UUID),
		// validation.Field(&c.BillingDetails.Address.City, validation.Length(5, 50)),
		// validation.Field(&c.BillingDetails.Address.Country, validation.Length(5, 50), is.CountryCode2),
		// validation.Field(&c.BillingDetails.Name, validation.Required),
		// validation.Field(&c.BillingDetails.Email, is.Email),
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
		// validation.Field(&c.OnBehalfOf),
		// validation.Field(&c.Order, is.UUID),
		// validation.Field(&c.Outcome.NetworkStatus, validation.In("approved_by_network", "declined_by_network", "not_sent_to_network", "reversed_after_approval", "reversed_after_approval ")),
		// validation.Field(&c.Outcome.Reason, validation.In("type", "highest_risk_level", "elevated_risk_level", "rule")),
		// validation.Field(&c.Outcome.RiskLevel, validation.In("normal", "elevated", "highest", "not_assessed", "unknown")),
		// validation.Field(&c.Outcome.RiskScore, validation.Min(0), validation.Max(100)),
		// validation.Field(&c.Outcome.SellerMessage),
		// validation.Field(&c.Outcome.Type, validation.In("authorized", "manual_review", "issuer_declined", "blocked", "invalid")),

		validation.Field(&c.Paid, validation.Required),
		validation.Field(&c.PaymentIntent, is.UUID),
		validation.Field(&c.PaymentMethod, validation.Required, is.UUID),

		validation.Field(&c.ReceiptEmail, is.Email),
		validation.Field(&c.ReceiptNumber),
		validation.Field(&c.ReceiptURL, is.URL),
		validation.Field(&c.Redaction),
		validation.Field(&c.Refunded),
		// validation.Field(&c.Refunds.Id, is.UUID),
		// validation.Field(&c.Refunds.HasMore),
		// validation.Field(&c.Refunds.Object),
		// validation.Field(&c.Refunds.URL),
		validation.Field(&c.Review, is.UUID),
		validation.Field(&c.Shipping),
		// validation.Field(&c.Shipping.Address.City, validation.Length(5, 50)),
		// validation.Field(&c.Shipping.Address.Country, validation.Length(5, 50), is.CountryCode2),
		// validation.Field(&c.Shipping.Address.Line1, validation.Length(5, 50)),
		// validation.Field(&c.Shipping.Address.Line2, validation.Length(5, 50)),
		// validation.Field(&c.Shipping.Address.State, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// validation.Field(&c.Shipping.Address.Zip, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		// validation.Field(&c.Shipping.Carrier),
		// validation.Field(&c.Shipping.Name),
		// validation.Field(&c.Shipping.Phone),
		// validation.Field(&c.Shipping.TrackingNumber),
		validation.Field(&c.StatementDescriptor),
		validation.Field(&c.StatementDescriptorSuffix),
		validation.Field(&c.Status, validation.In("succeeded", "pending", "failed")),
	)
}
