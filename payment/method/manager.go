package method

import (
	"github.com/driver005/gateway/payment/method/methods"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*PaymentMethod, error) {
	type Alias PaymentMethod
	var m PaymentMethod
	var err error

	model := struct {
		*Alias
		Customer                    uuid.NullUUID `json:"customer,omitempty"`
		AccountNumber               string        `json:"account_number,omitempty"`
		AccountHolderType           string        `json:"account_holder_type,omitempty"`
		AccountType                 string        `json:"account_type,omitempty"`
		FinancialConnectionsAccount string        `json:"financial_connections_account,omitempty"`
		RoutingNumber               string        `json:"routing_number,omitempty"`
		InstitutionNumber           string        `json:"institution_number,omitempty"`
		TransitNumber               string        `json:"transit_number,omitempty"`
		BsbNumber                   string        `json:"bsb_number,omitempty"`
		SortCode                    string        `json:"sort_code,omitempty"`
		TaxId                       string        `json:"tax_id,omitempty"`
		ExpMonth                    int           `json:"exp_month,omitempty"`
		ExpYear                     int           `json:"exp_year,omitempty"`
		Number                      string        `json:"number,omitempty"`
		Cvc                         string        `json:"cvc,omitempty"`
		Bank                        string        `json:"bank,omitempty"`
		Day                         int           `json:"day,omitempty"`
		Month                       int           `json:"month,omitempty"`
		Year                        int           `json:"year,omitempty"`
		Iban                        string        `json:"iban,omitempty"`
		Country                     string        `json:"country,omitempty"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context.Context(), model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}

	if m.Type == "acss_debit" {
		m.AcssDebit = &methods.PaymentMethodAcssDebit{
			// AccountNumber: model.AccountNumber,
			InstitutionNumber: model.InstitutionNumber,
			TransitNumber:     model.TransitNumber,
		}
	} else if m.Type == "affirm" {
		m.Affirm = &methods.PaymentMethodAffirm{}
	} else if m.Type == "afterpay_clearpay" {
		m.AfterpayClearpay = &methods.PaymentMethodAfterpayClearpay{}
	} else if m.Type == "alipay" {
		m.Alipay = &methods.PaymentMethodAlipay{}
	} else if m.Type == "au_becs_debit" {
		m.AuBecsDebit = &methods.PaymentMethodAuBecsDebit{
			// AccountNumber: model.AccountNumber,
			BsbNumber: model.BsbNumber,
		}
	} else if m.Type == "bacs_debit" {
		m.BacsDebit = &methods.PaymentMethodBacsDebit{
			// AccountNumber: model.AccountNumber,
			SortCode: model.SortCode,
		}
	} else if m.Type == "bancontact" {
		m.Bancontact = &methods.PaymentMethodBancontact{}
	} else if m.Type == "blik" {
		m.Blik = &methods.PaymentMethodBlik{}
	} else if m.Type == "boleto" {
		m.Boleto = &methods.PaymentMethodBoleto{
			TaxId: model.TaxId,
		}
	} else if m.Type == "btcpay" {
		m.BtcPay = &methods.PaymentMethodBtcPay{}
	} else if m.Type == "card" {
		m.Card = &methods.PaymentMethodCard{
			ExpMonth: model.ExpMonth,
			ExpYear:  model.ExpYear,
			// Number: model.Number,
			// Cvc: model.Cvc,
		}
	} else if m.Type == "card_present" {
		m.CardPresent = &methods.PaymentMethodCardPresent{
			ExpMonth: model.ExpMonth,
			ExpYear:  model.ExpYear,
			// Number: model.Number,
			// Cvc: model.Cvc,
		}
	} else if m.Type == "customer_balance" {
		m.CustomerBalance = &methods.PaymentMethodCustomerBalance{}
	} else if m.Type == "eps" {
		m.Eps = &methods.PaymentMethodEps{
			Bank: model.Bank,
		}
	} else if m.Type == "fpx" {
		m.Fpx = &methods.PaymentMethodFpx{
			Bank: model.Bank,
		}
	} else if m.Type == "giropay" {
		m.Giropay = &methods.PaymentMethodGiropay{}
	} else if m.Type == "grabpay" {
		m.Grabpay = &methods.PaymentMethodGrabpay{}
	} else if m.Type == "ideal" {
		m.Ideal = &methods.PaymentMethodIdeal{
			Bank: model.Bank,
		}
	} else if m.Type == "interac_present" {
		m.InteracPresent = &methods.PaymentMethodInteracPresent{
			ExpMonth: model.ExpMonth,
			ExpYear:  model.ExpYear,
			// Number: model.Number,
			// Cvc: model.Cvc,
		}
	} else if m.Type == "klarna" {
		m.Klarna = &methods.PaymentMethodKlarna{
			Day:   model.Day,
			Month: model.Month,
			Year:  model.Year,
		}
	} else if m.Type == "konbini" {
		m.Konbini = &methods.PaymentMethodKonbini{}
	} else if m.Type == "link" {
		m.Link = &methods.PaymentMethodLink{}
	} else if m.Type == "oxxo" {
		m.Oxxo = &methods.PaymentMethodOxxo{}
	} else if m.Type == "p24" {
		m.P24 = &methods.PaymentMethodP24{
			Bank: model.Bank,
		}
	} else if m.Type == "paynow" {
		m.Paynow = &methods.PaymentMethodPaynow{}
	} else if m.Type == "pix" {
		m.Pix = &methods.PaymentMethodPix{}
	} else if m.Type == "promptpay" {
		m.Promptpay = &methods.PaymentMethodPromptpay{}
	} else if m.Type == "sepa_debit" {
		m.SepaDebit = &methods.PaymentMethodSepaDebit{}
		// add SepaDebit
	} else if m.Type == "sofort" {
		m.Sofort = &methods.PaymentMethodSofort{
			Country: model.Country,
		}
	} else if m.Type == "us_bank_account" {
		m.UsBankAccount = &methods.PaymentMethodUsBankAccount{
			AccountHolderType: model.AccountHolderType,
			// AccountNumber: model.AccountNumber,
			AccountType:                 model.AccountType,
			FinancialConnectionsAccount: model.FinancialConnectionsAccount,
			RoutingNumber:               model.RoutingNumber,
		}
	} else if m.Type == "wechat_pay" {
		m.WechatPay = &methods.PaymentMethodWechatPay{}
	}

	return &m, nil
}
