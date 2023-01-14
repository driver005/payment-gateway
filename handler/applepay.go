package handler

type ApplePay struct {
	r Registry
}

func NewApplePay() *ApplePay {
	return &ApplePay{}
}

// ApplePayResponse - ...
type ApplePaySessionResponse struct {
	EpochTimestamp            int64  `json:"epochTimestamp"`
	ExpiresAt                 int64  `json:"expiresAt"`
	MerchantSessionIdentifier string `json:"merchantSessionIdentifier"`
	Nonce                     string `json:"nonce"`
	MerchantIdentifier        string `json:"merchantIdentifier"`
	DomainName                string `json:"domainName"`
	DisplayName               string `json:"displayName"`
	Signature                 string `json:"signature"`
}

// func (a *ApplePay) GetApplePaySession(context *fiber.Ctx) error {
// 	appleData := &struct {
// 		URL string `json:"url"`
// 	}{}
// 	if err := context.BodyParser(&appleData); err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "error",
// 		})
// 	}

// 	payload, err := a.r.Applepay().Session(appleData.URL)
// 	if err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "error",
// 		})
// 	}

// 	// s.logger.Debugw("Apple Session", "url", appleData.URL, "session", payload)

// 	// Payload is a JSON string in byte array format
// 	var response ApplePaySessionResponse
// 	err = json.Unmarshal(payload, &response)
// 	if err != nil {
// 		return context.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "error",
// 		})
// 	}

// 	return context.JSON(response)
// }

// func (a *ApplePay) ProcessApplePayResponse(context *fiber.Ctx) error {
// 	response := types.ApplePayAuthorizedPaymentResponse{}

// 	if err := context.BodyParser(&response); err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "error",
// 		})
// 	}

// 	input := a.r.Service().NewApplePayPaymentTokenInput(&response)

// 	paymentToken, err := a.r.Service().CreatePaymentToken(context.Context(), input)
// 	if err != nil {
// 		if a.r.Service().IsConnectionInvalid(err) {
// 			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"message": err.Error(),
// 				"type":    "error",
// 			})
// 		} else {
// 			a.r.Logger().Error("Create Payment Token from Apple Pay", "Params", input)
// 			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": err.Error(),
// 				"type":    "error",
// 			})
// 		}
// 	}

// 	// We now have the payment data.
// 	return context.JSON(&types.PublicPaymentToken{
// 		ID:       paymentToken.ID,
// 		Object:   paymentToken.Object,
// 		Amount:   paymentToken.Amount,
// 		Currency: paymentToken.Currency,
// 	})

// }
