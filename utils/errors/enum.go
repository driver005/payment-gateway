package errors

import "database/sql/driver"

type Type string

const (
	TypeApiError            Type = "api_error"
	TypeCardError           Type = "card_error"
	TypeIdempotencyError    Type = "idempotency_error"
	TypeInvalidRequestError Type = "invalid_request_error"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}
