package wechat

import (
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestDoForPay(t *testing.T) {
	client, err := mockNewClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("client is nil")
	}

	cases := []struct {
		req       *PayRequest
		resp      *PayResponse
		transport *mockTransport
		pass      bool
	}{
		{
			&PayRequest{
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
				TradeType: Native,
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			nil,
			true,
		},
		{
			&PayRequest{
				AppId:       client.config.AppId,
				MchId:       client.config.MchId,
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
				Payer:     &Payer{"openid"},
				TradeType: JSAPI,
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			nil,
			true,
		},
		{
			&PayRequest{
				AppId:       client.config.AppId,
				MchId:       client.config.MchId,
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			nil,
			true,
		},
		{
			&PayRequest{
				AppId:       client.config.AppId,
				MchId:       client.config.MchId,
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			&mockTransport{
				RoundTripFn: func(req *http.Request) (*http.Response, error) {
					var resp = &http.Response{
						StatusCode: http.StatusOK,
					}

					resp.Header = http.Header{}
					resp.Body = ioutil.NopCloser(strings.NewReader("{}"))
					return resp, nil
				},
			},
			false,
		},
		{
			&PayRequest{
				AppId:       client.config.AppId,
				MchId:       client.config.MchId,
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
				Payer: &Payer{},
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			nil,
			false,
		},
		{
			&PayRequest{
				AppId:       client.config.AppId,
				MchId:       client.config.MchId,
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
				Payer:     &Payer{},
				TradeType: JSAPI,
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			nil,
			false,
		},
		{
			&PayRequest{
				AppId:       client.config.AppId,
				MchId:       client.config.MchId,
				Description: "for testing",
				OutTradeNo:  "forxxxxxxxxx",
				TimeExpire:  time.Now().Add(10 * time.Minute),
				Attach:      "cipher code",
				NotifyUrl:   "https://luoji.live/notify",
				Amount: PayAmount{
					Total:    1,
					Currency: "CNY",
				},
				TradeType: JSAPI,
			},
			&PayResponse{
				CodeUrl: "weixin://wxpay/bizpayurl/up?pr=NwY5Mz9&groupid=00",
			},
			nil,
			false,
		},
	}

	ctx := context.Background()
	for _, c := range cases {
		if c.transport != nil {
			client.config.opts.transport = c.transport
			client.secrets.clear()
		}

		resp, err := c.req.Do(ctx, client)
		pass := err == nil
		if pass != c.pass {
			t.Fatalf("expect %v, got %v, err: %v", c.pass, pass, err)
		}

		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.resp, resp) {
			t.Fatalf("expect %v, got %v", c.resp, resp)
		}
	}
}
