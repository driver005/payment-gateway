package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/driver005/gateway/btcpay"
	"github.com/driver005/gateway/helper"
	"github.com/julienschmidt/httprouter"
)

type BtcPay struct {
	r      Registry
	client *btcpay.Client
	ctx    context.Context
}

func NewBtcPay(r Registry) *BtcPay {
	return &BtcPay{
		r:      r,
		client: btcpay.NewClient("https://testnet.demo.btcpayserver.org", "d0b9689840047ed7b0737d9e981ab5e83c1b9d35"),
		ctx:    context.Background(),
	}
}

func (b *BtcPay) query(r *http.Request, name string) string {
	return r.FormValue(name)
}

func (n *BtcPay) state(r *http.Request, name string, defaultState bool) bool {
	if n.query(r, name) == "false" {
		return false
	} else if n.query(r, name) != "true" {
		return true
	} else {
		return defaultState
	}
}

// Basic operations

func (b *BtcPay) GetHealth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.GetHealth(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetServerInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.GetServerInfo(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetLanguageCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.GetLanguageCodes(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetInvoiceCheckoutPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	response, statusCode, err := b.client.GetInvoiceCheckoutPage(b.ctx, &invoiceId, b.query(r, "lang"))
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

// ApiKey

func (b *BtcPay) RevokeAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	apiKey := btcpay.APIKey(ps.ByName("apikey"))
	statusCode, err := b.client.RevokeAPIKey(b.ctx, &apiKey)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

func (b *BtcPay) GetCurrentAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.GetCurrentAPIKey(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) RevokeCurrentAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.RevokeCurrentAPIKey(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreateAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.APIKeyRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.CreateAPIKey(b.ctx, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

// Apps

// Authorize

func (b *BtcPay) Authorize(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.AuthorizationRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	statusCode, err := b.client.Authorize(b.ctx, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

// User

func (b *BtcPay) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.GetUser(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.CreateUser(b.ctx, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

// Store

func (b *BtcPay) GetStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, statusCode, err := b.client.GetStores(b.ctx)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.GetStore(b.ctx, &storeId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreateStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.StoreRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.CreateStore(b.ctx, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) UpdateStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	var m btcpay.StoreUpdate

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.UpdateStore(b.ctx, &storeId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) RemoveStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	statusCode, err := b.client.RemoveStore(b.ctx, &storeId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

// Payment Request

func (b *BtcPay) GetPaymentRequests(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.GetPaymentRequests(b.ctx, &storeId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetPaymentRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	paymentRequestId := btcpay.PaymentRequestID(ps.ByName("paymentRequestId"))
	response, statusCode, err := b.client.GetPaymentRequest(b.ctx, &storeId, &paymentRequestId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreatePaymentRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	var m btcpay.PaymentRequestRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.CreatePaymentRequest(b.ctx, &storeId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) ArchivePaymentRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	paymentRequestId := btcpay.PaymentRequestID(ps.ByName("paymentRequestId"))
	statusCode, err := b.client.ArchivePaymentRequest(b.ctx, &storeId, &paymentRequestId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

func (b *BtcPay) UpdatePaymentRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	paymentRequestId := btcpay.PaymentRequestID(ps.ByName("paymentRequestId"))
	var m btcpay.PaymentRequestRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.UpdatePaymentRequest(b.ctx, &storeId, &paymentRequestId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

// Pull payments

func (b *BtcPay) GetPullPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.GetPullPayments(b.ctx, &storeId, b.state(r, "includeArchived", false))
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetPullPayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pullPaymentId := btcpay.PullPaymentID(ps.ByName("pullPaymentId"))
	response, statusCode, err := b.client.GetPullPayment(b.ctx, &pullPaymentId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreatePullPayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	var m btcpay.PullPaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.CreatePullPayment(b.ctx, &storeId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) ArchivePullPayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	pullPaymentId := btcpay.PullPaymentID(ps.ByName("pullPaymentId"))
	statusCode, err := b.client.ArchivePullPayment(b.ctx, &storeId, &pullPaymentId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

func (b *BtcPay) GetPayouts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pullPaymentId := btcpay.PullPaymentID(ps.ByName("pullPaymentId"))

	response, statusCode, err := b.client.GetPayouts(b.ctx, &pullPaymentId, b.state(r, "includeCancelled", false))
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreatePayout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pullPaymentId := btcpay.PullPaymentID(ps.ByName("pullPaymentId"))
	var m btcpay.PayoutRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.CreatePayout(b.ctx, &pullPaymentId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) ApprovePayout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	payoutId := btcpay.PayoutID(ps.ByName("payoutId"))
	var m btcpay.PayoutApproveRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	response, statusCode, err := b.client.ApprovePayout(b.ctx, &storeId, &payoutId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CancelPayout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	payoutId := btcpay.PayoutID(ps.ByName("payoutId"))

	statusCode, err := b.client.CancelPayout(b.ctx, &storeId, &payoutId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

// Invoice

func (b *BtcPay) GetInvoices(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.GetInvoices(b.ctx, &storeId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) CreateInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.InvoiceRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.CreateInvoice(b.ctx, &storeId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.GetInvoice(b.ctx, &storeId, &invoiceId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) ArchiveInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	statusCode, err := b.client.ArchiveInvoice(b.ctx, &storeId, &invoiceId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, nil)
}

func (b *BtcPay) UpdateInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.InvoiceUpdate

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.UpdateInvoice(b.ctx, &storeId, &invoiceId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) GetInvoicePaymentMethod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.GetInvoicePaymentMethod(b.ctx, &storeId, &invoiceId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) MarkInvoiceStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m btcpay.MarkInvoiceStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		b.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.MarkInvoiceStatus(b.ctx, &storeId, &invoiceId, &m)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) UnarchiveInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	response, statusCode, err := b.client.UnarchiveInvoice(b.ctx, &storeId, &invoiceId)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}

func (b *BtcPay) ActivatePaymentMethod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoiceId := btcpay.InvoiceID(ps.ByName("invoiceId"))
	storeId := btcpay.StoreID(ps.ByName("storeId"))
	paymentMethod := ps.ByName("paymentMethod")
	response, statusCode, err := b.client.ActivatePaymentMethod(b.ctx, &storeId, &invoiceId, &paymentMethod)
	if err != nil {
		b.r.Writer().WriteErrorCode(w, r, statusCode, helper.WithStack(err))
		return
	}

	b.r.Writer().WriteCode(w, r, statusCode, response)
}
