package handler

type GooglePay struct {
	r Registry
}

func NewGooglePay() *GooglePay {
	return &GooglePay{}
}

// func (g *GooglePay) ProcessGooglePayResponse(context *fiber.Ctx) error {
// 	response := &types.GooglePayAuthorizedPaymentResponse{}
// 	if err := context.BodyParser(&response); err != nil {
// 		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 			"type":    "error",
// 		})
// 	}

// 	input := g.r.Service().NewGooglePayPaymentTokenInput(response)

// 	paymentToken, err := g.r.Service().CreatePaymentToken(context.Context(), input)
// 	if err != nil {
// 		if g.r.Service().IsConnectionInvalid(err) {
// 			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"message": err.Error(),
// 				"type":    "error",
// 			})
// 		} else {
// 			g.r.Logger().Error("Create Payment Token from Google Pay", "Params", input)
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
