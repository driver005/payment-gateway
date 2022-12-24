package handler

import (
	"context"

	"github.com/driver005/gateway/crypto/btcpay"
	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
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

func (b *BtcPay) query(context *fiber.Ctx, name string) string {
	return context.Query(name)
}

func (n *BtcPay) state(context *fiber.Ctx, name string, defaultState bool) bool {
	if context.Query(name) == "false" {
		return false
	} else if context.Query(name) != "true" {
		return true
	} else {
		return defaultState
	}
}

// Basic operations

func (b *BtcPay) GetHealth(context *fiber.Ctx) error {
	response, statusCode, err := b.client.GetHealth(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetServerInfo(context *fiber.Ctx) error {
	response, statusCode, err := b.client.GetServerInfo(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetLanguageCodes(context *fiber.Ctx) error {
	response, statusCode, err := b.client.GetLanguageCodes(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetInvoiceCheckoutPage(context *fiber.Ctx) error {
	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	response, statusCode, err := b.client.GetInvoiceCheckoutPage(b.ctx, &invoiceId, b.query(context, "lang"))
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

// ApiKey

func (b *BtcPay) RevokeAPIKey(context *fiber.Ctx) error {
	apiKey := btcpay.APIKey(context.Params("apikey"))
	statusCode, err := b.client.RevokeAPIKey(b.ctx, &apiKey)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

func (b *BtcPay) GetCurrentAPIKey(context *fiber.Ctx) error {
	response, statusCode, err := b.client.GetCurrentAPIKey(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) RevokeCurrentAPIKey(context *fiber.Ctx) error {
	response, statusCode, err := b.client.RevokeCurrentAPIKey(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreateAPIKey(context *fiber.Ctx) error {
	var m btcpay.APIKeyRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.CreateAPIKey(b.ctx, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

// Apps

// Authorize

func (b *BtcPay) Authorize(context *fiber.Ctx) error {
	var m btcpay.AuthorizationRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	statusCode, err := b.client.Authorize(b.ctx, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

// User

func (b *BtcPay) GetUser(context *fiber.Ctx) error {
	response, statusCode, err := b.client.GetUser(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreateUser(context *fiber.Ctx) error {
	var m btcpay.UserRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.CreateUser(b.ctx, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

// Store

func (b *BtcPay) GetStores(context *fiber.Ctx) error {
	response, statusCode, err := b.client.GetStores(b.ctx)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetStore(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.GetStore(b.ctx, &storeId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreateStore(context *fiber.Ctx) error {
	var m btcpay.StoreRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.CreateStore(b.ctx, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) UpdateStore(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	var m btcpay.StoreUpdate

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.UpdateStore(b.ctx, &storeId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) RemoveStore(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	statusCode, err := b.client.RemoveStore(b.ctx, &storeId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

// Payment Request

func (b *BtcPay) GetPaymentRequests(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.GetPaymentRequests(b.ctx, &storeId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetPaymentRequest(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	paymentRequestId := btcpay.PaymentRequestID(context.Params("paymentRequestId"))
	response, statusCode, err := b.client.GetPaymentRequest(b.ctx, &storeId, &paymentRequestId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreatePaymentRequest(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	var m btcpay.PaymentRequestRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.CreatePaymentRequest(b.ctx, &storeId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) ArchivePaymentRequest(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	paymentRequestId := btcpay.PaymentRequestID(context.Params("paymentRequestId"))
	statusCode, err := b.client.ArchivePaymentRequest(b.ctx, &storeId, &paymentRequestId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

func (b *BtcPay) UpdatePaymentRequest(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	paymentRequestId := btcpay.PaymentRequestID(context.Params("paymentRequestId"))
	var m btcpay.PaymentRequestRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.UpdatePaymentRequest(b.ctx, &storeId, &paymentRequestId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

// Pull payments

func (b *BtcPay) GetPullPayments(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.GetPullPayments(b.ctx, &storeId, b.state(context, "includeArchived", false))
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetPullPayment(context *fiber.Ctx) error {
	pullPaymentId := btcpay.PullPaymentID(context.Params("pullPaymentId"))
	response, statusCode, err := b.client.GetPullPayment(b.ctx, &pullPaymentId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreatePullPayment(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	var m btcpay.PullPaymentRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.CreatePullPayment(b.ctx, &storeId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) ArchivePullPayment(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	pullPaymentId := btcpay.PullPaymentID(context.Params("pullPaymentId"))
	statusCode, err := b.client.ArchivePullPayment(b.ctx, &storeId, &pullPaymentId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

func (b *BtcPay) GetPayouts(context *fiber.Ctx) error {
	pullPaymentId := btcpay.PullPaymentID(context.Params("pullPaymentId"))

	response, statusCode, err := b.client.GetPayouts(b.ctx, &pullPaymentId, b.state(context, "includeCancelled", false))
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreatePayout(context *fiber.Ctx) error {
	pullPaymentId := btcpay.PullPaymentID(context.Params("pullPaymentId"))
	var m btcpay.PayoutRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.CreatePayout(b.ctx, &pullPaymentId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) ApprovePayout(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	payoutId := btcpay.PayoutID(context.Params("payoutId"))
	var m btcpay.PayoutApproveRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	response, statusCode, err := b.client.ApprovePayout(b.ctx, &storeId, &payoutId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CancelPayout(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	payoutId := btcpay.PayoutID(context.Params("payoutId"))

	statusCode, err := b.client.CancelPayout(b.ctx, &storeId, &payoutId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

// Invoice

func (b *BtcPay) GetInvoices(context *fiber.Ctx) error {
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.GetInvoices(b.ctx, &storeId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) CreateInvoice(context *fiber.Ctx) error {
	var m btcpay.InvoiceRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.CreateInvoice(b.ctx, &storeId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetInvoice(context *fiber.Ctx) error {
	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.GetInvoice(b.ctx, &storeId, &invoiceId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) ArchiveInvoice(context *fiber.Ctx) error {
	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	statusCode, err := b.client.ArchiveInvoice(b.ctx, &storeId, &invoiceId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(nil)
}

func (b *BtcPay) UpdateInvoice(context *fiber.Ctx) error {
	var m btcpay.InvoiceUpdate

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.UpdateInvoice(b.ctx, &storeId, &invoiceId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) GetInvoicePaymentMethod(context *fiber.Ctx) error {
	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.GetInvoicePaymentMethod(b.ctx, &storeId, &invoiceId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) MarkInvoiceStatus(context *fiber.Ctx) error {
	var m btcpay.MarkInvoiceStatusRequest

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.MarkInvoiceStatus(b.ctx, &storeId, &invoiceId, &m)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) UnarchiveInvoice(context *fiber.Ctx) error {
	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	response, statusCode, err := b.client.UnarchiveInvoice(b.ctx, &storeId, &invoiceId)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}

func (b *BtcPay) ActivatePaymentMethod(context *fiber.Ctx) error {
	invoiceId := btcpay.InvoiceID(context.Params("invoiceId"))
	storeId := btcpay.StoreID(context.Params("storeId"))
	paymentMethod := context.Params("paymentMethod")
	response, statusCode, err := b.client.ActivatePaymentMethod(b.ctx, &storeId, &invoiceId, &paymentMethod)
	if err != nil {
		return context.Status(statusCode).JSON(helper.WithStack(err))
	}

	return context.Status(statusCode).JSON(response)
}
