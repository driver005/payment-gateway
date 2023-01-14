package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/helper"
)

type BatchJobResultErrorsCode struct {
}

type BatchJobResultErrors struct {
	Message string `json:"message,omitempty" database:"default:null"`

	Code BatchJobResultErrorsCode `json:"code,omitempty" database:"default:null"`

	Err []string `json:"err,omitempty" database:"default:null"`
}

type BatchJobResultStatDescriptors struct {
	Key string `json:"key,omitempty" database:"default:null"`

	Name string `json:"name,omitempty" database:"default:null"`

	Message string `json:"message,omitempty" database:"default:null"`
}

// BatchJobResult - The result of the batch job.
type BatchJobResult struct {
	Count float64 `json:"count,omitempty" database:"default:null"`

	AdvancementCount float64 `json:"advancement_count,omitempty" database:"default:null"`

	Progress float64 `json:"progress,omitempty" database:"default:null"`

	Errors BatchJobResultErrors `json:"errors,omitempty" database:"default:null"`

	StatDescriptors BatchJobResultStatDescriptors `json:"stat_descriptors,omitempty" database:"default:null"`

	FileKey string `json:"file_key,omitempty" database:"default:null"`

	FileSize float64 `json:"file_size,omitempty" database:"default:null"`
}

func (b BatchJobResult) Interface() interface{} {
	return BatchJobResult(b)
}

func (BatchJobResult) DBDataType() string {
	return "result"
}

func (b BatchJobResult) DBValue(ctx context.Context, db *database.DB) clause.Expr {
	return clause.Expr{
		SQL: "ROW(?, ?, ?, ?, ?, ?, ?, ?)::result",
		Vars: []interface{}{
			b.Count,
			b.AdvancementCount,
			b.Progress,
			b.Errors,
			b.StatDescriptors,
			b.FileKey,
			b.FileSize,
		},
	}
}

func (b *BatchJobResult) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	s := strings.Split(strings.Trim(string(bytes), "()"), ",")
	for i := 0; i < reflect.TypeOf(BatchJobResult{}).NumField(); i++ {
		helper.SetField(&b, reflect.TypeOf(BatchJobResult{}).Field(i).Name, s[i])
	}

	return nil
}
