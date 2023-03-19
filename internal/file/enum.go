package file

import "database/sql/driver"

type Purpose string

const (
	PurposeAccountRequirement               Purpose = "account_requirement"
	PurposeAdditionalVerification           Purpose = "additional_verification"
	PurposeBusinessIcon                     Purpose = "business_icon"
	PurposeBusinesslogo                     Purpose = "business_logo"
	PurposeCustomerSignature                Purpose = "customer_signature"
	PurposeDisputeEvidence                  Purpose = "dispute_evidence"
	PurposeDocumentProviderIdentityDocument Purpose = "document_provider_identity_document"
	PurposeFinanceReportRun                 Purpose = "finance_report_run"
	PurposeIdentityDocument                 Purpose = "identity_document"
	PurposeIdentityDocumentDownloadable     Purpose = "identity_document_downloadable"
	PurposePciDocument                      Purpose = "pci_document"
	PurposeSelfie                           Purpose = "selfie"
	PurposeSigmaScheduledQuery              Purpose = "sigma_scheduled_query"
	PurposeTaxDocumentUserUpload            Purpose = "tax_document_user_upload"
	PurposeTerminalReaderSplashscreen       Purpose = "terminal_reader_splashscreen"
)

func (ct *Purpose) Scan(value interface{}) error {
	*ct = Purpose(value.(string))
	return nil
}

func (ct Purpose) Value() (driver.Value, error) {
	return string(ct), nil
}
