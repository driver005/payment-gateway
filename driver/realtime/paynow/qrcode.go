// Package paynow generates QR Code for Singapore PayNow payment system.
package paynow

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// QRCodeOptions is a configuration struct.
type QRCodeOptions struct {
	// QR Code content
	Content string
	// Default: 256
	Size int
}

// QRCode returns a graphical representation of the Copy and Paste code in a QR Code form.
func QRCode(options QRCodeOptions) (barcode.Barcode, error) {
	if options.Size == 0 {
		options.Size = 256
	}

	bytes, err := qr.Encode(options.Content, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	bytes, err = barcode.Scale(bytes, options.Size, options.Size)

	if err != nil {
		return nil, err
	}

	return bytes, err
}
