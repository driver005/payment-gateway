package alipay_test

import (
	"testing"

	"github.com/driver005/gateway/driver/wallet/alipay"
)

func TestClient_TradeWapPay(t *testing.T) {
	t.Log("========== TradeWapPay ==========")
	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.ReturnURL = "http://203.86.24.181:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_20170623021124"
	p.TotalAmount = "10.00"
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}