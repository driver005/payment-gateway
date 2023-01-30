package handler

import (
	"github.com/gofiber/fiber/v2"
)

const (
	NbxplorerPath = "/nbxplorer"
	BtcpayPath    = "/btcpay"
)

func (h *Handler) Routes(public fiber.Router) {
	h.NbxplorerRoutes(public)
	h.BtcpayRoutes(public)
	// h.ApplePayRoutes(public)
	// h.GooglePayRoutes(public)
}

func (h *Handler) NbxplorerRoutes(public fiber.Router) {
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/status", h.n.GetStatus)
	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/rescan", h.n.RescanTransactions)
	// public.Post(NbxplorerPath+"/cryptos/:cryptoCode/rpc", h.n.Rpc)
	// public.Post(NbxplorerPath+"/cryptos/:cryptoCode/connect", h.n.GetConnect)

	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations", h.n.CreateWallet)
	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme", h.n.TrackDerivationScheme)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/transactions", h.n.GetDerivationSchemeTransactions)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/transactions/:txId", h.n.GetDerivationSchemeTransaction)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/balance", h.n.GetCurrentBalance)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/addresses/unused", h.n.NewUnusedAddress)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/scripts/:script", h.n.GetScriptPubKeyInfos)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/utxos", h.n.GetDerivationSchemeUTXOs)
	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/utxos/scan", h.n.ScanUTXOSet)
	// public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/utxos/wipe", h.n.WipeUTXOSet)
	public.Put(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/metadata/:key", h.n.AttachDerivationSchemeMetadata)
	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/metadata/:key", h.n.DetachDerivationSchemeMetadata)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/metadata/:key", h.n.GetDerivationSchemeMetadata)
	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/prune", h.n.PruneUTXOSet)

	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address", h.n.TrackAddress)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address/transactions", h.n.GetAddressTransactions)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address/transactions/:txId", h.n.GetAddressTransaction)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address/utxos", h.n.GetAddressUTXOs)

	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/transactions", h.n.BroadcastTransaction)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/transactions/:txId", h.n.GetTransaction)

	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/fees/:blockCount", h.n.GetFeeRate)

	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/events", h.n.GetEventStream)
	public.Get(NbxplorerPath+"/cryptos/:cryptoCode/events/latest", h.n.GetRecentEventStream)

	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/psbt/create", h.n.CreatePSBT)
	public.Post(NbxplorerPath+"/cryptos/:cryptoCode/psbt/update", h.n.UpdatePSBT)
}

func (h *Handler) BtcpayRoutes(public fiber.Router) {
	public.Get(BtcpayPath+"/health", h.b.GetHealth)
	public.Get(BtcpayPath+"/server/info", h.b.GetServerInfo)
	public.Get(BtcpayPath+"/misc/lang", h.b.GetLanguageCodes)
	public.Get(BtcpayPath+"/i/:invoiceId", h.b.GetInvoiceCheckoutPage)

	public.Delete(BtcpayPath+"/api-keys/:apikey", h.b.RevokeAPIKey)
	public.Get(BtcpayPath+"/api-keys/current", h.b.GetCurrentAPIKey)
	// public.Delete(BtcpayPath+"/api-keys/current", h.b.RevokeCurrentAPIKey)
	public.Post(BtcpayPath+"/api-keys", h.b.CreateAPIKey)

	public.Get(BtcpayPath+"/api-keys/authorize", h.b.Authorize)

	// public.Post(BtcpayPath+"api/v1/stores/:storeId/apps/pos", h.b.CreateAPIKey)
	// public.Get(BtcpayPath+"api/v1/apps/:appId", h.b.CreateAPIKey)
	// public.Delete(BtcpayPath+"api/v1/apps/:appId", h.b.CreateAPIKey)

	public.Get(BtcpayPath+"/users/me", h.b.GetUser)
	public.Get(BtcpayPath+"/users", h.b.CreateUser)

	public.Get(BtcpayPath+"/stores", h.b.GetStores)
	public.Post(BtcpayPath+"/stores", h.b.CreateStore)
	public.Get(BtcpayPath+"/stores/:storeId", h.b.GetStore)
	public.Put(BtcpayPath+"/stores/:storeId", h.b.UpdateStore)
	public.Delete(BtcpayPath+"/stores/:storeId", h.b.RemoveStore)

	public.Get(BtcpayPath+"/stores/:storeId/payment-requests", h.b.GetPaymentRequests)
	public.Post(BtcpayPath+"/stores/:storeId/payment-requests", h.b.CreatePaymentRequest)
	public.Get(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", h.b.GetPaymentRequest)
	public.Delete(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", h.b.ArchivePaymentRequest)
	public.Put(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", h.b.UpdatePaymentRequest)

	public.Get(BtcpayPath+"/stores/:storeId/pull-payments", h.b.GetPullPayments)
	public.Post(BtcpayPath+"/stores/:storeId/pull-payments", h.b.CreatePullPayment)
	public.Get(BtcpayPath+"/pull-payments/:pullPaymentId", h.b.GetPullPayment)
	public.Delete(BtcpayPath+"/stores/:storeId/pull-payments/:pullPaymentId", h.b.ArchivePullPayment)
	public.Get(BtcpayPath+"/pull-payments/:pullPaymentId/payouts", h.b.GetPayouts)
	public.Post(BtcpayPath+"/pull-payments/:pullPaymentId/payouts", h.b.CreatePayout)
	public.Post(BtcpayPath+"/stores/:storeId/payouts/:payoutId", h.b.ApprovePayout)
	public.Delete(BtcpayPath+"/stores/:storeId/payouts/:payoutId", h.b.CancelPayout)

	public.Get(BtcpayPath+"/stores/:storeId/invoices", h.b.GetInvoices)
	public.Post(BtcpayPath+"/stores/:storeId/invoices", h.b.CreateInvoice)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", h.b.GetInvoice)
	public.Delete(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", h.b.ArchiveInvoice)
	public.Put(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", h.b.UpdateInvoice)
	public.Get(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/payment-methods", h.b.GetInvoicePaymentMethod)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/status", h.b.MarkInvoiceStatus)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/unarchive", h.b.UnarchiveInvoice)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/payment-methods/:paymentMethod/activate", h.b.ActivatePaymentMethod)
}

// func (h *Handler) ApplePayRoutes(public fiber.Router) {
// 	public.Post(ApplePayPath+"/session", h.a.GetApplePaySession)
// 	public.Post(ApplePayPath+"/process", h.a.ProcessApplePayResponse)
// }

// func (h *Handler) GooglePayRoutes(public fiber.Router) {
// 	public.Post(GooglePayPath+"/process", h.g.ProcessGooglePayResponse)
// }
