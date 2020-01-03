package kyc

type Kyc struct {
	ReferenceID string `json:"referenceId"`
}

type Create struct {
	CustomerID string     `json:"customerId"`
	Kyc        *CreateKyc `json:"kyc,omitempty"`
	Type       *Type      `json:"type,omitempty"`
}

type CreateKyc struct {
	ReferenceID string `json:"referenceId"`
}

type CheckEvent struct {
	CustomerID string                 `json:"customerId"`
	Payload    map[string]interface{} `json:"payload"`
	Type       string                 `json:"type"`
}

type Check struct {
	CustomerID *string   `json:"customerId,omitempty"`
	ID         *string   `json:"id,omitempty"`
	Kyc        *CheckKyc `json:"kyc,omitempty"`
	Type       *Type     `json:"type,omitempty"`
}

type CheckKyc struct {
	ReferenceID string `json:"referenceId"`
}

type PurpleKyc struct {
	Kyc  KycKyc `json:"kyc"`
	Type *Type  `json:"type,omitempty"`
}

type KycKyc struct {
	ReferenceID string `json:"referenceId"`
}

type Purpose string

const (
	Deposit          Purpose = "DEPOSIT"
	ManualAdjustment Purpose = "MANUAL_ADJUSTMENT"
	Purchase         Purpose = "PURCHASE"
	Transfer         Purpose = "TRANSFER"
	Withdrawal       Purpose = "WITHDRAWAL"
)

type Flow string

const (
	Inbound  Flow = "INBOUND"
	Internal Flow = "INTERNAL"
	Outbound Flow = "OUTBOUND"
)

type Type string

const (
	TypeKyc Type = "kyc"
)
