package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/driver005/database"
	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
)

type Currency struct {
	r handler.Registry
}

func Paginate(r *http.Request, offset int, pageSize int) func(db *database.DB) *database.DB {
	return func(db *database.DB) *database.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func (c *Currency) GetCurrency(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m models.Currency
	var id = ps.ByName("id")

	if err := c.r.Manager(r.Context()).Where("id = ?", id).First(&m).Error; err != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	c.r.Writer().Write(w, r, &m)
}

func (c *Currency) ListCurrency(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m = make([]models.Currency, 0)
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	if err := c.r.Manager(r.Context()).Scopes(Paginate(r, offset, pageSize)).Order("id").Find(&m).Error; err != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n := c.r.Manager(r.Context()).Find(&models.Currency{})
	if n.Error != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(n.Error))
		return
	}

	helper.Header(w, r.URL, n.RowsAffected, pageSize, offset)

	if m == nil {
		m = []models.Currency{}
	}

	c.r.Writer().Write(w, r, &m)
}

func (c *Currency) UpdateCurrency(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m models.Currency
	var o models.Currency
	var err error

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	m.Id, err = uuid.FromString(ps.ByName("id"))
	if err != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(err))
	}

	if err = c.r.Manager(r.Context()).Where("id = ?", m.Id).First(&o).Error; err != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := c.r.Manager(r.Context()).Model(&o).Updates(m).Error; err != nil {
		c.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	c.r.Writer().Write(w, r, &m)
}
