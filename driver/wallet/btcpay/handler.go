package btcpay

import (
	"context"
	"fmt"
	"log"
)

type Handler struct {
	client  *Client
	storeID StoreID
	ctx     context.Context
}

func NewHandler() *Handler {
	handler := &Handler{
		client:  &Client{},
		storeID: "",
		ctx:     context.Background(),
	}
	return handler
}

func (h *Handler) SetStoreID() {
	stores, code, err := h.client.GetStores(h.ctx)
	if err != nil {
		fmt.Println(code)
		panic(err)
	}

	for _, v := range stores {
		if v.Name == "Gateway" {
			h.storeID = v.ID
		}
	}
}

func (h *Handler) SetClient(url string, key string) {
	// h.client = NewBasicClient(url, username, password)
	h.client = NewClient(url, APIKey(key))
}

func (h *Handler) GetAndDeleteNotification() {
	notifications, _, err := h.client.GetNotifications(h.ctx)
	if err != nil {
		panic(err)
	}
	nClient := h.client.Notification
	if len(notifications) > 0 {
		nClient.ID = notifications[0].ID
		fmt.Print(nClient.RemoveNotification(h.ctx))
	}
}
func (h *Handler) GetAndPutNotification() {
	notifications, _, err := h.client.GetNotifications(h.ctx)
	if err != nil {
		panic(err)
	}
	nClient := h.client.Notification
	nClient.ID = notifications[0].ID
	fmt.Print(nClient.UpdateNotification(h.ctx))
}
func (h *Handler) GetNotifications() {
	notifications, _, err := h.client.GetNotifications(h.ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range notifications {
		fmt.Println(v)
	}
}
func (h *Handler) GetPaymentRequests() {
	paymentRequests, _, err := h.client.GetPaymentRequests(h.ctx, &h.storeID)
	if err != nil {
		panic(err)
	}
	for _, v := range paymentRequests {
		fmt.Println(v)
		h.client.ArchivePaymentRequest(h.ctx, &h.storeID, &v.ID)
	}
}
func (h *Handler) CreateInvoicePage() *InvoiceCheckoutPage {
	invoice, _, err := h.client.CreateInvoice(h.ctx, &h.storeID, &InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
	})
	if err != nil {
		panic(err)
	}
	page, _, err := h.client.GetInvoiceCheckoutPage(h.ctx, &invoice.ID, "de-DE")
	if err != nil {
		panic(err)
	}

	return page
}

func (h *Handler) CreateInvoice() *InvoiceResponse {
	invoice, _, err := h.client.CreateInvoice(h.ctx, &h.storeID, &InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
	})
	if err != nil {
		panic(err)
	}

	return invoice
}

func (h *Handler) CreateNewUser() {
	fmt.Println(h.client.CreateUser(h.ctx, &UserRequest{
		Email:           "test@test.com",
		Password:        "asdfasdf",
		IsAdministrator: false,
	}))
}

func (h *Handler) CreateNewStore() {
	fmt.Println(h.client.CreateStore(h.ctx, &StoreRequest{Name: "test03"}))
}

func (h *Handler) CreateInvoiceByStoreGetAndDeleteInvoiceByID() {
	// create a new storeClient
	storeClient := h.client.Store
	// assign a store ID to the storeClient
	storeClient.ID = h.storeID
	// create a new invoice for the store
	invoice, _, err := storeClient.CreateInvoice(h.ctx, &InvoiceRequest{Amount: "10", Currency: "USD"})
	if err != nil {
		panic(err)
	}
	// create a new invoiceClient, based on the current client
	invoiceClient := h.client.Invoice
	// assign a storeClient to the invoiceClient
	invoiceClient.Store = storeClient
	// assign a invoice ID to the invoiceClient
	invoiceClient.ID = invoice.ID
	fmt.Println(invoiceClient.GetInvoice(h.ctx))
	invoiceClient.ArchiveInvoice(h.ctx)
}
func (h *Handler) GetInvoicesByStore() {
	storeClient := h.client.Store
	storeClient.ID = h.storeID
	fmt.Println(storeClient.GetInvoices(h.ctx))
}

func (h *Handler) CreateAndDeleteInvoice() {
	fmt.Println(h.storeID)
	invoice, code, err := h.client.CreateInvoice(h.ctx, &h.storeID, &InvoiceRequest{Amount: "11", Currency: "USD", Metadata: InvoiceMetadata{"test": "asdf", "test2": "aaaa"}})
	if err != nil {
		fmt.Println(code)
		panic(err)
	}
	fmt.Println(invoice)
	h.client.UpdateInvoice(h.ctx, &h.storeID, &invoice.ID, &InvoiceUpdate{Metadata: InvoiceMetadata{"test3": "ccccc"}})
	fmt.Println(h.client.GetInvoices(h.ctx, &h.storeID))
	h.client.ArchiveInvoice(h.ctx, &h.storeID, &invoice.ID)
}

func (h *Handler) CreateAndDeleteAPIKey() {
	// create new APIKey
	apiKey, _, err := h.client.CreateAPIKey(h.ctx, &APIKeyRequest{
		Permissions: []BTCPayPermission{CreateCustomPermission(GetPermission().StoreCanviewinvoices, StoreID("66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn"))}})
	if err != nil {
		panic(err)
	}

	// delete the new APIKey
	_, err = h.client.RevokeAPIKey(h.ctx, &apiKey.APIKey)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) CreateAndDeleteCurrentAPIKey() {
	// create new APIKey
	apiKey, statusCode, err := h.client.CreateAPIKey(h.ctx, &APIKeyRequest{
		Permissions: []BTCPayPermission{CreateCustomPermission(GetPermission().StoreCanviewinvoices, StoreID("66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn"))}})
	if err != nil {
		log.Fatal(statusCode, err)
	}

	// add APIKey to client
	h.client.APIKey = apiKey.APIKey
	fmt.Println(h.client.GetCurrentAPIKey(h.ctx))
	h.client.RevokeCurrentAPIKey(h.ctx)
}

func start() {
	// client := NewBasicClient("http://localhost:14142/", "paulwolf005@gmail.com", "B9wwndwwnj8!")

	/* cont, cancel := context.WithTimeout(h.ctx, time.Second*5)
	defer cancel()
	fmt.Println(client.GetHealth(cont)) */

	//createAndDeleteCurrentAPIKey(client)

	// fmt.Println(client.GetServerInfo(ctx))

	// get store id

	/* stores, code, err := client.GetStores(h.ctx)
	if err != nil {
		fmt.Println(code)
		panic(err)
	}
	storeID := getStoreID(stores) */
	//createInvoiceByStoreGetAndDeleteInvoiceByID(client, storeID)

	//getInvoicesByStore(client, storeID)
	//createAndDeleteInvoice(client, storeID)

	//createNewStore(client)

	// createNewUser(client)
	//fmt.Println(client.GetUser(h.ctx))

	/* fmt.Println((&btcpay.Client{URL: client.URL}).GetLanguageCodes(h.ctx))
	langs, _, err := client.GetLanguageCodes(h.ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range langs {
		fmt.Println(v)
	} */

	//createPrintPageDeleteInvoice(client, &storeID)

	//getPaymentRequests(client, &storeID)

	//GetNotifications(client)
	//getAndDeleteNotification(client)
}
