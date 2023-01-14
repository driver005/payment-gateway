package iso3166_test

import (
	"testing"

	"github.com/driver005/gateway/driver/debit/ach/internal/iso3166"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert.True(t, iso3166.Valid("US"))
	assert.True(t, iso3166.Valid("SS"))
	assert.False(t, iso3166.Valid(""))
	assert.False(t, iso3166.Valid("QZ"))
}
