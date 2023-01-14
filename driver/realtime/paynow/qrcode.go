// Package paynow generates QR Code for Singapore PayNow payment system.
package paynow

import (
	"bytes"
	"time"

	qrcode "github.com/yeqown/go-qrcode"
)

// QRSpec to set qrcode version
type QRSpec struct {
	Version int
}

// QRCode struct, all fields are optional.
type QRCode struct {
	Payee
	Options []qrcode.ImageOption
	Spec    QRSpec // optional qrcode version
}

func NewUENQRCode(merchantName string, uen string) *QRCode {
	q := &QRCode{
		Payee: *NewUEN(merchantName, uen),
	}
	return q
}

func NewMobileQRCode(mobile string) *QRCode {
	q := &QRCode{
		Payee: *NewMobile(mobile),
	}
	return q
}

// QRCode returns editable qr code image
func (pp *QRCode) QRCode(amount float32, ref string) ([]byte, string, error) {
	return pp.QRCodeExpiry(amount, ref, true, time.Time{})
}

// QRCodeLocked returns uneditable qr code image
func (pp *QRCode) QRCodeLocked(amount float32, ref string) ([]byte, string, error) {
	return pp.QRCodeExpiry(amount, ref, false, time.Time{})
}

// QRCodeExpiry returns QR Code image, in jpeg format by default
// expirySGT should be in Asia/Singapore timezone. If expirySGT is zero, perpetual
func (pp *QRCode) QRCodeExpiry(amount float32, ref string, editable bool, expirySGT time.Time) ([]byte, string, error) {
	var (
		buf bytes.Buffer
		qrc *qrcode.QRCode
		err error
	)

	value := pp.New(amount, ref, editable, expirySGT).String()

	if pp.Spec.Version != 0 {
		qrc, err = qrcode.NewWithSpecV(value, pp.Spec.Version, qrcode.ErrorCorrectionMedium, pp.Options...)
	} else {
		qrc, err = qrcode.New(value, pp.Options...)
	}

	if err != nil {
		return nil, "", err
	}

	err = qrc.SaveTo(&buf)
	if err != nil {
		return nil, "", err
	}

	return buf.Bytes(), value, nil
}
