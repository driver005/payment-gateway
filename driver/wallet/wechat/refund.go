package wechat

import (
	"context"
	"errors"
	"net/http"
	"time"
)

// RefundRequest is request when apply refund, TransactionId
// and OutTradeNo is required.
type RefundRequest struct {
	TransactionId string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	OutRefundNo   string `json:"out_refund_no"`
	Reason        string `json:"reason,omitempty"`
	NotifyUrl     string `json:"notify_url,omitempty"`
	FundsAccount  string `json:"funds_account,omitempty"`

	Amount      RefundAmount       `json:"amount"`
	GoodsDetail []RefundGoodDetail `json:"goods_detail,omitempty"`
}

// RefundAmount is total amount refund, have total and currency.
type RefundAmount struct {
	Refund   int    `json:"refund"`
	Total    int    `json:"tota"`
	Currency string `json:"currency"`
}

// RefundGoodDetail is the good information about refund transaction.
type RefundGoodDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"`
	GoodsName        string `json:"goods_name,omitempty"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}

// RefundResponse is the response for refund transaction.
type RefundResponse struct {
	RefundId            string    `json:"refund_id"`
	OutRefundNo         string    `json:"out_refund_no"`
	TransactionId       string    `json:"transaction_id"`
	OutTradeNo          string    `json:"out_trade_no"`
	Channel             string    `json:"channe"`
	UserReceivedAccount string    `json:"user_received_account"`
	SuccessTime         time.Time `json:"success_time,omitempty"`
	CreateTime          time.Time `json:"create_time"`
	Status              string    `json:"status"`
	FundsAccount        string    `json:"funds_account,omitempty"`

	Amount    RefundAmountInQueryResp  `json:"amount"`
	Promotion []*RefundPromotionDetail `json:"promotion_detail,omitempty"`
}

// RefundAmountInQueryResp is total amount refund.
type RefundAmountInQueryResp struct {
	Total            int    `json:"tota"`
	Refund           int    `json:"refund"`
	PayerTotal       int    `json:"payer_tota"`
	PayerRefund      int    `json:"payer_refund"`
	SettlementTotal  int    `json:"settlement_tota"`
	SettlementRefund int    `json:"settlement_refund"`
	DiscountRefund   int    `json:"discount_refund"`
	Currency         string `json:"currency"`
}

// RefundPromotionDetail is the promotion information about refund transaction.
type RefundPromotionDetail struct {
	PromotionId  int    `json:"promotion_id"`
	Scope        string `json:"scope"`
	Type         string `json:"type"`
	Amount       int    `json:"amount"`
	RefundAmount int    `json:"refund_amount"`

	GoodsDetail []RefundGoodDetail `json:"goods_detail,omitempty"`
}

// Do send the refund request and return refund response.
func (r *RefundRequest) Do(ctx context.Context, c Client) (*RefundResponse, error) {
	url := r.url(c.Config().Options().Domain)

	if err := r.validate(); err != nil {
		return nil, err
	}

	resp := &RefundResponse{}
	if err := c.Do(ctx, http.MethodPost, url, r).Scan(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RefundRequest) validate() error {
	if r.TransactionId == "" {
		return errors.New("transaction_id can't be empty")
	}
	if r.OutRefundNo == "" {
		return errors.New("out_refund_no can't be empty")
	}
	if r.OutTradeNo == "" {
		return errors.New("out_trade_no can't be empty")
	}
	if r.Amount.Refund <= 0 {
		return errors.New("refund can't less than 0")
	}
	if r.Amount.Total <= 0 {
		return errors.New("total can't less than 0")
	}
	if r.Amount.Currency == "" {
		return errors.New("currency can't be empty")
	}

	return nil
}

func (r *RefundRequest) url(domain string) string {
	return domain + `/v3/refund/domestic/refunds`
}
