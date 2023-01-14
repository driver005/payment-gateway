package riverty

import "encoding/json"

const (
	checkoutAuthorize = "/api/v3/checkout/authorize"
)

type (
	Riverty struct {
		client Client
	}

	Address struct {
		Street                 string `json:"street"`
		StreetNumber           string `json:"streetNumber"`
		StreetNumberAdditional string `json:"streetNumberAdditional"`
		PostalCode             string `json:"postalCode"`
		PostalPlace            string `json:"postalPlace"`
		CountryCode            string `json:"countryCode"`
		CareOf                 string `json:"careOf"`
	}

	RiskData struct {
		ExistingCustomer               bool    `json:"existingCustomer,omitempty"`
		VerifiedCustomerIdentification bool    `json:"verifiedCustomerIdentification,omitempty"`
		MarketingOptIn                 bool    `json:"marketingOptIn,omitempty"`
		CustomerSince                  string  `json:"customerSince,omitempty"`
		CustomerClassification         string  `json:"customerClassification,omitempty"`
		AcquisitionChannel             string  `json:"acquisitionChannel,omitempty"`
		HasCustomerCard                bool    `json:"hasCustomerCard,omitempty"`
		CustomerCardSince              string  `json:"customerCardSince,omitempty"`
		CustomerCardClassification     string  `json:"customerCardClassification,omitempty"`
		ProfileTrackingID              string  `json:"profileTrackingId,omitempty"`
		IPAddress                      string  `json:"ipAddress"`
		NumberOfTransactions           int     `json:"numberOfTransactions,omitempty"`
		CustomerIndividualScore        string  `json:"customerIndividualScore,omitempty"`
		UserAgent                      string  `json:"userAgent,omitempty"`
		AmountOfTransactions           float64 `json:"amountOfTransactions,omitempty"`
		OtherPaymentMethods            bool    `json:"otherPaymentMethods,omitempty"`
	}

	Items struct {
		ProductID             string  `json:"productId"`
		GroupID               string  `json:"groupId,omitempty"`
		Description           string  `json:"description"`
		Type                  string  `json:"type,omitempty"`
		NetUnitPrice          float64 `json:"netUnitPrice"`
		GrossUnitPrice        float64 `json:"grossUnitPrice"`
		Quantity              float64 `json:"quantity"`
		VatPercent            float64 `json:"vatPercent"`
		VatAmount             float64 `json:"vatAmount"`
		VatCategory           string  `json:"vatCategory"`
		ImageURL              string  `json:"imageUrl,omitempty"`
		ProductURL            string  `json:"productUrl,omitempty"`
		MarketPlaceSellerID   string  `json:"marketPlaceSellerId,omitempty"`
		AdditionalInformation string  `json:"additionalInformation,omitempty"`
		Specification         string  `json:"specification,omitempty"`
		ServiceStart          string  `json:"serviceStart,omitempty"`
		ServiceEnd            string  `json:"serviceEnd,omitempty"`
	}

	Payment struct {
		Type        string `json:"type"`
		ContractID  string `json:"contractId,omitempty"`
		DirectDebit struct {
			BankCode    string `json:"bankCode,omitempty"`
			BankAccount string `json:"bankAccount,omitempty"`
			Token       string `json:"token,omitempty"`
		} `json:"directDebit,omitempty"`
		RecurringDirectDebit struct {
			DirectDebitAgreement string `json:"directDebitAgreement,omitempty"`
			BankCode             string `json:"bankCode,omitempty"`
			BankAccount          string `json:"bankAccount,omitempty"`
		} `json:"recurringDirectDebit,omitempty"`
		Campaign struct {
			CampaignNumber int `json:"campaignNumber,omitempty"`
		} `json:"campaign,omitempty"`
		Invoice string `json:"invoice,omitempty"`
		Account struct {
			ProfileNo int `json:"profileNo,omitempty"`
		} `json:"account,omitempty"`
		ConsolidatedInvoice struct {
			InvoiceDate string `json:"invoiceDate,omitempty"`
		} `json:"consolidatedInvoice,omitempty"`
		Installment struct {
			ProfileNo            int     `json:"profileNo,omitempty"`
			NumberOfInstallments int     `json:"numberOfInstallments,omitempty"`
			CustomerInterestRate float64 `json:"customerInterestRate,omitempty"`
		} `json:"installment,omitempty"`
		// PayinX string `json:"payinX,omitempty"`
	}

	Customer struct {
		IdentificationNumber string   `json:"identificationNumber,omitempty"`
		Address              Address  `json:"address"`
		LegalForm            string   `json:"legalForm,omitempty"`
		RiskData             RiskData `json:"riskData"`
		CustomerNumber       string   `json:"customerNumber,omitempty"`
		Salutation           string   `json:"salutation,omitempty"`
		FirstName            string   `json:"firstName"`
		LastName             string   `json:"lastName"`
		CompanyName          string   `json:"companyName,omitempty"`
		Email                string   `json:"email"`
		Phone                string   `json:"phone,omitempty"`
		MobilePhone          string   `json:"mobilePhone,omitempty"`
		BirthDate            string   `json:"birthDate"`
		CustomerCategory     string   `json:"customerCategory"`
		ConversationLanguage string   `json:"conversationLanguage,omitempty"`
		DistributionType     string   `json:"distributionType,omitempty"`
		VatID                string   `json:"vatId,omitempty"`
	}

	Order struct {
		Number           string  `json:"number,omitempty"`
		TotalNetAmount   float64 `json:"totalNetAmount"`
		TotalGrossAmount float64 `json:"totalGrossAmount"`
		Currency         string  `json:"currency"`
		MerchantImageURL string  `json:"merchantImageUrl,omitempty"`
		Items            []Items `json:"items"`
		ProductUser      string  `json:"productUser,omitempty"`
		CostCenter       string  `json:"costCenter,omitempty"`
	}

	Checkout struct {
		CheckoutID                 string   `json:"checkoutId,omitempty"`
		Payment                    Payment  `json:"payment"`
		Customer                   Customer `json:"customer"`
		Order                      Order    `json:"order"`
		ParentTransactionReference string   `json:"parentTransactionReference,omitempty"`
		YourReference              string   `json:"yourReference,omitempty"`
		OurReference               string   `json:"ourReference,omitempty"`
		Nonce                      string   `json:"nonce,omitempty"`
	}

	Response struct {
		Outcome  string `json:"outcome"`
		Customer struct {
			CustomerNumber    string `json:"customerNumber"`
			CustomerAccountID string `json:"customerAccountId"`
			FirstName         string `json:"firstName"`
			LastName          string `json:"lastName"`
			AddressList       []struct {
				Street                 string `json:"street"`
				StreetNumber           string `json:"streetNumber"`
				StreetNumberAdditional string `json:"streetNumberAdditional"`
				PostalCode             string `json:"postalCode"`
				PostalPlace            string `json:"postalPlace"`
				CountryCode            string `json:"countryCode"`
				CareOf                 string `json:"careOf"`
			} `json:"addressList"`
		} `json:"customer"`
		DeliveryCustomer struct {
			CustomerNumber    string `json:"customerNumber"`
			CustomerAccountID string `json:"customerAccountId"`
			FirstName         string `json:"firstName"`
			LastName          string `json:"lastName"`
			AddressList       []struct {
				Street                 string `json:"street"`
				StreetNumber           string `json:"streetNumber"`
				StreetNumberAdditional string `json:"streetNumberAdditional"`
				PostalCode             string `json:"postalCode"`
				PostalPlace            string `json:"postalPlace"`
				CountryCode            string `json:"countryCode"`
				CareOf                 string `json:"careOf"`
			} `json:"addressList"`
		} `json:"deliveryCustomer"`
		SecureLoginURL    string `json:"secureLoginUrl"`
		RiskCheckMessages []struct {
			Type                  string `json:"type"`
			Code                  string `json:"code"`
			Message               string `json:"message"`
			CustomerFacingMessage string `json:"customerFacingMessage"`
			ActionCode            string `json:"actionCode"`
			FieldReference        string `json:"fieldReference"`
		} `json:"riskCheckMessages"`
		ExpirationDate string `json:"expirationDate"`
	}
)

func NewRiverty(c Client) Riverty {
	return Riverty{client: c}
}

func (rv *Riverty) CreateCheckout(c *Checkout) (*Response, error) {
	var m Response
	res, err := rv.client.Post(checkoutAuthorize, c)
	if nil != err {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}

	return &m, nil
}