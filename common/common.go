package common

type Address struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type Get struct {
	ID string `json:"id"`
}

type List struct {
	ID    []string `json:"id"`
	Query *string  `json:"query,omitempty"`
}

type Tag struct {
	ID   string   `json:"id"`
	Tags []string `json:"tags"`
}

type Untag struct {
	ID   string   `json:"id"`
	Tags []string `json:"tags"`
}

type Annotate struct {
	Annotations map[string]string `json:"annotations"`
	ID          string            `json:"id"`
}

type RemoveAnnotations struct {
	ID   string   `json:"id"`
	Keys []string `json:"keys"`
}

type Terms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type PreferredInstrument struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type PreferredInstruments struct {
	BankAccount BankAccount `json:"bankAccount"`
	Card        Card        `json:"card"`
}

type BankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Card struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Participants struct {
	Participant1 string `json:"participant1"`
	Participant2 string `json:"participant2"`
	Participant3 string `json:"participant3"`
	Participant4 string `json:"participant4"`
	Participant5 string `json:"participant5"`
	Participant6 string `json:"participant6"`
	Prefix       string `json:"prefix"`
}

type Endpoint struct {
	Authenticate bool `json:"authenticate"`
	Idempotent   bool `json:"idempotent"`
	Queryable    bool `json:"queryable"`
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
