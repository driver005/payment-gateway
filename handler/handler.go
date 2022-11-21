package handler

import (
	"github.com/gofiber/fiber/v2"
)

const (
	NbxplorerPath = "/nbxplorer"
	BtcpayPath    = "/btcpay"
)

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

func (n *Handler) BtcpayRoutes(public fiber.Router) {
	public.Get(BtcpayPath+"/health", n.b.GetHealth)
	public.Get(BtcpayPath+"/server/info", n.b.GetServerInfo)
	public.Get(BtcpayPath+"/misc/lang", n.b.GetLanguageCodes)
	public.Get(BtcpayPath+"/i/:invoiceId", n.b.GetInvoiceCheckoutPage)

	public.Delete(BtcpayPath+"/api-keys/:apikey", n.b.RevokeAPIKey)
	public.Get(BtcpayPath+"/api-keys/current", n.b.GetCurrentAPIKey)
	// public.Delete(BtcpayPath+"/api-keys/current", n.b.RevokeCurrentAPIKey)
	public.Post(BtcpayPath+"/api-keys", n.b.CreateAPIKey)

	public.Get(BtcpayPath+"/api-keys/authorize", n.b.Authorize)

	// public.Post(BtcpayPath+"api/v1/stores/:storeId/apps/pos", n.b.CreateAPIKey)
	// public.Get(BtcpayPath+"api/v1/apps/:appId", n.b.CreateAPIKey)
	// public.Delete(BtcpayPath+"api/v1/apps/:appId", n.b.CreateAPIKey)

	public.Get(BtcpayPath+"/users/me", n.b.GetUser)
	public.Get(BtcpayPath+"/users", n.b.CreateUser)

	public.Get(BtcpayPath+"/stores", n.b.GetStores)
	public.Post(BtcpayPath+"/stores", n.b.CreateStore)
	public.Get(BtcpayPath+"/stores/:storeId", n.b.GetStore)
	public.Put(BtcpayPath+"/stores/:storeId", n.b.UpdateStore)
	public.Delete(BtcpayPath+"/stores/:storeId", n.b.RemoveStore)

	public.Get(BtcpayPath+"/stores/:storeId/payment-requests", n.b.GetPaymentRequests)
	public.Post(BtcpayPath+"/stores/:storeId/payment-requests", n.b.CreatePaymentRequest)
	public.Get(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", n.b.GetPaymentRequest)
	public.Delete(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", n.b.ArchivePaymentRequest)
	public.Put(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", n.b.UpdatePaymentRequest)

	public.Get(BtcpayPath+"/stores/:storeId/pull-payments", n.b.GetPullPayments)
	public.Post(BtcpayPath+"/stores/:storeId/pull-payments", n.b.CreatePullPayment)
	public.Get(BtcpayPath+"/pull-payments/:pullPaymentId", n.b.GetPullPayment)
	public.Delete(BtcpayPath+"/stores/:storeId/pull-payments/:pullPaymentId", n.b.ArchivePullPayment)
	public.Get(BtcpayPath+"/pull-payments/:pullPaymentId/payouts", n.b.GetPayouts)
	public.Post(BtcpayPath+"/pull-payments/:pullPaymentId/payouts", n.b.CreatePayout)
	public.Post(BtcpayPath+"/stores/:storeId/payouts/:payoutId", n.b.ApprovePayout)
	public.Delete(BtcpayPath+"/stores/:storeId/payouts/:payoutId", n.b.CancelPayout)

	public.Get(BtcpayPath+"/stores/:storeId/invoices", n.b.GetInvoices)
	public.Post(BtcpayPath+"/stores/:storeId/invoices", n.b.CreateInvoice)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", n.b.GetInvoice)
	public.Delete(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", n.b.ArchiveInvoice)
	public.Put(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", n.b.UpdateInvoice)
	public.Get(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/payment-methods", n.b.GetInvoicePaymentMethod)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/status", n.b.MarkInvoiceStatus)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/unarchive", n.b.UnarchiveInvoice)
	public.Post(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/payment-methods/:paymentMethod/activate", n.b.ActivatePaymentMethod)
}

// func (h *Handler) SetRoutes(public *helper.RouterPublic) {
// 	h.CardRoutes(public)
// }

// func (h *Handler) GetHealth(context *fiber.Ctx) error {
// 	respond.NewResponse(w).Ok("Working")
// }

// func (h *Handler) CardRoutes(public *helper.RouterPublic) {
// 	public.Get(ClientsHandlerPath, h.ListCard)
// 	public.Post(ClientsHandlerPath, h.CreateCard)
// 	public.Get(ClientsHandlerPath+"/:id", h.GetCard)
// 	public.Put(ClientsHandlerPath+"/:id", h.UpdateCard)
// 	public.Delete(ClientsHandlerPath+"/:id", h.DeleteCard)
// }

// func (h *Handler) Routes(public fiber.Router) {
// 	public.Get("/", h.GetHealth)
// }
