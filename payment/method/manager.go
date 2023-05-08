package method

import (
	"github.com/driver005/gateway/payment/method/methods"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias PaymentMethod

func (h *Handler) Bind(context *fiber.Ctx) (*PaymentMethod, error) {
	var m PaymentMethod
	var err error

	type request struct {
		*Alias
		Customer                    uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		AccountNumber               string        `json:"account_number,omitempty" swaggertype:"primitive,string" format:"uuid"`
		AccountHolderType           string        `json:"account_holder_type,omitempty" swaggertype:"primitive,string" format:"uuid"`
		AccountType                 string        `json:"account_type,omitempty" swaggertype:"primitive,string" format:"uuid"`
		FinancialConnectionsAccount string        `json:"financial_connections_account,omitempty" swaggertype:"primitive,string" format:"uuid"`
		RoutingNumber               string        `json:"routing_number,omitempty" swaggertype:"primitive,string" format:"uuid"`
		InstitutionNumber           string        `json:"institution_number,omitempty" swaggertype:"primitive,string" format:"uuid"`
		TransitNumber               string        `json:"transit_number,omitempty" swaggertype:"primitive,string" format:"uuid"`
		BsbNumber                   string        `json:"bsb_number,omitempty" swaggertype:"primitive,string" format:"uuid"`
		SortCode                    string        `json:"sort_code,omitempty" swaggertype:"primitive,string" format:"uuid"`
		TaxId                       string        `json:"tax_id,omitempty" swaggertype:"primitive,string" format:"uuid"`
		ExpMonth                    int           `json:"exp_month,omitempty" swaggertype:"primitive,string" format:"uuid"`
		ExpYear                     int           `json:"exp_year,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Number                      string        `json:"number,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Cvc                         string        `json:"cvc,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Bank                        string        `json:"bank,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Day                         int           `json:"day,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Month                       int           `json:"month,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Year                        int           `json:"year,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Iban                        string        `json:"iban,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Country                     string        `json:"country,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context, model.Customer.UUID)
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
