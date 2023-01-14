package wechat

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestCloseRequestDo(t *testing.T) {
	client, err := mockNewClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("client is nil")
	}

	cases := []struct {
		req       *CloseRequest
		transport *mockTransport
		pass      bool
	}{
		{
			&CloseRequest{
				MchId:      client.config.MchId,
				OutTradeNo: "fortest",
			},
			nil,
			true,
		},
		{
			&CloseRequest{
				OutTradeNo: "fortest",
			},
			nil,
			true,
		},
		{
			&CloseRequest{
				MchId:      client.config.MchId,
				OutTradeNo: "fortest",
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
	}

	ctx := context.Background()
	for _, c := range cases {
		if c.transport != nil {
			client.config.opts.transport = c.transport
			client.secrets.clear()
		}

		err := c.req.Do(ctx, client)
		pass := err == nil
		if pass != c.pass {
			t.Fatalf("expect %v, got %v, err: %v", c.pass, pass, err)
		}

		if err != nil {
			continue
		}
	}
}
