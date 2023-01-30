package plaid

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Post("/create_link_token", h.LinkTokenRoute)
	r.Post("/exchange_public_token", h.AccessTokenRoute)
	r.Get("/link", h.RenderLink)
	r.Get("/auth", h.AuthRoute)
	r.Get("/webhook", h.WebbhookRoute)
}

func (h *Handler) LinkTokenRoute(context *fiber.Ctx) error {
	tocken, err := h.GenerateLinkToken(context.Context())
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"link_token": tocken,
	})
}

func (h *Handler) AccessTokenRoute(context *fiber.Ctx) error {
	token, item, err := h.ExchangePublicToken(context.Context(), context.FormValue("public_token"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	fmt.Println(token)
	fmt.Println(item)
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"public_token_exchange": "complete",
	})
}

func (h *Handler) AuthRoute(context *fiber.Ctx) error {
	_, err := h.Auth(context.Context(), "access-sandbox-c7cb4de3-0cd4-4253-a47d-20e67ec1a753")
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"public_token_exchange": "complete",
	})
}

func (h *Handler) WebbhookRoute(context *fiber.Ctx) error {
	var m Webbhook
	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "database_error",
		})
	}

	h.Webbhook(m)

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

func (h *Handler) RenderLink(context *fiber.Ctx) error {
	return context.Render("plaid", fiber.Map{})
}
