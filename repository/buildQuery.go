package repository

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/database"
	"github.com/fatih/structs"
)

type Result struct {
	Where   string
	Order   string
	Select  string
	Offset  int
	Limit   int
	Deleted bool
}

func MapToQuery(i []*structs.Field) Result {
	var query Result
	for _, f := range i {
		// fmt.Printf("%+v, %+v, %+v \n", f.Name(), f.Kind(), f.Value())
		name := strings.Replace(f.Tag("json"), ",omitempty", "", -1)
		if f.Kind() == reflect.Struct {
			if !f.IsZero() {
				if f.Name() == "N" {
					query.Where += fmt.Sprintf(`= '%+v'`, f.Value().(time.Time).Format(time.RFC3339))
				} else if f.Name() == "Lt" {
					query.Where += fmt.Sprintf(`< '%+v'`, f.Value().(time.Time).Format(time.RFC3339))
				} else if f.Name() == "Gt" {
					query.Where += fmt.Sprintf(`> '%+v'`, f.Value().(time.Time).Format(time.RFC3339))
				} else if f.Name() == "Lte" {
					query.Where += fmt.Sprintf(`<= '%+v'`, f.Value().(time.Time).Format(time.RFC3339))
				} else if f.Name() == "Gte" {
					query.Where += fmt.Sprintf(`>= '%+v'`, f.Value().(time.Time).Format(time.RFC3339))
				} else {
					if f.IsEmbedded() {
						val := MapToQuery(f.Fields())
						query = Result{
							Where:   val.Where,
							Order:   val.Order,
							Select:  val.Select,
							Offset:  val.Offset,
							Limit:   val.Limit,
							Deleted: val.Deleted,
						}
					} else {
						if len(query.Where) > 0 {
							query.Where += " AND "
						}
						query.Where += fmt.Sprintf("%+v %+v", name, MapToQuery(f.Fields()).Where)
					}
				}
			}
		} else if f.Kind() == reflect.String {
			if !f.IsZero() {
				if f.Name() == "N" {
					query.Where += fmt.Sprintf(`= '%+v'`, f.Value().(string))
				} else if f.Name() == "Lt" {
					query.Where += fmt.Sprintf(`< '%+v'`, f.Value().(string))
				} else if f.Name() == "Gt" {
					query.Where += fmt.Sprintf(`> '%+v'`, f.Value().(string))
				} else if f.Name() == "Lte" {
					query.Where += fmt.Sprintf(`<= '%+v'`, f.Value().(string))
				} else if f.Name() == "Gte" {
					query.Where += fmt.Sprintf(`>= '%+v'`, f.Value().(string))
				} else if f.Name() == "Select" {
					query.Select = f.Value().(string)
				} else if f.Name() == "Order" {
					query.Order = f.Value().(string)
				} else {
					if len(query.Where) > 0 {
						query.Where += " AND "
					}
					query.Where += fmt.Sprintf(`%+v = '%+v'`, name, f.Value())
				}
			}
		} else if f.Kind() == reflect.Int {
			if f.Name() == "N" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`= '%+v'`, f.Value().(int))
				}
			} else if f.Name() == "Lt" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`< '%+v'`, f.Value().(int))
				}
			} else if f.Name() == "Gt" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`> '%+v'`, f.Value().(int))
				}
			} else if f.Name() == "Lte" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`<= '%+v'`, f.Value().(int))
				}
			} else if f.Name() == "Gte" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`>= '%+v'`, f.Value().(int))
				}
			} else if f.Name() == "Offset" {
				if !f.IsZero() {
					query.Offset = f.Value().(int)
				} else {
					query.Offset = -1
				}
			} else if f.Name() == "Limit" {
				if !f.IsZero() {
					query.Limit = f.Value().(int)
				} else {
					query.Limit = -1
				}
			} else {
				if !f.IsZero() {
					if len(query.Where) > 0 {
						query.Where += " AND "
					}
					query.Where += fmt.Sprintf(`%+v = %+v`, name, f.Value())
				}
			}
		} else if f.Kind() == reflect.Bool {
			if !f.IsZero() {
				if f.Name() == "N" {
					query.Where += fmt.Sprintf(`= '%+v'`, f.Value().(bool))
				} else if f.Name() == "Lt" {
					query.Where += fmt.Sprintf(`< '%+v'`, f.Value().(bool))
				} else if f.Name() == "Gt" {
					query.Where += fmt.Sprintf(`> '%+v'`, f.Value().(bool))
				} else if f.Name() == "Lte" {
					query.Where += fmt.Sprintf(`<= '%+v'`, f.Value().(bool))
				} else if f.Name() == "Gte" {
					query.Where += fmt.Sprintf(`>= '%+v'`, f.Value().(bool))
				} else if f.Name() == "Deleted" {
					query.Deleted = f.Value().(bool)
				} else {
					if len(query.Where) > 0 {
						query.Where += " AND "
					}
					query.Where += fmt.Sprintf(`%+v = %+v`, name, f.Value())
				}
			}
		} else if f.Kind() == reflect.Slice {
			if !f.IsZero() {
				if len(query.Where) > 0 {
					query.Where += " AND "
				}
				query.Where += fmt.Sprintf(`%+v = %+v`, name, f.Value())
			}
		} else if f.Kind() == reflect.Array {
			if !f.IsZero() {
				if len(query.Where) > 0 {
					query.Where += " AND "
				}
				query.Where += fmt.Sprintf(`%+v = %+v`, name, f.Value())
			}
		}
	}
	return query
}

func BuildQuery(model interface{}) Result {
	return MapToQuery(structs.Fields(model))
}

func (r *Repositories) ToQuery(config interface{}, model interface{}) *database.DB {
	s := BuildQuery(config)
	m := BuildQuery(model)
	query := r.DB()

	if m.Where != "" {
		if len(s.Where) > 0 {
			s.Where += " AND "
		}
		s.Where += fmt.Sprintf(`%+v`, m.Where)
	}

	if s.Deleted {
		query = query.Unscoped()
	}
	if s.Select != "" {
		query = query.Select(s.Select)
	}
	if s.Where != "" {
		query = query.Where(s.Where)
	}
	if s.Limit != -1 {
		query = query.Limit(s.Limit)
	}
	if s.Offset != -1 {
		query = query.Offset(s.Offset)
	}
	if s.Order != "" {
		query = query.Order(s.Order)
	}

	return query
}

type FindOption struct {
	Select    map[string]interface{}
	Where     map[string]interface{}
	Order     map[string]interface{}
	Relations map[string]interface{}
	Offset    *int
	Limit     *int
	Deleted   bool
}

func (r *Repositories) BuildFindQuery(option FindOption) *database.DB {
	query := r.db
	for index, value := range option.Select {
		switch value.(type) {
		case bool:
			if value == true {
				query = query.Select(index)
			}
		}
	}

	for index, value := range option.Relations {
		switch value.(type) {
		case bool:
			if value == true {
				query = query.Preload(index)
			}
		}
	}

	for index, value := range option.Where {
		switch value.(type) {
		case string:
			if value == true {
				query = query.Where(fmt.Sprintf("%+v = ?", index), value)
			}
		}
	}

	for index, value := range option.Order {
		switch value.(type) {
		case string:
			if value == true {
				query = query.Order(fmt.Sprintf("%+v %+v", index, value))
			}
		}
	}

	if option.Offset != nil {
		query = query.Offset(*option.Offset)
	}

	if option.Limit != nil {
		query = query.Limit(*option.Limit)
	}

	if option.Deleted {
		query = query.Unscoped()
	}

	return query
}

func (r *Repositories) BuildFindWhereQuery(where map[string]interface{}) *database.DB {
	query := r.db

	for index, value := range where {
		switch value.(type) {
		case string:
			if value == true {
				query = query.Where(fmt.Sprintf("%+v = ?", index), value)
			}
		}
	}

	return query
}
