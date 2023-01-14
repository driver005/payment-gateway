package ach

import (
	"testing"
)

func TestOffsetAccountType(t *testing.T) {
	if err := OffsetChecking.validate(); err != nil {
		t.Fatal(err)
	}
	if err := OffsetSavings.validate(); err != nil {
		t.Fatal(err)
	}
	if err := OffsetAccountType("invalid").validate(); err == nil {
		t.Error("expected error")
	}
}
