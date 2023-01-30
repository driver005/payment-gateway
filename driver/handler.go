package driver

import (
	"encoding/base64"
	"fmt"
	"image/png"
	"net/url"

	"github.com/aaronarduino/goqrsvg"
	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode"
	"github.com/driver005/gateway/driver/realtime/paynow"
	"github.com/driver005/gateway/driver/realtime/pix"
	"github.com/driver005/gateway/driver/realtime/promptpay"
	"github.com/driver005/gateway/driver/voucher/boleto"
	"github.com/driver005/gateway/driver/wallet/alipay"
	"github.com/driver005/gateway/pdf"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SetRoutes(r fiber.Router) {
	r.Post("/notify", h.RouteNotify)
	r.Get("/code", h.RouteQrCode)
}

func (h *Handler) RouteNotify(context *fiber.Ctx) error {
	val, err := url.ParseQuery(string(context.Request().URI().QueryString()))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	ok, err := h.alipay.VerifySign(val)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	if !ok {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}

	var p = alipay.TradeQuery{}
	p.OutTradeNo = context.Query("out_trade_no")
	rsp, err := h.alipay.TradeQuery(p)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "internal_error",
		})
	}
	if !rsp.IsSuccess() {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("error %s: %s-%s \n", context.Query("out_trade_no"), rsp.Content.Msg, rsp.Content.SubMsg),
			"type":    "internal_error",
		})
	}

	return context.Status(fiber.StatusOK).JSON(context.Query("out_trade_no"))
}

func (h *Handler) RouteQrCode(context *fiber.Ctx) error {
	var code barcode.Barcode
	context.Type(context.Query("type"))

	data, err := base64.URLEncoding.DecodeString(context.Query("data"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "parse_error",
		})
	}

	if context.Query("paymentMethod") == "pix" {
		code, err = pix.QRCode(pix.QRCodeOptions{
			Content: string(data),
		})
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "intenal_error",
			})
		}
	} else if context.Query("paymentMethod") == "promptpay" {
		code, err = promptpay.QRCode(promptpay.QRCodeOptions{
			Content: string(data),
		})
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "intenal_error",
			})
		}
	} else if context.Query("paymentMethod") == "paynow" {
		code, err = paynow.QRCode(paynow.QRCodeOptions{
			Content: string(data),
		})
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"type":    "intenal_error",
			})
		}
	} else if context.Query("paymentMethod") == "boleto" {
		if context.Query("type") == "pdf" {
			code, err := boleto.GenrateBarcode(string(data))
			if err != nil {
				return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
					"type":    "intenal_error",
				})
			}

			url := fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", context.Query("data"), "boleto", "png")
			pdf.GenerateWithBarcode(context.Response().BodyWriter(), "Boleto", url, code, string(data))
			return nil
		} else {
			code, err = boleto.GenrateBarcode(string(data))
			if err != nil {
				return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
					"type":    "intenal_error",
				})
			}
		}
	} else {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "payment method not supportet",
			"type":    "intenal_error",
		})
	}

	if context.Query("type") == "png" {
		png.Encode(context.Response().BodyWriter(), code)
		return nil

	} else if context.Query("type") == "svg" {
		canvas := svg.New(context.Response().BodyWriter())
		qs := goqrsvg.NewQrSVG(code, 5)
		qs.StartQrSVG(canvas)
		qs.WriteQrSVG(canvas)
		canvas.End()
		return nil
	} else {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "code image type not supportet",
			"type":    "intenal_error",
		})
	}
}
