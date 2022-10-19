package handler

import "github.com/driver005/gateway/helper"

const (
	ClientsHandlerPath = "/clients"
)

func (h *Handler) SetRoutes(public *helper.RouterPublic) {
	h.CardRoutes(public)
}

func (h *Handler) CardRoutes(public *helper.RouterPublic) {
	public.GET(ClientsHandlerPath, h.ListCard)
	public.POST(ClientsHandlerPath, h.CreateCard)
	public.GET(ClientsHandlerPath+"/:id", h.GetCard)
	public.PUT(ClientsHandlerPath+"/:id", h.UpdateCard)
	public.DELETE(ClientsHandlerPath+"/:id", h.DeleteCard)
}
