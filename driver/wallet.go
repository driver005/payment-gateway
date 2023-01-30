package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/driver005/gateway/driver/wallet/alipay"
	"github.com/driver005/gateway/driver/wallet/applepay"
	"github.com/driver005/gateway/driver/wallet/btcpay"
	"github.com/driver005/gateway/driver/wallet/grabpay/dto"
	"github.com/driver005/gateway/driver/wallet/grabpay/merchant"
	"github.com/driver005/gateway/driver/wallet/grabpay/utils"
	"github.com/driver005/gateway/driver/wallet/wechat"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/intent/next"
)

//
// Alipay
//

func InitAlipay() *alipay.Client {
	privatKey := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCzNgtbsfADVychhO/C8q17O0AEN8vIEtPIxDHY0oVe77k6NEVf7PmAYwuez2qYQtE2YtNH/+sIuMb5bPwPw7mEWupQ8L+xIWojY5HIcti63Rgca+Fkp6FScM3jEj2dmkmruajc7PsPEm9C/ePq34gC2YfIZpZ92WYR2BKh6G2zfFsv7KjL9SPXzA04IJzMx9zNlwC5F83lDgaH6x9xdokT1zpNNVrluwPTsLKw7dTU+Fa7rvz7P/3ry1bWW2nZwq90rp5oon+70/u64OPccLek0i4A2q4UxmN9dzdIEL+BI5UeW6vkLFTJEoElBVvaqAdvvWAJ53gmdii3DimHePZfAgMBAAECggEAAtrVi3xmKmd1BB+uAhkDknb7KhYJdQJ3xTb65E1WOpDvDnJOuJx5zQXG5ZzAQfSuiINCgl8dS2JGKSsHSOkBR457NRbgtt2lZzLq25qYwUFMZQeC66uP8lRBs2BTAvO2QLoQapSL202b4CFTRCaTcBgvBY48z3xrF9WtVk7xkG5f1SkJhXOp0KuogfFKnI5kpHp3BomaDPddGoUVnMLewLeljakSWxK8pxOuf6WjElFNiQ/dUBf3GuHWZbcvB7l9/Uw6NpJDsRZoDYjCFGL3CrYaF3P8wKq2hl5YXF0bHdVSkRCKCNxj1JX+plofW0w8hQACzzScp2o1xP3DeIOn4QKBgQDyWE9RUFo/FXGscTB9nfpoZx/HAfQaXASvbP68pF1bQkuwqu9cbQRJfIqbdMKTgQK2+cZ0CCR6T6ppkQAA7ifnN3jftRqZAFm3xk9woTJzQDcHPr8Wh3Po/4j8HtVT1h03UwY/zSQyfd4GBrLw6xGQlj9jyC9PuDj/6r8U4YezKwKBgQC9TxBsc3fDVCMpdRY9Isk8xNhzVtzvt8UnkBaHnuFj3bUoNbysC5KBAbvF/Vycjrfd85aK6+2cad/goSsqagFZs1ks2aqoHgc6KVW2oP6nDAy7dw3PTLe0GLC/5Dak0SuKU+PdhkDac/PuxdtRFfC/IPatACUQ4kHF7cWmPDm/nQKBgQCrLxAQGxc0Wlxqd6XOBcp89uMuesXH3Nn+ZJpn/B5puSGoEIZ1nhgaJJvwBYDwGoAlabfx/FJwB1gTq7X58kBDhh6evtmj2+iD9NPmdich2+lC9+KY8mNA4UgM92avCp6mlsHiTXDVLa2oShEaNqG4pTjvrdclk7bHpzUEde6Z8wKBgFrbfYhvisVmtlGJbpj8/xR5bWE+CV/MFYW9c/K3YHmryOx5jgoMq83a5SJLTY2eOcT+yfv2692sOtT2xV7f7bH9kAkklvGFJ/bUK90xO8c000+N/kTGD00SpJIUl91WvszvImgJrUqraCHrOC2cmCVDTuENNrz5upOiYMA+G4FJAoGABAlYxcUfOxBSVMxw7swTDPsv/MP5bmUTJAmTK8xUTzIfQfb+F5ODq5wWx+bgDUc6yqWnnPF4DukvUlvO4AoM0ir5zeYs+bdu+eqcEZqlZRuQ7jF3aXmWzOUZyVBEnVYcEkm1wtgiHhYQVCddmROvenAwTfg2XXM+QWPMkRb12A8="
	client, err := alipay.New("5Y8V3J2YUS1100350", privatKey, true)
	if err != nil {
		fmt.Println(err)
	}

	if err = client.LoadAppPublicCertFromFile("cert/appPublicCert.crt"); err != nil {
		log.Println("error: ", err)
	}
	if err = client.LoadAliPayRootCertFromFile("cert/alipayRootCert.crt"); err != nil {
		log.Println("error: ", err)
	}
	if err = client.LoadAliPayPublicCertFromFile("cert/alipayPublicCert.crt"); err != nil {
		log.Println("error: ", err)
	}

	return client
}

func (h *Handler) Alipay(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://localhost/api/v1/notify"
	p.ReturnURL = m.NextAction.AlipayHandleRedirect.ReturnUrl
	p.Subject = m.Description
	p.OutTradeNo = fmt.Sprintf("trade_no_%+v", m.Id)
	p.TotalAmount = fmt.Sprint(m.Amount)
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := h.alipay.TradePagePay(p)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "alipay_handle_redirect",
		AlipayHandleRedirect: &next.PaymentIntentNextActionAlipayHandleRedirect{

			Url: url.String(),
		},
	}

	return m, nil
}

//
// ApplePay
//

func InitApplePay() *applepay.Merchant {
	client, err := applepay.New(
		"merchant.com.processout.test",
		applepay.MerchantDisplayName("ProcessOut Development Store"),
		applepay.MerchantDomainName("applepay.processout.com"),
		applepay.MerchantCertificateLocation(
			"certs/cert-merchant.crt",
			"certs/cert-merchant-key.pem",
		),
		applepay.ProcessingCertificateLocation(
			"certs/cert-processing.crt",
			"certs/cert-processing-key.pem",
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	return client
}

func (h *Handler) ApplePay() (*applepay.Token, error) {
	payload, err := h.applepay.Session("")
	if err != nil {
		return nil, err
	}

	response := applepay.Response{}
	err = json.Unmarshal(payload, &response)
	if err != nil {
		return nil, err
	}

	token, err := h.applepay.DecryptResponse(&response)
	if err != nil {
		return nil, err
	}

	return token, nil
}

//
// BtcPay
//

func InitBtcPay() *btcpay.Client {
	return btcpay.NewClient("https://testnet.demo.btcpayserver.org", "d0b9689840047ed7b0737d9e981ab5e83c1b9d35")
}

func (h *Handler) BtcPay(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	invoice, _, err := h.btcpay.CreateInvoice(context.TODO(), &h.btcpay.Store.ID, &btcpay.InvoiceRequest{
		Amount:   fmt.Sprint(m.Amount),
		Currency: m.Currency,
	})

	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{

			Url: invoice.CheckoutLink,
		},
	}

	return m, nil
}

//
// GrabPay
//

func InitGrabPay() *merchant.OnlineTransaction {
	ot := merchant.NewMerchantOnline("STG", "DE", "partnerID", "partnerSecret", "merchantID", "clientID", "clientSecret", "redirectUri")

	return &ot
}

func (h *Handler) GrabPay(m *intent.PaymentIntent) (string, error) {
	params := new(dto.OnaCreateWebUrlParams)
	params.PartnerTxID = utils.GenerateRandomString(32)
	params.PartnerGroupTxID = utils.GenerateRandomString(32)
	params.Amount = int64(m.Amount)
	params.Description = m.Description
	params.Currency = m.Currency

	// Web init
	response, _, err := h.grabpay.OnaCreateWebUrl(context.Background(), params)
	if err != nil {
		return "", err
	}

	return response, nil
}

//
// WeChat
//

func InitWeChat() *wechat.Client {
	client, err := wechat.NewClient(wechat.Config{
		AppId:       "appId",
		MchId:       "mchId",
		Apiv3Secret: "apiv3Secret",
		Cert: wechat.CertSuite{
			SerialNo:       "serialNo",
			PrivateKeyPath: "privateKeyPath",
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	return &client
}

func (h *Handler) WeChat(m *intent.PaymentIntent) (string, error) {
	// create a pay request
	req := &wechat.PayRequest{
		Description: m.Description,
		TimeExpire:  time.Now().Add(10 * time.Minute),
		Amount: wechat.PayAmount{
			Total:    m.Amount,
			Currency: m.Currency,
		},
		TradeType: wechat.Native,
	}

	resp, err := req.Do(context.Background(), h.wechat)
	if err != nil {
		return "", err
	}

	return resp.CodeUrl, nil
}
