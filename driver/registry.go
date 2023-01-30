package driver

import (
	"github.com/driver005/gateway/driver/card"
	"github.com/driver005/gateway/driver/paylater/affirm"
	"github.com/driver005/gateway/driver/paylater/klarna"
	"github.com/driver005/gateway/driver/paylater/riverty"
	"github.com/driver005/gateway/driver/redirect/bancontact"
	"github.com/driver005/gateway/driver/redirect/blik"
	"github.com/driver005/gateway/driver/redirect/eps"
	"github.com/driver005/gateway/driver/redirect/fpx"
	"github.com/driver005/gateway/driver/redirect/giropay"
	"github.com/driver005/gateway/driver/redirect/ideal"
	"github.com/driver005/gateway/driver/voucher/oxxo"
	"github.com/driver005/gateway/driver/wallet/alipay"
	"github.com/driver005/gateway/driver/wallet/applepay"
	"github.com/driver005/gateway/driver/wallet/btcpay"
	"github.com/driver005/gateway/driver/wallet/grabpay/merchant"
	"github.com/driver005/gateway/driver/wallet/wechat"
	"github.com/driver005/gateway/internal/intent"
)

type Handler struct {
	r          Registry
	affirm     affirm.Affirm
	riverty    riverty.Riverty
	klarna     klarna.PaymentSrv
	card       card.Card
	alipay     alipay.Client
	applepay   applepay.Merchant
	btcpay     btcpay.Client
	grabpay    merchant.OnlineTransaction
	wechat     wechat.Client
	oxxo       oxxo.OpClient
	giropay    giropay.GiroPay
	bancontact bancontact.Bankcontact
	blik       blik.Blik
	eps        eps.Eps
	fpx        fpx.Fpx
	ideal      ideal.Ideal
}

type Registry interface {
	PaymentIntent() *intent.Handler
}

func NewHandler(r Registry) *Handler {
	return &Handler{
		r:       r,
		affirm:  *InitAffirm(),
		riverty: *InitRiverty(),
		klarna:  *InitKlarna(),
		card:    *InitCard(),
		alipay:  *InitAlipay(),
		// applepay: *InitApplePay(),
		btcpay:     *InitBtcPay(),
		grabpay:    *InitGrabPay(),
		wechat:     *InitWeChat(),
		oxxo:       *InitOxxo(),
		giropay:    *InitGiroPay(),
		bancontact: *InitBancontact(),
		blik:       *InitBlik(),
		eps:        *InitEps(),
		fpx:        *InitFpx(),
		ideal:      *InitIdeal(),
	}
}
