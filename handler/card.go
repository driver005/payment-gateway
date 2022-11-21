package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/driver005/database"
// 	"github.com/driver005/gateway/helper"
// 	models "github.com/driver005/gateway/models_old"
// 	resp "github.com/driver005/gateway/respond"
// 	"github.com/gofrs/uuid"
// 	"github.com/julienschmidt/httprouter"
// )

// func (h *Handler) CreateCard(context *fiber.Ctx) error {
// 	var m models.Card

// 	if err := context.BodyParser(&m); err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	if err := m.Validate(); err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	m.CreatedAt = time.Now().UTC().Round(time.Second)
// 	m.UpdatedAt = m.CreatedAt
// 	if err := h.r.Manager(context.Context()).Create(&m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	resp.NewResponse(w).AddHeader("Location", ClientsHandlerPath).Ok(&m)
// }

// func (h *Handler) UpdateCard(context *fiber.Ctx) error {
// 	var m models.Card
// 	var o models.Card
// 	var err error

// 	if err := context.BodyParser(&m); err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	m.ID, err = uuid.FromString(context.Params("id"))
// 	if err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 	}

// 	if err = h.r.Manager(context.Context()).Where("id = ?", m.ID).First(&o).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	m.ID = o.ID
// 	m.UpdatedAt = time.Now().UTC().Round(time.Second)
// 	if err := h.r.Manager(context.Context()).Model(&o).Updates(m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}
// 	resp.NewResponse(w).Ok(&m)
// }

// func Paginate(r *http.Request, offset int, pageSize int) func(db *database.DB) *database.DB {
// 	return func(db *database.DB) *database.DB {
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }

// func (h *Handler) ListCard(context *fiber.Ctx) error {
// 	var m = make([]models.Card, 0)
// 	q := r.URL.Query()
// 	page, _ := strconv.Atoi(context.Query("page"))
// 	if page == 0 {
// 		page = 1
// 	}

// 	pageSize, _ := strconv.Atoi(context.Query("page_size"))
// 	switch {
// 	case pageSize > 100:
// 		pageSize = 100
// 	case pageSize <= 0:
// 		pageSize = 10
// 	}

// 	offset := (page - 1) * pageSize
// 	if err := h.r.Manager(context.Context()).Scopes(Paginate(offset, pageSize)).Order("id").Find(&m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	n := h.r.Manager(context.Context()).Find(&models.Card{})
// 	if n.Error != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(n.Error))
// 		return
// 	}

// 	helper.Header(context, context.Request().URI(), n.RowsAffected, pageSize, offset)

// 	if m == nil {
// 		m = []models.Card{}
// 	}

// 	resp.NewResponse(w).Ok(&m)
// }

// func (h *Handler) GetCard(context *fiber.Ctx) error {
// 	var m models.Card
// 	var id = context.Params("id")

// 	if err := h.r.Manager(context.Context()).Where("id = ?", id).First(&m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	resp.NewResponse(w).Ok(&m)
// }

// func (h *Handler) DeleteCard(context *fiber.Ctx) error {
// 	var id = context.Params("id")

// 	ID, err := uuid.FromString(id)
// 	if err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 	}

// 	if err := h.r.Manager(context.Context()).Delete(&models.Card{ID: ID}).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	resp.NewResponse(w).NoContent()
// }
