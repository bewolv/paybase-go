package document

type Type struct {
	Subtype string       `json:"subtype"`
	Type    DocumentType `json:"type"`
}

type Create struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	CustomerID  string            `json:"customerId"`
	FileName    string            `json:"fileName"`
	Tags        []string          `json:"tags"`
	Type        CreateType        `json:"type"`
}

type CreateType struct {
	Subtype string       `json:"subtype"`
	Type    DocumentType `json:"type"`
}

type Document struct {
	Annotations map[string]string  `json:"annotations,omitempty"`
	CreatedAt   *string            `json:"createdAt,omitempty"`
	CustomerID  *string            `json:"customerId,omitempty"`
	FileName    *string            `json:"fileName,omitempty"`
	ID          *string            `json:"id,omitempty"`
	StateID     *DocumentState     `json:"stateId,omitempty"`
	Tags        []string           `json:"tags"`
	Type        *DocumentTypeClass `json:"type,omitempty"`
	UpdatedAt   *string            `json:"updatedAt,omitempty"`
	Document    []byte             `json:document,omitempty`
}

type DocumentTypeClass struct {
	Subtype string       `json:"subtype"`
	Type    DocumentType `json:"type"`
}

type Types struct {
	Types map[string]interface{} `json:"types,omitempty"`
}

type DocumentType string

const (
	AdditionalInformation             DocumentType = "ADDITIONAL_INFORMATION"
	AdditionalInformationVerification DocumentType = "ADDITIONAL_INFORMATION_VERIFICATION"
	AddressVerification               DocumentType = "ADDRESS_VERIFICATION"
	BusinessPurpose                   DocumentType = "BUSINESS_PURPOSE"
	Compliance                        DocumentType = "COMPLIANCE"
	IDDocument                        DocumentType = "ID_DOCUMENT"
	Industry                          DocumentType = "INDUSTRY"
	Occupation                        DocumentType = "OCCUPATION"
	Other                             DocumentType = "OTHER"
	PurposeOfAccount                  DocumentType = "PURPOSE_OF_ACCOUNT"
	SourceOfFunds                     DocumentType = "SOURCE_OF_FUNDS"
	SourceOfWealth                    DocumentType = "SOURCE_OF_WEALTH"
)

type DocumentState string

const (
	Invalid DocumentState = "INVALID"
	Pending DocumentState = "PENDING"
	Valid   DocumentState = "VALID"
)
