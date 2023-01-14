package iso4217_test

import (
	"testing"

	"github.com/driver005/gateway/driver/debit/ach/internal/iso4217"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert.True(t, iso4217.Valid("USD"))
	assert.True(t, iso4217.Valid("eur"))
	assert.False(t, iso4217.Valid(""))
	assert.False(t, iso4217.Valid("QZA"))
}
