package card

type Card struct {
	Annotations      map[string]string   `json:"annotations,omitempty"`
	BillingAddress   *CardBillingAddress `json:"billingAddress,omitempty"`
	CardNumberSuffix *string             `json:"cardNumberSuffix,omitempty"`
	CreatedAt        *string             `json:"createdAt,omitempty"`
	Expiry           *string             `json:"expiry,omitempty"`
	ID               *string             `json:"id,omitempty"`
	OwnerID          *string             `json:"ownerId,omitempty"`
	Requirements     map[string]string   `json:"requirements,omitempty"`
	Scheme           *string             `json:"scheme,omitempty"`
	SetAsPreferred   *bool               `json:"setAsPreferred,omitempty"`
	StateID          *CardState          `json:"stateId,omitempty"`
	StorageType      *StorageType        `json:"storageType,omitempty"`
	Tags             []string            `json:"tags"`
	Type             *string             `json:"type,omitempty"`
	UpdatedAt        *string             `json:"updatedAt,omitempty"`
}

type CardBillingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CardList struct {
	Results []Result `json:"results"`
}

type Result struct {
	Annotations      map[string]string     `json:"annotations,omitempty"`
	BillingAddress   *ResultBillingAddress `json:"billingAddress,omitempty"`
	CardNumberSuffix *string               `json:"cardNumberSuffix,omitempty"`
	CreatedAt        *string               `json:"createdAt,omitempty"`
	Expiry           *string               `json:"expiry,omitempty"`
	ID               *string               `json:"id,omitempty"`
	OwnerID          *string               `json:"ownerId,omitempty"`
	Requirements     map[string]string     `json:"requirements,omitempty"`
	Scheme           *string               `json:"scheme,omitempty"`
	SetAsPreferred   *bool                 `json:"setAsPreferred,omitempty"`
	StateID          *CardState            `json:"stateId,omitempty"`
	StorageType      *StorageType          `json:"storageType,omitempty"`
	Tags             []string              `json:"tags"`
	Type             *string               `json:"type,omitempty"`
	UpdatedAt        *string               `json:"updatedAt,omitempty"`
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
	Annotations    map[string]string    `json:"annotations,omitempty"`
	BillingAddress CreateBillingAddress `json:"billingAddress"`
	CardNumber     string               `json:"cardNumber"`
	Cvv            string               `json:"cvv"`
	Expiry         string               `json:"expiry"`
	OwnerID        string               `json:"ownerId"`
	StorageType    *StorageType         `json:"storageType,omitempty"`
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
	BillingAddress UpdateBillingAddress `json:"billingAddress"`
	ID             string               `json:"id"`
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
	Annotations map[string]string `json:"annotations,omitempty"`
	FromStateID *CardState        `json:"fromStateId,omitempty"`
	ID          string            `json:"id"`
	Tags        []string          `json:"tags"`
	ToStateID   CardState         `json:"toStateId"`
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

type CardState string

const (
	Closed     CardState = "CLOSED"
	Disabled   CardState = "DISABLED"
	Enabled    CardState = "ENABLED"
	Failed     CardState = "FAILED"
	Inactive   CardState = "INACTIVE"
	Pending    CardState = "PENDING"
	Processing CardState = "PROCESSING"
)

type StorageType string

const (
	CardOnFile StorageType = "CARD_ON_FILE"
	OneTimeUse StorageType = "ONE_TIME_USE"
)
