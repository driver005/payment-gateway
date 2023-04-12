package intent

import (
	"context"
	"errors"

	"github.com/driver005/gateway/internal/intent/methods"
	"github.com/driver005/gateway/internal/intent/next"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s pq.StringArray, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (h *Handler) Bind(context *fiber.Ctx, isPayable bool) (*PaymentIntent, error) {
	type Alias PaymentIntent
	var m PaymentIntent
	var err error
	var paymentMethodOptions = &methods.PaymentIntentPaymentMethodOptions{}

	model := struct {
		*Alias
		Customer      uuid.NullUUID `json:"customer,omitempty gorm:"-:all"`
		PaymentMethod uuid.NullUUID `json:"payment_method,omitempty gorm:"-:all"`
		Confirm       bool          `json:"confirm,omitempty gorm:"-:all"`
		ReturnUrl     string        `json:"return_url,omitempty gorm:"-:all"`
		OffSession    bool          `json:"off_session,omitempty gorm:"-:all"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context.Context(), model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}

	if model.ReturnUrl != "" {
		m.NextAction = &next.PaymentIntentNextAction{
			RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
				ReturnUrl: model.ReturnUrl,
			},
		}
	}
	if model.OffSession {
		m.SetupFutureUsage = "off_session"
	}

	if contains(m.PaymentMethodTypes, "acss_debit") {
		paymentMethodOptions.AcssDebit = &methods.PaymentIntentPaymentMethodOptionsAcssDebit1{}
	}
	if contains(m.PaymentMethodTypes, "affirm") {
		paymentMethodOptions.Affirm = &methods.PaymentIntentPaymentMethodOptionsAffirm{}
	}
	if contains(m.PaymentMethodTypes, "afterpay_clearpay") {
		paymentMethodOptions.AfterpayClearpay = &methods.PaymentIntentPaymentMethodOptionsAfterpayClearpay{}
	}
	if contains(m.PaymentMethodTypes, "alipay") {
		paymentMethodOptions.Alipay = &methods.PaymentIntentPaymentMethodOptionsAlipay{}
	}
	if contains(m.PaymentMethodTypes, "au_becs_debit") {
		paymentMethodOptions.AuBecsDebit = &methods.PaymentIntentPaymentMethodOptionsAuBecsDebit1{}
	}
	if contains(m.PaymentMethodTypes, "bacs_debit") {
		paymentMethodOptions.BacsDebit = &methods.PaymentIntentPaymentMethodOptionsBacsDebit{}
	}
	if contains(m.PaymentMethodTypes, "bancontact") {
		paymentMethodOptions.Bancontact = &methods.PaymentIntentPaymentMethodOptionsBancontact{}
	}
	if contains(m.PaymentMethodTypes, "blik") {
		paymentMethodOptions.Blik = &methods.PaymentIntentPaymentMethodOptionsBlik{}
	}
	if contains(m.PaymentMethodTypes, "boleto") {
		paymentMethodOptions.Boleto = &methods.PaymentIntentPaymentMethodOptionsBoleto{}
	}
	if contains(m.PaymentMethodTypes, "btcpay") {
		paymentMethodOptions.BtcPay = &methods.PaymentIntentPaymentMethodOptionsBtcPay{}
	}
	if contains(m.PaymentMethodTypes, "card") {
		paymentMethodOptions.Card = &methods.PaymentIntentPaymentMethodOptionsCard1{}
	}
	if contains(m.PaymentMethodTypes, "card_present") {
		paymentMethodOptions.CardPresent = &methods.PaymentIntentPaymentMethodOptionsCardPresent{}
	}
	if contains(m.PaymentMethodTypes, "customer_balance") {
		paymentMethodOptions.CustomerBalance = &methods.PaymentIntentPaymentMethodOptionsCustomerBalance{}
	}
	if contains(m.PaymentMethodTypes, "eps") {
		paymentMethodOptions.Eps = &methods.PaymentIntentPaymentMethodOptionsEps1{}
	}
	if contains(m.PaymentMethodTypes, "fpx") {
		paymentMethodOptions.Fpx = &methods.PaymentIntentPaymentMethodOptionsFpx{}
	}
	if contains(m.PaymentMethodTypes, "giropay") {
		paymentMethodOptions.Giropay = &methods.PaymentIntentPaymentMethodOptionsGiropay{}
	}
	if contains(m.PaymentMethodTypes, "grabpay") {
		paymentMethodOptions.Grabpay = &methods.PaymentIntentPaymentMethodOptionsGrabpay{}
	}
	if contains(m.PaymentMethodTypes, "ideal") {
		paymentMethodOptions.Ideal = &methods.PaymentIntentPaymentMethodOptionsIdeal{}
	}
	if contains(m.PaymentMethodTypes, "interac_present") {
		paymentMethodOptions.InteracPresent = &methods.PaymentIntentPaymentMethodOptionsInteracPresent{}
	}
	if contains(m.PaymentMethodTypes, "klarna") {
		paymentMethodOptions.Klarna = &methods.PaymentIntentPaymentMethodOptionsKlarna{}
	}
	if contains(m.PaymentMethodTypes, "konbini") {
		paymentMethodOptions.Konbini = &methods.PaymentIntentPaymentMethodOptionsKonbini{}
	}
	if contains(m.PaymentMethodTypes, "link") {
		paymentMethodOptions.Link = &methods.PaymentIntentPaymentMethodOptionsLink1{}
	}
	if contains(m.PaymentMethodTypes, "oxxo") {
		paymentMethodOptions.Oxxo = &methods.PaymentIntentPaymentMethodOptionsOxxo{}
	}
	if contains(m.PaymentMethodTypes, "p24") {
		paymentMethodOptions.P24 = &methods.PaymentIntentPaymentMethodOptionsP24{}
	}
	if contains(m.PaymentMethodTypes, "paynow") {
		paymentMethodOptions.Paynow = &methods.PaymentIntentPaymentMethodOptionsPaynow{}
	}
	if contains(m.PaymentMethodTypes, "pix") {
		paymentMethodOptions.Pix = &methods.PaymentIntentPaymentMethodOptionsPix{}
	}
	if contains(m.PaymentMethodTypes, "promptpay") {
		paymentMethodOptions.Promptpay = &methods.PaymentIntentPaymentMethodOptionsPromptpay{}
	}
	if contains(m.PaymentMethodTypes, "sepa_debit") {
		paymentMethodOptions.SepaDebit = &methods.PaymentIntentPaymentMethodOptionsSepaDebit1{}
	}
	if contains(m.PaymentMethodTypes, "sofort") {
		paymentMethodOptions.Sofort = &methods.PaymentIntentPaymentMethodOptionsSofort{}
	}
	if contains(m.PaymentMethodTypes, "us_bank_account") {
		paymentMethodOptions.UsBankAccount = &methods.PaymentIntentPaymentMethodOptionsUsBankAccount1{}
	}
	if contains(m.PaymentMethodTypes, "wechat_pay") {
		paymentMethodOptions.WechatPay = &methods.PaymentIntentPaymentMethodOptionsWechatPay{}
	}

	m.PaymentMethodOptions = paymentMethodOptions

	if m.CaptureMethod == "" {
		m.CaptureMethod = CaptureMethodAutomatic
	}

	if m.ConfirmationMethod == "" {
		m.ConfirmationMethod = ConfirmationMethodAutomatic
	}

	if m.Status == StatusRequiresPaymentMethod && m.PaymentMethod != nil {
		m.Status = StatusRequiresAction
	}

	if m.Status == "" {
		if m.PaymentMethod == nil {
			m.Status = StatusRequiresPaymentMethod
		} else {
			m.Status = StatusRequiresAction
		}
	}

	if isPayable {
		if model.Confirm && m.Status != StatusRequiresPaymentMethod {
			r, err := h.Pay(context.Context(), m)
			if err != nil {
				return nil, err
			}

			m = *r
		} else {
			if model.Confirm && m.Status == StatusRequiresPaymentMethod {
				return nil, errors.New("to confirm payment method is required")
			}
		}
	}

	r, err := h.Upsert(context.Context(), &m)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (h *Handler) Pay(ctx context.Context, m PaymentIntent) (*PaymentIntent, error) {
	r, err := h.r.Pay(ctx, &m)
	if err != nil {
		return nil, err
	}

	return r, nil
}
