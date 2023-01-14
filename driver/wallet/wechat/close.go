package wechat

import (
	"context"
	"net/http"
)

// CloseRequest is the request for close transaction.
type CloseRequest struct {
	MchId      string `json:"mchid"`
	OutTradeNo string `json:"-"`
}

// Do send the request of close transaction.
func (r *CloseRequest) Do(ctx context.Context, c Client) error {
	if r.MchId == "" {
		r.MchId = c.Config().MchId
	}

	url := r.url(c.Config().Options().Domain)

	if err := c.Do(ctx, http.MethodPost, url, r).Error(); err != nil {
		return err
	}

	return nil
}

// return the url for close transcation
func (r *CloseRequest) url(domain string) string {
	return domain + "/v3/pay/transactions/out-trade-no/" + r.OutTradeNo + "/close"
}
