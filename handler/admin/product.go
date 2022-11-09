package admin

import (
	"net/http"
	"strconv"

	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/julienschmidt/httprouter"
)

type Product struct {
	r handler.Registry
}

func (p *Product) ListProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m = make([]models.Product, 0)
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
	if err := p.r.Manager(r.Context()).Scopes(Paginate(r, offset, pageSize)).Order("id").Find(&m).Error; err != nil {
		p.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n := p.r.Manager(r.Context()).Find(&models.Product{})
	if n.Error != nil {
		p.r.Writer().WriteError(w, r, helper.WithStack(n.Error))
		return
	}

	helper.Header(w, r.URL, n.RowsAffected, pageSize, offset)

	if m == nil {
		m = []models.Product{}
	}

	p.r.Writer().Write(w, r, &m)
}
