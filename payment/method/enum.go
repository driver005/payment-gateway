package method

import "database/sql/driver"

type Type string

const (
	TypeAcssDebit        Type = "acss_debit"
	TypeAffirm           Type = "affirm"
	TypeAfterpayClearpay Type = "afterpay_clearpay"
	TypeAlipay           Type = "alipay"
	TypeAuBecsDebit      Type = "au_becs_debit"
	TypeBacsDebit        Type = "bacs_debit"
	TypeBancontact       Type = "bancontact"
	TypeBlik             Type = "blik"
	TypeBoleto           Type = "boleto"
	TypeBtcpay           Type = "btcpay"
	TypeCard             Type = "card"
	TypeCardPresent      Type = "card_present"
	TypeCustomerBalance  Type = "customer_balance"
	TypeEps              Type = "eps"
	TypeFpx              Type = "fpx"
	TypeGiropay          Type = "giropay"
	TypeGrabpay          Type = "grabpay"
	TypeIdeal            Type = "ideal"
	TypeInteracPresent   Type = "interac_present"
	TypeKlarna           Type = "klarna"
	TypeKonbini          Type = "konbini"
	TypeLink             Type = "link"
	TypeOxxo             Type = "oxxo"
	TypeP24              Type = "p24"
	TypePaynow           Type = "paynow"
	TypePix              Type = "pix"
	TypePromptpay        Type = "promptpay"
	TypeSepaDebit        Type = "sepa_debit"
	TypeSofort           Type = "sofort"
	TypeUsBankAccount    Type = "us_bank_account"
	TypeWechatPay        Type = "wechat_pay"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}
