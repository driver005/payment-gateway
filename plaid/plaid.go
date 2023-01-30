package plaid

import (
	"context"
	"fmt"

	"github.com/plaid/plaid-go/plaid"
)

func InitPlaid() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", "63c426c5ff1ded001227a2d5")
	configuration.AddDefaultHeader("PLAID-SECRET", "d63fdd5b20f4077bd0f245ec9483ff")
	configuration.UseEnvironment(plaid.Sandbox)
	return plaid.NewAPIClient(configuration)
}

func (h *Handler) GenerateLinkToken(ctx context.Context) (string, error) {
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "1",
	}
	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	request.SetProducts([]plaid.Products{
		plaid.PRODUCTS_TRANSACTIONS,
		plaid.PRODUCTS_AUTH,
		plaid.PRODUCTS_IDENTITY,
	})
	request.SetLinkCustomizationName("default")
	request.SetWebhook("https://webhook.site/4a819294-d934-451d-9c99-72bd80fd3065")
	request.SetRedirectUri("http://localhost/link")
	resp, _, err := h.client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		return "", err
	}
	return resp.GetLinkToken(), nil
}

func (h *Handler) ExchangePublicToken(ctx context.Context, publicToken string) (string, string, error) {
	// exchange the public_token for an access_token
	exchangePublicTokenReq := plaid.NewItemPublicTokenExchangeRequest(publicToken)
	exchangePublicTokenResp, _, err := h.client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*exchangePublicTokenReq).Execute()
	if err != nil {
		return "", "", err
	}
	// These values should be saved to a persistent database and
	// associated with the currently signed-in user
	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()

	return accessToken, itemID, nil
}

func (h *Handler) Auth(ctx context.Context, accessToken string) (string, error) {
	authGetResp, _, err := h.client.PlaidApi.AuthGet(ctx).AuthGetRequest(*plaid.NewAuthGetRequest(accessToken)).Execute()
	if err != nil {
		return "", err
	}
	return authGetResp.RequestId, nil

}

type Webbhook struct {
	WebhookType string `json:"webhook_type,omitempty"`
	WebhookCode string `json:"webhook_code,omitempty"`
	ItemId      string `json:"item_id,omitempty"`
	AccountId   string `json:"account_id,omitempty"`
	Environment string `json:"environment,omitempty"`
}

func (h *Handler) Webbhook(w Webbhook) {
	if w.WebhookType == "AUTH" {
		if w.WebhookCode == "VERIFICATION_EXPIRED" {
			fmt.Println("EXPIRED")
		} else if w.WebhookCode == "AUTOMATICALLY_VERIFIED" {
			fmt.Println("VERIFIED")
		}
	}
}
