package account

type Account struct {
	ActiveAt          *string           `json:"activeAt,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	AvailableBalance  *string           `json:"availableBalance,omitempty"`
	CompleteBalance   *string           `json:"completeBalance,omitempty"`
	CreatedAt         *string           `json:"createdAt,omitempty"`
	CurrencyISO       *string           `json:"currencyISO,omitempty"`
	DisputedCredits   *string           `json:"disputedCredits,omitempty"`
	DisputedDebits    *string           `json:"disputedDebits,omitempty"`
	ID                *string           `json:"id,omitempty"`
	IncompleteCredits *string           `json:"incompleteCredits,omitempty"`
	IncompleteDebits  *string           `json:"incompleteDebits,omitempty"`
	OwnerID           *string           `json:"ownerId,omitempty"`
	StateID           *AccountState     `json:"stateId,omitempty"`
	Tags              []string          `json:"tags"`
	Type              *AccountType      `json:"type,omitempty"`
	UpdatedAt         *string           `json:"updatedAt,omitempty"`
	WithheldBalance   *string           `json:"withheldBalance,omitempty"`
}

type AccountList struct {
	Results []Result `json:"results"`
}

type Result struct {
	ActiveAt          *string           `json:"activeAt,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	AvailableBalance  *string           `json:"availableBalance,omitempty"`
	CompleteBalance   *string           `json:"completeBalance,omitempty"`
	CreatedAt         *string           `json:"createdAt,omitempty"`
	CurrencyISO       *string           `json:"currencyISO,omitempty"`
	DisputedCredits   *string           `json:"disputedCredits,omitempty"`
	DisputedDebits    *string           `json:"disputedDebits,omitempty"`
	ID                *string           `json:"id,omitempty"`
	IncompleteCredits *string           `json:"incompleteCredits,omitempty"`
	IncompleteDebits  *string           `json:"incompleteDebits,omitempty"`
	OwnerID           *string           `json:"ownerId,omitempty"`
	StateID           *AccountState     `json:"stateId,omitempty"`
	Tags              []string          `json:"tags"`
	Type              *AccountType      `json:"type,omitempty"`
	UpdatedAt         *string           `json:"updatedAt,omitempty"`
	WithheldBalance   *string           `json:"withheldBalance,omitempty"`
}

type Create struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	CurrencyISO string            `json:"currencyISO"`
	OwnerID     string            `json:"ownerId"`
	Tags        []string          `json:"tags"`
}

type Transition struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	FromStateID *AccountState     `json:"fromStateId,omitempty"`
	ID          string            `json:"id"`
	Tags        []string          `json:"tags"`
	ToStateID   AccountState      `json:"toStateId"`
}

type AccountState string

const (
	Closed   AccountState = "CLOSED"
	Disabled AccountState = "DISABLED"
	Enabled  AccountState = "ENABLED"
	Pending  AccountState = "PENDING"
)

type AccountType string

const (
	Fiat    AccountType = "FIAT"
	Loyalty AccountType = "LOYALTY"
)
