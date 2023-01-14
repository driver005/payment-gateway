package driver

import (
	"context"
	"encoding/json"
	"time"

	"github.com/driver005/gateway/driver/wallet/alipay"
	"github.com/driver005/gateway/driver/wallet/applepay"
	"github.com/driver005/gateway/driver/wallet/btcpay"
	"github.com/driver005/gateway/driver/wallet/grabpay/dto"
	"github.com/driver005/gateway/driver/wallet/grabpay/merchant"
	"github.com/driver005/gateway/driver/wallet/grabpay/utils"
	"github.com/driver005/gateway/driver/wallet/wechat"
)

func (h *Handler) Alipay() (string, error) {
	client, err := alipay.New("SANDBOX_5Y8V3J2YUS1100350", "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB", true)
	if err != nil {
		return "", err
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		return "", err
	}

	var payURL = url.String()

	return payURL, nil
}

func (h *Handler) ApplePay() (*applepay.Token, error) {
	ap, err := applepay.New(
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
		return nil, err
	}

	payload, err := ap.Session("")
	if err != nil {
		return nil, err
	}

	response := applepay.Response{}
	err = json.Unmarshal(payload, &response)
	if err != nil {
		return nil, err
	}

	token, err := ap.DecryptResponse(&response)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (h *Handler) BtcPay() (string, error) {
	client := btcpay.NewClient("https://testnet.demo.btcpayserver.org", "d0b9689840047ed7b0737d9e981ab5e83c1b9d35")

	invoice, _, err := client.CreateInvoice(context.TODO(), &client.Store.ID, &btcpay.InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
	})

	if err != nil {
		return "", err
	}

	return invoice.CheckoutLink, nil
}

func (h *Handler) GrabPay() (string, error) {
	merchant := merchant.NewMerchantOnline("STG", "DE", "partnerID", "partnerSecret", "merchantID", "clientID", "clientSecret", "redirectUri")

	params := new(dto.OnaCreateWebUrlParams)
	params.PartnerTxID = utils.GenerateRandomString(32)
	params.PartnerGroupTxID = utils.GenerateRandomString(32)
	params.Amount = int64(2000)
	params.Description = "testing"
	params.Currency = "VND"

	// Web init
	response, _, err := merchant.OnaCreateWebUrl(context.Background(), params)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (h *Handler) WeChat() (string, error) {
	payClient, err := wechat.NewClient(wechat.Config{
		AppId:       "appId",
		MchId:       "mchId",
		Apiv3Secret: "apiv3Secret",
		Cert: wechat.CertSuite{
			SerialNo:       "serialNo",
			PrivateKeyPath: "privateKeyPath",
		},
	})
	if err != nil {
		return "", err
	}

	// create a pay request
	req := &wechat.PayRequest{
		Description: "for testing",
		TimeExpire:  time.Now().Add(10 * time.Minute),
		Amount: wechat.PayAmount{
			Total:    int(10 * 100),
			Currency: "CNY",
		},
		TradeType: wechat.Native,
	}

	resp, err := req.Do(context.Background(), payClient)
	if err != nil {
		return "", err
	}

	return resp.CodeUrl, nil
}
