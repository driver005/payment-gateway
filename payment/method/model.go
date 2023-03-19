package method

import (
	"time"

	"github.com/driver005/database"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/method/methods"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/radar"
	"github.com/google/uuid"
)

type PaymentMethod struct {
	core.Model

	AcssDebit        *methods.PaymentMethodAcssDebit        `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Affirm           *methods.PaymentMethodAffirm           `json:"affirm,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AfterpayClearpay *methods.PaymentMethodAfterpayClearpay `json:"afterpay_clearpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Alipay           *methods.PaymentMethodAlipay           `json:"alipay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AuBecsDebit      *methods.PaymentMethodAuBecsDebit      `json:"au_becs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BacsDebit        *methods.PaymentMethodBacsDebit        `json:"bacs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Bancontact       *methods.PaymentMethodBancontact       `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BillingDetails   *address.BillingDetails                `json:"billing_details,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Blik             *methods.PaymentMethodBlik             `json:"blik,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Boleto           *methods.PaymentMethodBoleto           `json:"boleto,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BtcPay           *methods.PaymentMethodBtcPay           `json:"btc_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card             *methods.PaymentMethodCard             `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CardPresent      *methods.PaymentMethodCardPresent      `json:"card_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CustomerBalance  *methods.PaymentMethodCustomerBalance  `json:"customer_balance,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Eps              *methods.PaymentMethodEps              `json:"eps,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Fpx              *methods.PaymentMethodFpx              `json:"fpx,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Giropay          *methods.PaymentMethodGiropay          `json:"giropay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Grabpay          *methods.PaymentMethodGrabpay          `json:"grabpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Ideal            *methods.PaymentMethodIdeal            `json:"ideal,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	InteracPresent   *methods.PaymentMethodInteracPresent   `json:"interac_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Klarna           *methods.PaymentMethodKlarna           `json:"klarna,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Konbini          *methods.PaymentMethodKonbini          `json:"konbini,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Link             *methods.PaymentMethodLink             `json:"link,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Oxxo             *methods.PaymentMethodOxxo             `json:"oxxo,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	P24              *methods.PaymentMethodP24              `json:"p24,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Paynow           *methods.PaymentMethodPaynow           `json:"paynow,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Pix              *methods.PaymentMethodPix              `json:"pix,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Promptpay        *methods.PaymentMethodPromptpay        `json:"promptpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	RadarOptions     *radar.RadarOptions                    `json:"radar_options,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SepaDebit        *methods.PaymentMethodSepaDebit        `json:"sepa_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Sofort           *methods.PaymentMethodSofort           `json:"sofort,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount    *methods.PaymentMethodUsBankAccount    `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	WechatPay        *methods.PaymentMethodWechatPay        `json:"wechat_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Type             Type                                   `json:"type"`

	CustomerId *uuid.UUID         `json:"customer_id" swaggerignore:"true"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`
}

func (m *PaymentMethod) BeforeCreate(tx *database.DB) (err error) {
	if m.Id == uuid.Nil {
		m.Id, err = uuid.NewUUID()
		if err != nil {
			return err
		}
	}

	if m.Object == "" {
		m.Object = tx.Statement.Schema.Table
	}

	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt

	return nil
}
