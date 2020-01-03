package cardholder

type Create struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Email       string            `json:"email"`
	FullName    string            `json:"fullName"`
	Tags        []string          `json:"tags"`
	Terms       CreateTerms       `json:"terms"`
	Type        *CardholderType   `json:"type,omitempty"`
	Role        *CreateRole       `json:"role,omitempty"`
	RoleID      *string           `json:"roleId,omitempty"`
	RoleSlug    *string           `json:"roleSlug,omitempty"`
}

type CreateTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type Transition struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	FromStateID *CardholderState  `json:"fromStateId,omitempty"`
	ID          string            `json:"id"`
	Tags        []string          `json:"tags"`
	ToStateID   CardholderState   `json:"toStateId"`
}

type CardholderList struct {
	Results []Result `json:"results"`
}

type Result struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Email       *string           `json:"email,omitempty"`
	FullName    *string           `json:"fullName,omitempty"`
	ID          *string           `json:"id,omitempty"`
	RoleID      *string           `json:"roleId,omitempty"`
	StateID     *CardholderState  `json:"stateId,omitempty"`
	Tags        []string          `json:"tags"`
	Terms       *ResultTerms      `json:"terms,omitempty"`
	Type        *CardholderType   `json:"type,omitempty"`
}

type ResultTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type Cardholder struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Email       *string           `json:"email,omitempty"`
	FullName    *string           `json:"fullName,omitempty"`
	ID          *string           `json:"id,omitempty"`
	RoleID      *string           `json:"roleId,omitempty"`
	StateID     *CardholderState  `json:"stateId,omitempty"`
	Tags        []string          `json:"tags"`
	Terms       *CardholderTerms  `json:"terms,omitempty"`
	Type        *CardholderType   `json:"type,omitempty"`
}

type CardholderTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CardholderToken struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

type CreateToken struct {
	ID string `json:"id"`
}

type UpdateCardholder struct {
	Email    *string                `json:"email,omitempty"`
	FullName *string                `json:"fullName,omitempty"`
	ID       string                 `json:"id"`
	Terms    *UpdateCardholderTerms `json:"terms,omitempty"`
}

type UpdateCardholderTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type RoleID struct {
	Role   *RoleIDRole `json:"role,omitempty"`
	RoleID string      `json:"roleId"`
}

type RoleSlug struct {
	Role     *RoleSlugRole `json:"role,omitempty"`
	RoleSlug string        `json:"roleSlug"`
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

type CreateRole string

const (
	PurpleRoleID   CreateRole = "roleId"
	PurpleRoleSlug CreateRole = "roleSlug"
)

type CardholderType string

const (
	Company    CardholderType = "COMPANY"
	Individual CardholderType = "INDIVIDUAL"
)

type CardholderState string

const (
	Closed   CardholderState = "CLOSED"
	Disabled CardholderState = "DISABLED"
	Enabled  CardholderState = "ENABLED"
	Pending  CardholderState = "PENDING"
)

type RoleIDRole string

const (
	FluffyRoleID RoleIDRole = "roleId"
)

type RoleSlugRole string

const (
	FluffyRoleSlug RoleSlugRole = "roleSlug"
)
