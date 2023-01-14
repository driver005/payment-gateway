package wechat

import (
	"context"
	"net/http"
)

// CertificatesRequest is the request for certificates.
type CertificatesRequest struct {
}

// CertificatesResponse is the response for certificates.
type CertificatesResponse struct {
	Certificates []Certificate `json:"data"`
}

// Certificate is certificate information
type Certificate struct {
	SerialNo      string             `json:"serial_no"`
	EffectiveTime string             `json:"effective_time"`
	ExpireTime    string             `json:"expire_time"`
	Encrypt       EncryptCertificate `json:"encrypt_certificate"`
}

// EncryptCertificate is the information of encrypt certificate.
type EncryptCertificate struct {
	Algorithm  string `json:"algorithm"`
	Nonce      string `json:"nonce"`
	Associated string `json:"associated_data"`
	CipherText string `json:"ciphertext"`
}

// Do get certificates from wechat pay.
func (r *CertificatesRequest) Do(ctx context.Context, c Client) (*CertificatesResponse, error) {
	url := c.Config().Options().CertUrl

	resp := &CertificatesResponse{}
	if err := c.Do(ctx, http.MethodGet, url).Scan(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
