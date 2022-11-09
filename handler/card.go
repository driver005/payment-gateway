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

// func (h *Handler) CreateCard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	var m models.Card

// 	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	if err := m.Validate(); err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	m.CreatedAt = time.Now().UTC().Round(time.Second)
// 	m.UpdatedAt = m.CreatedAt
// 	if err := h.r.Manager(r.Context()).Create(&m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	resp.NewResponse(w).AddHeader("Location", ClientsHandlerPath).Ok(&m)
// }

// func (h *Handler) UpdateCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	var m models.Card
// 	var o models.Card
// 	var err error

// 	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	m.ID, err = uuid.FromString(ps.ByName("id"))
// 	if err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 	}

// 	if err = h.r.Manager(r.Context()).Where("id = ?", m.ID).First(&o).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	m.ID = o.ID
// 	m.UpdatedAt = time.Now().UTC().Round(time.Second)
// 	if err := h.r.Manager(r.Context()).Model(&o).Updates(m).Error; err != nil {
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

// func (h *Handler) ListCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	var m = make([]models.Card, 0)
// 	q := r.URL.Query()
// 	page, _ := strconv.Atoi(q.Get("page"))
// 	if page == 0 {
// 		page = 1
// 	}

// 	pageSize, _ := strconv.Atoi(q.Get("page_size"))
// 	switch {
// 	case pageSize > 100:
// 		pageSize = 100
// 	case pageSize <= 0:
// 		pageSize = 10
// 	}

// 	offset := (page - 1) * pageSize
// 	if err := h.r.Manager(r.Context()).Scopes(Paginate(r, offset, pageSize)).Order("id").Find(&m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	n := h.r.Manager(r.Context()).Find(&models.Card{})
// 	if n.Error != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(n.Error))
// 		return
// 	}

// 	helper.Header(w, r.URL, n.RowsAffected, pageSize, offset)

// 	if m == nil {
// 		m = []models.Card{}
// 	}

// 	resp.NewResponse(w).Ok(&m)
// }

// func (h *Handler) GetCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	var m models.Card
// 	var id = ps.ByName("id")

// 	if err := h.r.Manager(r.Context()).Where("id = ?", id).First(&m).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	resp.NewResponse(w).Ok(&m)
// }

// func (h *Handler) DeleteCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	var id = ps.ByName("id")

// 	ID, err := uuid.FromString(id)
// 	if err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 	}

// 	if err := h.r.Manager(r.Context()).Delete(&models.Card{ID: ID}).Error; err != nil {
// 		resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 		return
// 	}

// 	resp.NewResponse(w).NoContent()
// }
