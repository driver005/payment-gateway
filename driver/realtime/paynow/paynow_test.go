package paynow

import (
	"testing"
	"time"
)

func TestPayNowUEN(t *testing.T) {
	want := "00020101021226520009SG.PAYNOW010120213S99345678ABCD03011040899991231520400005303702540512.345802SG5912ACME Pte Ltd6009Singapore62110107INV123463041E74"
	payee := NewUEN("ACME Pte Ltd", "S99345678ABCD")
	got := payee.New(12.34, "INV1234", true, time.Time{}).String()

	if want != got {
		t.Fatal("want", want, "got", got)
	}
	t.Log(got)
}

func TestPayNowMobile(t *testing.T) {
	want := "00020101021226500009SG.PAYNOW010100211+659099123403011040899991231520400005303702540512.345802SG5902NA6009Singapore62110107INV123463048AF5"
	payee := NewMobile("+6590991234")
	got := payee.New(12.34, "INV1234", true, time.Time{}).String()

	if want != got {
		t.Fatal("want", want, "got", got)
	}
	t.Log(got)
}

func BenchmarkPN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payee := NewUEN("ACME Pte Ltd", "S99345678ABCD")
		_ = payee.New(12.34, "INV1234", true, time.Time{}).String()
	}
}
