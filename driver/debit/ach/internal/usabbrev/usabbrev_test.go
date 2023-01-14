package usabbrev

import (
	"testing"
)

func TestUSAbbrev(t *testing.T) {
	if !Valid("WA") {
		t.Error("WA (Washington) is a sate")
	}
	if !Valid("GU") {
		t.Error("GU (Guam) is a territory")
	}
	if Valid("XX") {
		t.Error("XX is not a valid US state or territory")
	}
}
