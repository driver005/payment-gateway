package types

import "github.com/driver005/gateway/core"

// SourceTypeCardPresent struct for SourceTypeCardPresent
type SourceTypeCardPresent struct {
	core.Model

	ApplicationCryptogram          string `json:"application_cryptogram,omitempty"`
	ApplicationPreferredName       string `json:"application_preferred_name,omitempty"`
	AuthorizationCode              string `json:"authorization_code,omitempty"`
	AuthorizationResponseCode      string `json:"authorization_response_code,omitempty"`
	Brand                          string `json:"brand,omitempty"`
	Country                        string `json:"country,omitempty"`
	CvmType                        string `json:"cvm_type,omitempty"`
	DataType                       string `json:"data_type,omitempty"`
	DedicatedFileName              string `json:"dedicated_file_name,omitempty"`
	EmvAuthData                    string `json:"emv_auth_data,omitempty"`
	EvidenceCustomerSignature      string `json:"evidence_customer_signature,omitempty"`
	EvidenceTransactionCertificate string `json:"evidence_transaction_certificate,omitempty"`
	ExpMonth                       int    `json:"exp_month,omitempty"`
	ExpYear                        int    `json:"exp_year,omitempty"`
	Fingerprint                    string `json:"fingerprint,omitempty"`
	Funding                        string `json:"funding,omitempty"`
	Last4                          string `json:"last4,omitempty"`
	PosDeviceId                    string `json:"pos_device_id,omitempty"`
	PosEntryMode                   string `json:"pos_entry_mode,omitempty"`
	ReadMethod                     string `json:"read_method,omitempty"`
	Reader                         string `json:"reader,omitempty"`
	TerminalVerificationResults    string `json:"terminal_verification_results,omitempty"`
	TransactionStatusInformation   string `json:"transaction_status_information,omitempty"`
}
