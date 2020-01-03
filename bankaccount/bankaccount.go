package bankaccount

type BankAccount struct {
	AccountNumber  *string                    `json:"accountNumber,omitempty"`
	Annotations    map[string]string          `json:"annotations,omitempty"`
	Bic            *string                    `json:"bic,omitempty"`
	BillingAddress *BankAccountBillingAddress `json:"billingAddress,omitempty"`
	CountryISO     *string                    `json:"countryISO,omitempty"`
	CreatedAt      *string                    `json:"createdAt,omitempty"`
	CurrencyISO    *string                    `json:"currencyISO,omitempty"`
	Iban           *string                    `json:"iban,omitempty"`
	ID             *string                    `json:"id,omitempty"`
	OwnerID        *string                    `json:"ownerId,omitempty"`
	SortCode       *string                    `json:"sortCode,omitempty"`
	StateID        *BankAccountState          `json:"stateId,omitempty"`
	Tags           []string                   `json:"tags"`
	UpdatedAt      *string                    `json:"updatedAt,omitempty"`
}

type BankAccountBillingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type BankAccountList struct {
	Results []Result `json:"results"`
}

type Result struct {
	AccountNumber  *string               `json:"accountNumber,omitempty"`
	Annotations    map[string]string     `json:"annotations,omitempty"`
	Bic            *string               `json:"bic,omitempty"`
	BillingAddress *ResultBillingAddress `json:"billingAddress,omitempty"`
	CountryISO     *string               `json:"countryISO,omitempty"`
	CreatedAt      *string               `json:"createdAt,omitempty"`
	CurrencyISO    *string               `json:"currencyISO,omitempty"`
	Iban           *string               `json:"iban,omitempty"`
	ID             *string               `json:"id,omitempty"`
	OwnerID        *string               `json:"ownerId,omitempty"`
	SortCode       *string               `json:"sortCode,omitempty"`
	StateID        *BankAccountState     `json:"stateId,omitempty"`
	Tags           []string              `json:"tags"`
	UpdatedAt      *string               `json:"updatedAt,omitempty"`
}

type ResultBillingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type Create struct {
	AccountNumber  *string              `json:"accountNumber,omitempty"`
	Annotations    map[string]string    `json:"annotations,omitempty"`
	Bic            *string              `json:"bic,omitempty"`
	BillingAddress CreateBillingAddress `json:"billingAddress"`
	CountryISO     string               `json:"countryISO"`
	CurrencyISO    string               `json:"currencyISO"`
	Iban           *string              `json:"iban,omitempty"`
	OwnerID        string               `json:"ownerId"`
	SetAsPreferred *bool                `json:"setAsPreferred,omitempty"`
	SortCode       *string              `json:"sortCode,omitempty"`
	Tags           []string             `json:"tags"`
}

type CreateBillingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type Update struct {
	BillingAddress *UpdateBillingAddress `json:"billingAddress,omitempty"`
	ID             string                `json:"id"`
	SetAsPreferred *bool                 `json:"setAsPreferred,omitempty"`
}

type UpdateBillingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type Transition struct {
	FromStateID *BankAccountState `json:"fromStateId,omitempty"`
	ID          string            `json:"id"`
	ToStateID   BankAccountState  `json:"toStateId"`
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

type BankAccountState string

const (
	Closed   BankAccountState = "CLOSED"
	Disabled BankAccountState = "DISABLED"
	Enabled  BankAccountState = "ENABLED"
	Pending  BankAccountState = "PENDING"
)
