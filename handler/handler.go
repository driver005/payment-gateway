package handler

import (
	"github.com/driver005/gateway/helper"
)

const (
	NbxplorerPath = "/nbxplorer"
	BtcpayPath    = "/btcpay"
)

func (h *Handler) NbxplorerRoutes(public *helper.RouteGroup) {
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/status", h.n.GetStatus)
	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/rescan", h.n.RescanTransactions)
	// public.POST(NbxplorerPath+"/cryptos/:cryptoCode/rpc", h.n.Rpc)
	// public.POST(NbxplorerPath+"/cryptos/:cryptoCode/connect", h.n.GetConnect)

	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations", h.n.CreateWallet)
	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme", h.n.TrackDerivationScheme)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/transactions", h.n.GetDerivationSchemeTransactions)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/transactions/:txId", h.n.GetDerivationSchemeTransaction)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/balance", h.n.GetCurrentBalance)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/addresses/unused", h.n.NewUnusedAddress)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/scripts/:script", h.n.GetScriptPubKeyInfos)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/utxos", h.n.GetDerivationSchemeUTXOs)
	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/utxos/scan", h.n.ScanUTXOSet)
	// public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/utxos/wipe", h.n.WipeUTXOSet)
	public.PUT(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/metadata/:key", h.n.AttachDerivationSchemeMetadata)
	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/metadata/:key", h.n.DetachDerivationSchemeMetadata)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/metadata/:key", h.n.GetDerivationSchemeMetadata)
	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/prune", h.n.PruneUTXOSet)

	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address", h.n.TrackAddress)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address/transactions", h.n.GetAddressTransactions)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address/transactions/:txId", h.n.GetAddressTransaction)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/addresses/:address/utxos", h.n.GetAddressUTXOs)

	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/transactions", h.n.BroadcastTransaction)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/transactions/:txId", h.n.GetTransaction)

	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/fees/:blockCount", h.n.GetFeeRate)

	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/events", h.n.GetEventStream)
	public.GET(NbxplorerPath+"/cryptos/:cryptoCode/events/latest", h.n.GetRecentEventStream)

	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/derivations/:derivationScheme/psbt/create", h.n.CreatePSBT)
	public.POST(NbxplorerPath+"/cryptos/:cryptoCode/psbt/update", h.n.UpdatePSBT)
}

func (n *Handler) BtcpayRoutes(public *helper.RouteGroup) {
	public.GET(BtcpayPath+"/health", n.b.GetHealth)
	public.GET(BtcpayPath+"/server/info", n.b.GetServerInfo)
	public.GET(BtcpayPath+"/misc/lang", n.b.GetLanguageCodes)
	public.GET(BtcpayPath+"/i/:invoiceId", n.b.GetInvoiceCheckoutPage)

	public.DELETE(BtcpayPath+"/api-keys/:apikey", n.b.RevokeAPIKey)
	public.GET(BtcpayPath+"/api-keys/current", n.b.GetCurrentAPIKey)
	// public.DELETE(BtcpayPath+"/api-keys/current", n.b.RevokeCurrentAPIKey)
	public.POST(BtcpayPath+"/api-keys", n.b.CreateAPIKey)

	public.GET(BtcpayPath+"/api-keys/authorize", n.b.Authorize)

	// public.POST(BtcpayPath+"api/v1/stores/:storeId/apps/pos", n.b.CreateAPIKey)
	// public.GET(BtcpayPath+"api/v1/apps/:appId", n.b.CreateAPIKey)
	// public.DELETE(BtcpayPath+"api/v1/apps/:appId", n.b.CreateAPIKey)

	public.GET(BtcpayPath+"/users/me", n.b.GetUser)
	public.GET(BtcpayPath+"/users", n.b.CreateUser)

	public.GET(BtcpayPath+"/stores", n.b.GetStores)
	public.POST(BtcpayPath+"/stores", n.b.CreateStore)
	public.GET(BtcpayPath+"/stores/:storeId", n.b.GetStore)
	public.PUT(BtcpayPath+"/stores/:storeId", n.b.UpdateStore)
	public.DELETE(BtcpayPath+"/stores/:storeId", n.b.RemoveStore)

	public.GET(BtcpayPath+"/stores/:storeId/payment-requests", n.b.GetPaymentRequests)
	public.POST(BtcpayPath+"/stores/:storeId/payment-requests", n.b.CreatePaymentRequest)
	public.GET(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", n.b.GetPaymentRequest)
	public.DELETE(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", n.b.ArchivePaymentRequest)
	public.PUT(BtcpayPath+"/stores/:storeId/payment-requests/:paymentRequestId", n.b.UpdatePaymentRequest)

	public.GET(BtcpayPath+"/stores/:storeId/pull-payments", n.b.GetPullPayments)
	public.POST(BtcpayPath+"/stores/:storeId/pull-payments", n.b.CreatePullPayment)
	public.GET(BtcpayPath+"/pull-payments/:pullPaymentId", n.b.GetPullPayment)
	public.DELETE(BtcpayPath+"/stores/:storeId/pull-payments/:pullPaymentId", n.b.ArchivePullPayment)
	public.GET(BtcpayPath+"/pull-payments/:pullPaymentId/payouts", n.b.GetPayouts)
	public.POST(BtcpayPath+"/pull-payments/:pullPaymentId/payouts", n.b.CreatePayout)
	public.POST(BtcpayPath+"/stores/:storeId/payouts/:payoutId", n.b.ApprovePayout)
	public.DELETE(BtcpayPath+"/stores/:storeId/payouts/:payoutId", n.b.CancelPayout)

	public.GET(BtcpayPath+"/stores/:storeId/invoices", n.b.GetInvoices)
	public.POST(BtcpayPath+"/stores/:storeId/invoices", n.b.CreateInvoice)
	public.POST(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", n.b.GetInvoice)
	public.DELETE(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", n.b.ArchiveInvoice)
	public.PUT(BtcpayPath+"/stores/:storeId/invoices/:invoiceId", n.b.UpdateInvoice)
	public.GET(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/payment-methods", n.b.GetInvoicePaymentMethod)
	public.POST(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/status", n.b.MarkInvoiceStatus)
	public.POST(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/unarchive", n.b.UnarchiveInvoice)
	public.POST(BtcpayPath+"/stores/:storeId/invoices/:invoiceId/payment-methods/:paymentMethod/activate", n.b.ActivatePaymentMethod)
}

// func (h *Handler) SetRoutes(public *helper.RouterPublic) {
// 	h.CardRoutes(public)
// }

// func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	respond.NewResponse(w).Ok("Working")
// }

// func (h *Handler) CardRoutes(public *helper.RouterPublic) {
// 	public.GET(ClientsHandlerPath, h.ListCard)
// 	public.POST(ClientsHandlerPath, h.CreateCard)
// 	public.GET(ClientsHandlerPath+"/:id", h.GetCard)
// 	public.PUT(ClientsHandlerPath+"/:id", h.UpdateCard)
// 	public.DELETE(ClientsHandlerPath+"/:id", h.DeleteCard)
// }

// func (h *Handler) Routes(public *helper.RouteGroup) {
// 	public.GET("/", h.GetHealth)
// }
