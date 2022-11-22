package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type Query struct {
	Where     string
	Order     string
	Select    string
	Offset    int
	Limit     int
	DeletedAt bool
}

func MapToQuery(i []*structs.Field) Query {
	var query Query
	for _, f := range i {
		name := strings.Replace(f.Tag("json"), ",omitempty", "", -1)
		if f.Kind() == reflect.Struct {
			if f.Name() == "Lt" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`< "%+v"`, f.Value())
				}
			} else if f.Name() == "Gt" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`> "%+v"`, f.Value())
				}
			} else if f.Name() == "Lte" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`<= "%+v"`, f.Value())
				}
			} else if f.Name() == "Gte" {
				if !f.IsZero() {
					query.Where += fmt.Sprintf(`>= "%+v"`, f.Value())
				}
			} else {
				if !f.IsZero() {
					if f.IsEmbedded() {
						query.Where = MapToQuery(f.Fields()).Where
					} else {
						if len(query.Where) > 0 {
							query.Where += " AND "
						}
						query.Where += fmt.Sprintf("%+v %+v", name, MapToQuery(f.Fields()))
					}

				}
			}
		} else if f.Kind() == reflect.String {
			if !f.IsZero() {
				if f.Name() == "Select" {
					query.Select = fmt.Sprintf(`%+v`, f.Value())
				} else if f.Name() == "Order" {
					query.Order = fmt.Sprintf(`%+v`, f.Value())
				} else {
					if len(query.Where) > 0 {
						query.Where += " AND "
					}
					query.Where += fmt.Sprintf(`%+v = "%+v"`, name, f.Value())
				}
			}
		} else if f.Kind() == reflect.Int {
			if !f.IsZero() {
				if f.Name() == "Offset" {
					query.Offset = f.Value().(int)
				} else if f.Name() == "Limit" {
					query.Limit = f.Value().(int)
				} else {
					if len(query.Where) > 0 {
						query.Where += " AND "
					}
					query.Where += fmt.Sprintf(`%+v = %+v`, name, f.Value())
				}
			}
		} else if f.Kind() == reflect.Bool {
			if !f.IsZero() {
				if f.Name() == "DeletedAt" {
					query.DeletedAt = f.Value().(bool)
				} else {
					if len(query.Where) > 0 {
						query.Where += " AND "
					}
					query.Where += fmt.Sprintf(`%+v = %+v`, name, f.Value())
				}
			}
		}
	}
	return query
}

func BuildQuery(model interface{}) Query {
	return MapToQuery(structs.Fields(model))
}
