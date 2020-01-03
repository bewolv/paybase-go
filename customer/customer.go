package customer

import (
	"time"
)

type BusinessPersonProfile struct {
	Dob                string                                  `json:"dob"`
	FirstName          string                                  `json:"firstName"`
	LastName           string                                  `json:"lastName"`
	ResidentialAddress BusinessPersonProfileResidentialAddress `json:"residentialAddress"`
	Roles              []string                                `json:"roles"`
}

type BusinessPersonProfileResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateBusinessPersonProfile struct {
	Dob                *string                                        `json:"dob,omitempty"`
	FirstName          *string                                        `json:"firstName,omitempty"`
	LastName           *string                                        `json:"lastName,omitempty"`
	ResidentialAddress *UpdateBusinessPersonProfileResidentialAddress `json:"residentialAddress,omitempty"`
	Roles              []string                                       `json:"roles"`
}

type UpdateBusinessPersonProfileResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type BusinessPerson struct {
	CustomerID *string                     `json:"customerId,omitempty"`
	ID         *string                     `json:"id,omitempty"`
	Profile    *BusinessPersonProfileClass `json:"profile,omitempty"`
}

type BusinessPersonProfileClass struct {
	Dob                string                   `json:"dob"`
	FirstName          string                   `json:"firstName"`
	LastName           string                   `json:"lastName"`
	ResidentialAddress PurpleResidentialAddress `json:"residentialAddress"`
	Roles              []string                 `json:"roles"`
}

type PurpleResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CreateBusinessPerson struct {
	CustomerID string                      `json:"customerId"`
	Profile    CreateBusinessPersonProfile `json:"profile"`
}

type CreateBusinessPersonProfile struct {
	Dob                string                   `json:"dob"`
	FirstName          string                   `json:"firstName"`
	LastName           string                   `json:"lastName"`
	ResidentialAddress FluffyResidentialAddress `json:"residentialAddress"`
	Roles              []string                 `json:"roles"`
}

type FluffyResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateBusinessPerson struct {
	CustomerID string                           `json:"customerId"`
	ID         string                           `json:"id"`
	Profile    UpdateBusinessPersonProfileClass `json:"profile"`
}

type UpdateBusinessPersonProfileClass struct {
	Dob                *string                      `json:"dob,omitempty"`
	FirstName          *string                      `json:"firstName,omitempty"`
	LastName           *string                      `json:"lastName,omitempty"`
	ResidentialAddress *TentacledResidentialAddress `json:"residentialAddress,omitempty"`
	Roles              []string                     `json:"roles"`
}

type TentacledResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IncorporatedBusinessProfile struct {
	Email             string                                       `json:"email"`
	IncorporatedDate  *string                                      `json:"incorporatedDate,omitempty"`
	Industry          *string                                      `json:"industry,omitempty"`
	OrgChart          *string                                      `json:"orgChart,omitempty"`
	PhoneNumber       string                                       `json:"phoneNumber"`
	RegisteredAddress IncorporatedBusinessProfileRegisteredAddress `json:"registeredAddress"`
	RegisteredName    string                                       `json:"registeredName"`
	TaxCountryISO     string                                       `json:"taxCountryISO"`
	TradingAddress    IncorporatedBusinessProfileTradingAddress    `json:"tradingAddress"`
	TradingName       *string                                      `json:"tradingName,omitempty"`
	URL               *string                                      `json:"url,omitempty"`
}

type IncorporatedBusinessProfileRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IncorporatedBusinessProfileTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateIncorporatedBusinessProfile struct {
	Email             *string                                             `json:"email,omitempty"`
	IncorporatedDate  *string                                             `json:"incorporatedDate,omitempty"`
	Industry          *string                                             `json:"industry,omitempty"`
	OrgChart          *string                                             `json:"orgChart,omitempty"`
	PhoneNumber       *string                                             `json:"phoneNumber,omitempty"`
	RegisteredAddress *UpdateIncorporatedBusinessProfileRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredName    *string                                             `json:"registeredName,omitempty"`
	TaxCountryISO     *string                                             `json:"taxCountryISO,omitempty"`
	TradingAddress    *UpdateIncorporatedBusinessProfileTradingAddress    `json:"tradingAddress,omitempty"`
	TradingName       *string                                             `json:"tradingName,omitempty"`
	URL               *string                                             `json:"url,omitempty"`
}

type UpdateIncorporatedBusinessProfileRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateIncorporatedBusinessProfileTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IncorporatedBusiness struct {
	ActiveAt             *string                                   `json:"activeAt,omitempty"`
	Annotations          map[string]string                         `json:"annotations,omitempty"`
	BusinessPersons      []IncorporatedBusinessBusinessPerson      `json:"businessPersons"`
	CddLevel             *CDDLevel                                 `json:"cddLevel,omitempty"`
	CompanyNumber        *string                                   `json:"companyNumber,omitempty"`
	CreatedAt            *string                                   `json:"createdAt,omitempty"`
	ID                   *string                                   `json:"id,omitempty"`
	PreferredInstruments *IncorporatedBusinessPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *IncorporatedBusinessProfileClass         `json:"profile,omitempty"`
	RoleID               *string                                   `json:"roleId,omitempty"`
	ShortID              *string                                   `json:"shortId,omitempty"`
	StateID              *CustomerState                            `json:"stateId,omitempty"`
	Tags                 []string                                  `json:"tags"`
	TargetCDDLevel       *CDDLevel                                 `json:"targetCDDLevel,omitempty"`
	Terms                *IncorporatedBusinessTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                                   `json:"updatedAt,omitempty"`
}

type IncorporatedBusinessBusinessPerson struct {
	CustomerID *string        `json:"customerId,omitempty"`
	ID         *string        `json:"id,omitempty"`
	Profile    *PurpleProfile `json:"profile,omitempty"`
}

type PurpleProfile struct {
	Dob                string                   `json:"dob"`
	FirstName          string                   `json:"firstName"`
	LastName           string                   `json:"lastName"`
	ResidentialAddress StickyResidentialAddress `json:"residentialAddress"`
	Roles              []string                 `json:"roles"`
}

type StickyResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IncorporatedBusinessPreferredInstruments struct {
	BankAccount PurpleBankAccount `json:"bankAccount"`
	Card        PurpleCard        `json:"card"`
}

type PurpleBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type PurpleCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type IncorporatedBusinessProfileClass struct {
	Email             string                  `json:"email"`
	IncorporatedDate  *string                 `json:"incorporatedDate,omitempty"`
	Industry          *string                 `json:"industry,omitempty"`
	OrgChart          *string                 `json:"orgChart,omitempty"`
	PhoneNumber       string                  `json:"phoneNumber"`
	RegisteredAddress PurpleRegisteredAddress `json:"registeredAddress"`
	RegisteredName    string                  `json:"registeredName"`
	TaxCountryISO     string                  `json:"taxCountryISO"`
	TradingAddress    PurpleTradingAddress    `json:"tradingAddress"`
	TradingName       *string                 `json:"tradingName,omitempty"`
	URL               *string                 `json:"url,omitempty"`
}

type PurpleRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type PurpleTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IncorporatedBusinessTerms struct {
	AcceptedAt time.Time `json:"acceptedAt"`
	Revision   string    `json:"revision"`
}

type IncorporatedBusinessRole string

const (
	FullCustomer     IncorporatedBusinessRole = "customer"
	CardOnlyCustomer IncorporatedBusinessRole = "customer-nocdd"
)

type CreateIncorporatedBusiness struct {
	Annotations   map[string]string                 `json:"annotations,omitempty"`
	CompanyNumber string                            `json:"companyNumber"`
	Profile       CreateIncorporatedBusinessProfile `json:"profile"`
	Tags          []string                          `json:"tags"`
	Terms         CreateIncorporatedBusinessTerms   `json:"terms"`
	Role          *CreateIncorporatedBusinessRole   `json:"role,omitempty"`
	RoleID        *string                           `json:"roleId,omitempty"`
	RoleSlug      IncorporatedBusinessRole          `json:"roleSlug,omitempty"`
}

type CreateIncorporatedBusinessProfile struct {
	Email             string                  `json:"email"`
	IncorporatedDate  time.Time               `json:"incorporatedDate,omitempty"`
	Industry          string                  `json:"industry,omitempty"`
	OrgChart          string                  `json:"orgChart,omitempty"`
	PhoneNumber       string                  `json:"phoneNumber"`
	RegisteredAddress FluffyRegisteredAddress `json:"registeredAddress"`
	RegisteredName    string                  `json:"registeredName"`
	TaxCountryISO     string                  `json:"taxCountryISO"`
	TradingAddress    FluffyTradingAddress    `json:"tradingAddress"`
	TradingName       string                  `json:"tradingName,omitempty"`
	URL               string                  `json:"url,omitempty"`
}

type FluffyRegisteredAddress struct {
	CountryISO      string `json:"countryISO"`
	FlatNumber      string `json:"flatNumber,omitempty"`
	HouseNameNumber string `json:"houseNameNumber,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	Region          string `json:"region,omitempty"`
	Street          string `json:"street,omitempty"`
	TownCity        string `json:"townCity,omitempty"`
}

type FluffyTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CreateIncorporatedBusinessTerms struct {
	AcceptedAt time.Time `json:"acceptedAt"`
	Revision   string    `json:"revision"`
}

type UpdateIncorporatedBusiness struct {
	ID       string                                  `json:"id"`
	Profile  *UpdateIncorporatedBusinessProfileClass `json:"profile,omitempty"`
	Terms    *UpdateIncorporatedBusinessTerms        `json:"terms,omitempty"`
	Role     *CreateIncorporatedBusinessRole         `json:"role,omitempty"`
	RoleID   *string                                 `json:"roleId,omitempty"`
	RoleSlug *string                                 `json:"roleSlug,omitempty"`
}

type UpdateIncorporatedBusinessProfileClass struct {
	Email             *string                     `json:"email,omitempty"`
	IncorporatedDate  *string                     `json:"incorporatedDate,omitempty"`
	Industry          *string                     `json:"industry,omitempty"`
	OrgChart          *string                     `json:"orgChart,omitempty"`
	PhoneNumber       *string                     `json:"phoneNumber,omitempty"`
	RegisteredAddress *TentacledRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredName    *string                     `json:"registeredName,omitempty"`
	TaxCountryISO     *string                     `json:"taxCountryISO,omitempty"`
	TradingAddress    *TentacledTradingAddress    `json:"tradingAddress,omitempty"`
	TradingName       *string                     `json:"tradingName,omitempty"`
	URL               *string                     `json:"url,omitempty"`
}

type TentacledRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type TentacledTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateIncorporatedBusinessTerms struct {
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

type IndividualProfile struct {
	Dob                string                              `json:"dob"`
	Email              string                              `json:"email"`
	FirstName          string                              `json:"firstName"`
	LastName           string                              `json:"lastName"`
	PhoneNumber        string                              `json:"phoneNumber"`
	ResidentialAddress IndividualProfileResidentialAddress `json:"residentialAddress"`
}

type IndividualProfileResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateIndividualProfile struct {
	Dob                *string                                    `json:"dob,omitempty"`
	Email              *string                                    `json:"email,omitempty"`
	FirstName          *string                                    `json:"firstName,omitempty"`
	LastName           *string                                    `json:"lastName,omitempty"`
	PhoneNumber        *string                                    `json:"phoneNumber,omitempty"`
	ResidentialAddress *UpdateIndividualProfileResidentialAddress `json:"residentialAddress,omitempty"`
}

type UpdateIndividualProfileResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type Individual struct {
	ActiveAt             *string                         `json:"activeAt,omitempty"`
	Annotations          map[string]string               `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                       `json:"cddLevel,omitempty"`
	CreatedAt            *string                         `json:"createdAt,omitempty"`
	ID                   *string                         `json:"id,omitempty"`
	PreferredInstruments *IndividualPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *IndividualProfileClass         `json:"profile,omitempty"`
	RoleID               *string                         `json:"roleId,omitempty"`
	ShortID              *string                         `json:"shortId,omitempty"`
	StateID              *CustomerState                  `json:"stateId,omitempty"`
	Tags                 []string                        `json:"tags"`
	TargetCDDLevel       *CDDLevel                       `json:"targetCDDLevel,omitempty"`
	Terms                *IndividualTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                         `json:"updatedAt,omitempty"`
}

type IndividualPreferredInstruments struct {
	BankAccount FluffyBankAccount `json:"bankAccount"`
	Card        FluffyCard        `json:"card"`
}

type FluffyBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type FluffyCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type IndividualProfileClass struct {
	Dob                string                   `json:"dob"`
	Email              string                   `json:"email"`
	FirstName          string                   `json:"firstName"`
	LastName           string                   `json:"lastName"`
	PhoneNumber        string                   `json:"phoneNumber"`
	ResidentialAddress IndigoResidentialAddress `json:"residentialAddress"`
}

type IndigoResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IndividualTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CreateIndividual struct {
	Annotations map[string]string               `json:"annotations,omitempty"`
	Profile     CreateIndividualProfile         `json:"profile"`
	Tags        []string                        `json:"tags"`
	Terms       CreateIndividualTerms           `json:"terms"`
	Role        *CreateIncorporatedBusinessRole `json:"role,omitempty"`
	RoleID      *string                         `json:"roleId,omitempty"`
	RoleSlug    *string                         `json:"roleSlug,omitempty"`
}

type CreateIndividualProfile struct {
	Dob                string                     `json:"dob"`
	Email              string                     `json:"email"`
	FirstName          string                     `json:"firstName"`
	LastName           string                     `json:"lastName"`
	PhoneNumber        string                     `json:"phoneNumber"`
	ResidentialAddress IndecentResidentialAddress `json:"residentialAddress"`
}

type IndecentResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CreateIndividualTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type UpdateIndividual struct {
	ID       string                          `json:"id"`
	Profile  *UpdateIndividualProfileClass   `json:"profile,omitempty"`
	Terms    *UpdateIndividualTerms          `json:"terms,omitempty"`
	Role     *CreateIncorporatedBusinessRole `json:"role,omitempty"`
	RoleID   *string                         `json:"roleId,omitempty"`
	RoleSlug *string                         `json:"roleSlug,omitempty"`
}

type UpdateIndividualProfileClass struct {
	Dob                *string                      `json:"dob,omitempty"`
	Email              *string                      `json:"email,omitempty"`
	FirstName          *string                      `json:"firstName,omitempty"`
	LastName           *string                      `json:"lastName,omitempty"`
	PhoneNumber        *string                      `json:"phoneNumber,omitempty"`
	ResidentialAddress *HilariousResidentialAddress `json:"residentialAddress,omitempty"`
}

type HilariousResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateIndividualTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type Customer struct {
	IncorporatedBusiness *CustomerIncorporatedBusiness `json:"incorporatedBusiness,omitempty"`
	Type                 *CustomerType                 `json:"type,omitempty"`
	Individual           *CustomerIndividual           `json:"individual,omitempty"`
	SoleTrader           *CustomerSoleTrader           `json:"soleTrader,omitempty"`
	Organisation         *CustomerOrganisation         `json:"organisation,omitempty"`
}

type CustomerIncorporatedBusiness struct {
	ActiveAt             *string                     `json:"activeAt,omitempty"`
	Annotations          map[string]string           `json:"annotations,omitempty"`
	BusinessPersons      []PurpleBusinessPerson      `json:"businessPersons"`
	CddLevel             *CDDLevel                   `json:"cddLevel,omitempty"`
	CompanyNumber        *string                     `json:"companyNumber,omitempty"`
	CreatedAt            *string                     `json:"createdAt,omitempty"`
	ID                   *string                     `json:"id,omitempty"`
	PreferredInstruments *PurplePreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *TentacledProfile           `json:"profile,omitempty"`
	RoleID               *string                     `json:"roleId,omitempty"`
	ShortID              *string                     `json:"shortId,omitempty"`
	StateID              *CustomerState              `json:"stateId,omitempty"`
	Tags                 []string                    `json:"tags"`
	TargetCDDLevel       *CDDLevel                   `json:"targetCDDLevel,omitempty"`
	Terms                *PurpleTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                     `json:"updatedAt,omitempty"`
}

type PurpleBusinessPerson struct {
	CustomerID *string        `json:"customerId,omitempty"`
	ID         *string        `json:"id,omitempty"`
	Profile    *FluffyProfile `json:"profile,omitempty"`
}

type FluffyProfile struct {
	Dob                string                      `json:"dob"`
	FirstName          string                      `json:"firstName"`
	LastName           string                      `json:"lastName"`
	ResidentialAddress AmbitiousResidentialAddress `json:"residentialAddress"`
	Roles              []string                    `json:"roles"`
}

type AmbitiousResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type PurplePreferredInstruments struct {
	BankAccount TentacledBankAccount `json:"bankAccount"`
	Card        TentacledCard        `json:"card"`
}

type TentacledBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type TentacledCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type TentacledProfile struct {
	Email             string                  `json:"email"`
	IncorporatedDate  *string                 `json:"incorporatedDate,omitempty"`
	Industry          *string                 `json:"industry,omitempty"`
	OrgChart          *string                 `json:"orgChart,omitempty"`
	PhoneNumber       string                  `json:"phoneNumber"`
	RegisteredAddress StickyRegisteredAddress `json:"registeredAddress"`
	RegisteredName    string                  `json:"registeredName"`
	TaxCountryISO     string                  `json:"taxCountryISO"`
	TradingAddress    StickyTradingAddress    `json:"tradingAddress"`
	TradingName       *string                 `json:"tradingName,omitempty"`
	URL               *string                 `json:"url,omitempty"`
}

type StickyRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type StickyTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type PurpleTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CustomerIndividual struct {
	ActiveAt             *string                     `json:"activeAt,omitempty"`
	Annotations          map[string]string           `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                   `json:"cddLevel,omitempty"`
	CreatedAt            *string                     `json:"createdAt,omitempty"`
	ID                   *string                     `json:"id,omitempty"`
	PreferredInstruments *FluffyPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *StickyProfile              `json:"profile,omitempty"`
	RoleID               *string                     `json:"roleId,omitempty"`
	ShortID              *string                     `json:"shortId,omitempty"`
	StateID              *CustomerState              `json:"stateId,omitempty"`
	Tags                 []string                    `json:"tags"`
	TargetCDDLevel       *CDDLevel                   `json:"targetCDDLevel,omitempty"`
	Terms                *FluffyTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                     `json:"updatedAt,omitempty"`
}

type FluffyPreferredInstruments struct {
	BankAccount StickyBankAccount `json:"bankAccount"`
	Card        StickyCard        `json:"card"`
}

type StickyBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type StickyCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type StickyProfile struct {
	Dob                string                    `json:"dob"`
	Email              string                    `json:"email"`
	FirstName          string                    `json:"firstName"`
	LastName           string                    `json:"lastName"`
	PhoneNumber        string                    `json:"phoneNumber"`
	ResidentialAddress CunningResidentialAddress `json:"residentialAddress"`
}

type CunningResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type FluffyTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CustomerOrganisation struct {
	ActiveAt             *string                        `json:"activeAt,omitempty"`
	Annotations          map[string]string              `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                      `json:"cddLevel,omitempty"`
	CreatedAt            *string                        `json:"createdAt,omitempty"`
	ID                   *string                        `json:"id,omitempty"`
	PreferredInstruments *TentacledPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *IndigoProfile                 `json:"profile,omitempty"`
	RoleID               *string                        `json:"roleId,omitempty"`
	ShortID              *string                        `json:"shortId,omitempty"`
	StateID              *CustomerState                 `json:"stateId,omitempty"`
	Tags                 []string                       `json:"tags"`
	TargetCDDLevel       *CDDLevel                      `json:"targetCDDLevel,omitempty"`
	Terms                *TentacledTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                        `json:"updatedAt,omitempty"`
}

type TentacledPreferredInstruments struct {
	BankAccount IndigoBankAccount `json:"bankAccount"`
	Card        IndigoCard        `json:"card"`
}

type IndigoBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type IndigoCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type IndigoProfile struct {
	Email                    *string                  `json:"email,omitempty"`
	Name                     *string                  `json:"name,omitempty"`
	PhoneNumber              *string                  `json:"phoneNumber,omitempty"`
	RegisteredAddress        *IndigoRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                  `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string        `json:"supplementaryInformation,omitempty"`
}

type IndigoRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type TentacledTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CustomerSoleTrader struct {
	ActiveAt             *string                     `json:"activeAt,omitempty"`
	Annotations          map[string]string           `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                   `json:"cddLevel,omitempty"`
	CreatedAt            *string                     `json:"createdAt,omitempty"`
	ID                   *string                     `json:"id,omitempty"`
	PreferredInstruments *StickyPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *IndecentProfile            `json:"profile,omitempty"`
	RoleID               *string                     `json:"roleId,omitempty"`
	ShortID              *string                     `json:"shortId,omitempty"`
	StateID              *CustomerState              `json:"stateId,omitempty"`
	Tags                 []string                    `json:"tags"`
	TargetCDDLevel       *CDDLevel                   `json:"targetCDDLevel,omitempty"`
	Terms                *StickyTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                     `json:"updatedAt,omitempty"`
}

type StickyPreferredInstruments struct {
	BankAccount IndecentBankAccount `json:"bankAccount"`
	Card        IndecentCard        `json:"card"`
}

type IndecentBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type IndecentCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type IndecentProfile struct {
	Dob                string                    `json:"dob"`
	Email              string                    `json:"email"`
	FirstName          string                    `json:"firstName"`
	Industry           *string                   `json:"industry,omitempty"`
	LastName           string                    `json:"lastName"`
	PhoneNumber        string                    `json:"phoneNumber"`
	ResidentialAddress MagentaResidentialAddress `json:"residentialAddress"`
	TradingName        *string                   `json:"tradingName,omitempty"`
	URL                *string                   `json:"url,omitempty"`
}

type MagentaResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type StickyTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CustomerList struct {
	Results []Result `json:"results"`
}

type Result struct {
	IncorporatedBusiness *ResultIncorporatedBusiness `json:"incorporatedBusiness,omitempty"`
	Type                 *CustomerType               `json:"type,omitempty"`
	Individual           *ResultIndividual           `json:"individual,omitempty"`
	SoleTrader           *ResultSoleTrader           `json:"soleTrader,omitempty"`
	Organisation         *ResultOrganisation         `json:"organisation,omitempty"`
}

type ResultIncorporatedBusiness struct {
	ActiveAt             *string                     `json:"activeAt,omitempty"`
	Annotations          map[string]string           `json:"annotations,omitempty"`
	BusinessPersons      []FluffyBusinessPerson      `json:"businessPersons"`
	CddLevel             *CDDLevel                   `json:"cddLevel,omitempty"`
	CompanyNumber        *string                     `json:"companyNumber,omitempty"`
	CreatedAt            *string                     `json:"createdAt,omitempty"`
	ID                   *string                     `json:"id,omitempty"`
	PreferredInstruments *IndigoPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *AmbitiousProfile           `json:"profile,omitempty"`
	RoleID               *string                     `json:"roleId,omitempty"`
	ShortID              *string                     `json:"shortId,omitempty"`
	StateID              *CustomerState              `json:"stateId,omitempty"`
	Tags                 []string                    `json:"tags"`
	TargetCDDLevel       *CDDLevel                   `json:"targetCDDLevel,omitempty"`
	Terms                *IndigoTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                     `json:"updatedAt,omitempty"`
}

type FluffyBusinessPerson struct {
	CustomerID *string           `json:"customerId,omitempty"`
	ID         *string           `json:"id,omitempty"`
	Profile    *HilariousProfile `json:"profile,omitempty"`
}

type HilariousProfile struct {
	Dob                string                   `json:"dob"`
	FirstName          string                   `json:"firstName"`
	LastName           string                   `json:"lastName"`
	ResidentialAddress FriskyResidentialAddress `json:"residentialAddress"`
	Roles              []string                 `json:"roles"`
}

type FriskyResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IndigoPreferredInstruments struct {
	BankAccount HilariousBankAccount `json:"bankAccount"`
	Card        HilariousCard        `json:"card"`
}

type HilariousBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type HilariousCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type AmbitiousProfile struct {
	Email             string                    `json:"email"`
	IncorporatedDate  *string                   `json:"incorporatedDate,omitempty"`
	Industry          *string                   `json:"industry,omitempty"`
	OrgChart          *string                   `json:"orgChart,omitempty"`
	PhoneNumber       string                    `json:"phoneNumber"`
	RegisteredAddress IndecentRegisteredAddress `json:"registeredAddress"`
	RegisteredName    string                    `json:"registeredName"`
	TaxCountryISO     string                    `json:"taxCountryISO"`
	TradingAddress    IndigoTradingAddress      `json:"tradingAddress"`
	TradingName       *string                   `json:"tradingName,omitempty"`
	URL               *string                   `json:"url,omitempty"`
}

type IndecentRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IndigoTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IndigoTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type ResultIndividual struct {
	ActiveAt             *string                       `json:"activeAt,omitempty"`
	Annotations          map[string]string             `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                     `json:"cddLevel,omitempty"`
	CreatedAt            *string                       `json:"createdAt,omitempty"`
	ID                   *string                       `json:"id,omitempty"`
	PreferredInstruments *IndecentPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *CunningProfile               `json:"profile,omitempty"`
	RoleID               *string                       `json:"roleId,omitempty"`
	ShortID              *string                       `json:"shortId,omitempty"`
	StateID              *CustomerState                `json:"stateId,omitempty"`
	Tags                 []string                      `json:"tags"`
	TargetCDDLevel       *CDDLevel                     `json:"targetCDDLevel,omitempty"`
	Terms                *IndecentTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                       `json:"updatedAt,omitempty"`
}

type IndecentPreferredInstruments struct {
	BankAccount AmbitiousBankAccount `json:"bankAccount"`
	Card        AmbitiousCard        `json:"card"`
}

type AmbitiousBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type AmbitiousCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type CunningProfile struct {
	Dob                string                        `json:"dob"`
	Email              string                        `json:"email"`
	FirstName          string                        `json:"firstName"`
	LastName           string                        `json:"lastName"`
	PhoneNumber        string                        `json:"phoneNumber"`
	ResidentialAddress MischievousResidentialAddress `json:"residentialAddress"`
}

type MischievousResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IndecentTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type ResultOrganisation struct {
	ActiveAt             *string                        `json:"activeAt,omitempty"`
	Annotations          map[string]string              `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                      `json:"cddLevel,omitempty"`
	CreatedAt            *string                        `json:"createdAt,omitempty"`
	ID                   *string                        `json:"id,omitempty"`
	PreferredInstruments *HilariousPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *MagentaProfile                `json:"profile,omitempty"`
	RoleID               *string                        `json:"roleId,omitempty"`
	ShortID              *string                        `json:"shortId,omitempty"`
	StateID              *CustomerState                 `json:"stateId,omitempty"`
	Tags                 []string                       `json:"tags"`
	TargetCDDLevel       *CDDLevel                      `json:"targetCDDLevel,omitempty"`
	Terms                *HilariousTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                        `json:"updatedAt,omitempty"`
}

type HilariousPreferredInstruments struct {
	BankAccount CunningBankAccount `json:"bankAccount"`
	Card        CunningCard        `json:"card"`
}

type CunningBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type CunningCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type MagentaProfile struct {
	Email                    *string                     `json:"email,omitempty"`
	Name                     *string                     `json:"name,omitempty"`
	PhoneNumber              *string                     `json:"phoneNumber,omitempty"`
	RegisteredAddress        *HilariousRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                     `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string           `json:"supplementaryInformation,omitempty"`
}

type HilariousRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type HilariousTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type ResultSoleTrader struct {
	ActiveAt             *string                        `json:"activeAt,omitempty"`
	Annotations          map[string]string              `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                      `json:"cddLevel,omitempty"`
	CreatedAt            *string                        `json:"createdAt,omitempty"`
	ID                   *string                        `json:"id,omitempty"`
	PreferredInstruments *AmbitiousPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *FriskyProfile                 `json:"profile,omitempty"`
	RoleID               *string                        `json:"roleId,omitempty"`
	ShortID              *string                        `json:"shortId,omitempty"`
	StateID              *CustomerState                 `json:"stateId,omitempty"`
	Tags                 []string                       `json:"tags"`
	TargetCDDLevel       *CDDLevel                      `json:"targetCDDLevel,omitempty"`
	Terms                *AmbitiousTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                        `json:"updatedAt,omitempty"`
}

type AmbitiousPreferredInstruments struct {
	BankAccount MagentaBankAccount `json:"bankAccount"`
	Card        MagentaCard        `json:"card"`
}

type MagentaBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type MagentaCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type FriskyProfile struct {
	Dob                string                          `json:"dob"`
	Email              string                          `json:"email"`
	FirstName          string                          `json:"firstName"`
	Industry           *string                         `json:"industry,omitempty"`
	LastName           string                          `json:"lastName"`
	PhoneNumber        string                          `json:"phoneNumber"`
	ResidentialAddress BraggadociousResidentialAddress `json:"residentialAddress"`
	TradingName        *string                         `json:"tradingName,omitempty"`
	URL                *string                         `json:"url,omitempty"`
}

type BraggadociousResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type AmbitiousTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CustomerToken struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

type CreateToken struct {
	ID string `json:"id"`
}

type Transition struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	FromStateID *CustomerState    `json:"fromStateId,omitempty"`
	ID          string            `json:"id"`
	Tags        []string          `json:"tags"`
	ToStateID   CustomerState     `json:"toStateId"`
}

type BusinessPersonID struct {
	CustomerID string `json:"customerId"`
	ID         string `json:"id"`
}

type PurpleIncorporatedBusiness struct {
	IncorporatedBusiness IncorporatedBusinessIncorporatedBusiness `json:"incorporatedBusiness"`
	Type                 *IncorporatedBusinessType                `json:"type,omitempty"`
}

type IncorporatedBusinessIncorporatedBusiness struct {
	ActiveAt             *string                      `json:"activeAt,omitempty"`
	Annotations          map[string]string            `json:"annotations,omitempty"`
	BusinessPersons      []TentacledBusinessPerson    `json:"businessPersons"`
	CddLevel             *CDDLevel                    `json:"cddLevel,omitempty"`
	CompanyNumber        *string                      `json:"companyNumber,omitempty"`
	CreatedAt            *string                      `json:"createdAt,omitempty"`
	ID                   *string                      `json:"id,omitempty"`
	PreferredInstruments *CunningPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *BraggadociousProfile        `json:"profile,omitempty"`
	RoleID               *string                      `json:"roleId,omitempty"`
	ShortID              *string                      `json:"shortId,omitempty"`
	StateID              *CustomerState               `json:"stateId,omitempty"`
	Tags                 []string                     `json:"tags"`
	TargetCDDLevel       *CDDLevel                    `json:"targetCDDLevel,omitempty"`
	Terms                *CunningTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                      `json:"updatedAt,omitempty"`
}

type TentacledBusinessPerson struct {
	CustomerID *string             `json:"customerId,omitempty"`
	ID         *string             `json:"id,omitempty"`
	Profile    *MischievousProfile `json:"profile,omitempty"`
}

type MischievousProfile struct {
	Dob                string              `json:"dob"`
	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	ResidentialAddress ResidentialAddress1 `json:"residentialAddress"`
	Roles              []string            `json:"roles"`
}

type ResidentialAddress1 struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CunningPreferredInstruments struct {
	BankAccount FriskyBankAccount `json:"bankAccount"`
	Card        FriskyCard        `json:"card"`
}

type FriskyBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type FriskyCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type BraggadociousProfile struct {
	Email             string                     `json:"email"`
	IncorporatedDate  *string                    `json:"incorporatedDate,omitempty"`
	Industry          *string                    `json:"industry,omitempty"`
	OrgChart          *string                    `json:"orgChart,omitempty"`
	PhoneNumber       string                     `json:"phoneNumber"`
	RegisteredAddress AmbitiousRegisteredAddress `json:"registeredAddress"`
	RegisteredName    string                     `json:"registeredName"`
	TaxCountryISO     string                     `json:"taxCountryISO"`
	TradingAddress    IndecentTradingAddress     `json:"tradingAddress"`
	TradingName       *string                    `json:"tradingName,omitempty"`
	URL               *string                    `json:"url,omitempty"`
}

type AmbitiousRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type IndecentTradingAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CunningTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type PurpleIndividual struct {
	Individual IndividualIndividual `json:"individual"`
	Type       *IndividualType      `json:"type,omitempty"`
}

type IndividualIndividual struct {
	ActiveAt             *string                      `json:"activeAt,omitempty"`
	Annotations          map[string]string            `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                    `json:"cddLevel,omitempty"`
	CreatedAt            *string                      `json:"createdAt,omitempty"`
	ID                   *string                      `json:"id,omitempty"`
	PreferredInstruments *MagentaPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *Profile1                    `json:"profile,omitempty"`
	RoleID               *string                      `json:"roleId,omitempty"`
	ShortID              *string                      `json:"shortId,omitempty"`
	StateID              *CustomerState               `json:"stateId,omitempty"`
	Tags                 []string                     `json:"tags"`
	TargetCDDLevel       *CDDLevel                    `json:"targetCDDLevel,omitempty"`
	Terms                *MagentaTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                      `json:"updatedAt,omitempty"`
}

type MagentaPreferredInstruments struct {
	BankAccount MischievousBankAccount `json:"bankAccount"`
	Card        MischievousCard        `json:"card"`
}

type MischievousBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type MischievousCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Profile1 struct {
	Dob                string              `json:"dob"`
	Email              string              `json:"email"`
	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	PhoneNumber        string              `json:"phoneNumber"`
	ResidentialAddress ResidentialAddress2 `json:"residentialAddress"`
}

type ResidentialAddress2 struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type MagentaTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type SoleTrader struct {
	SoleTrader SoleTraderSoleTrader `json:"soleTrader"`
	Type       *SoleTraderType      `json:"type,omitempty"`
}

type SoleTraderSoleTrader struct {
	ActiveAt             *string                     `json:"activeAt,omitempty"`
	Annotations          map[string]string           `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                   `json:"cddLevel,omitempty"`
	CreatedAt            *string                     `json:"createdAt,omitempty"`
	ID                   *string                     `json:"id,omitempty"`
	PreferredInstruments *FriskyPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *Profile2                   `json:"profile,omitempty"`
	RoleID               *string                     `json:"roleId,omitempty"`
	ShortID              *string                     `json:"shortId,omitempty"`
	StateID              *CustomerState              `json:"stateId,omitempty"`
	Tags                 []string                    `json:"tags"`
	TargetCDDLevel       *CDDLevel                   `json:"targetCDDLevel,omitempty"`
	Terms                *FriskyTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                     `json:"updatedAt,omitempty"`
}

type FriskyPreferredInstruments struct {
	BankAccount BraggadociousBankAccount `json:"bankAccount"`
	Card        BraggadociousCard        `json:"card"`
}

type BraggadociousBankAccount struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type BraggadociousCard struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Profile2 struct {
	Dob                string              `json:"dob"`
	Email              string              `json:"email"`
	FirstName          string              `json:"firstName"`
	Industry           *string             `json:"industry,omitempty"`
	LastName           string              `json:"lastName"`
	PhoneNumber        string              `json:"phoneNumber"`
	ResidentialAddress ResidentialAddress3 `json:"residentialAddress"`
	TradingName        *string             `json:"tradingName,omitempty"`
	URL                *string             `json:"url,omitempty"`
}

type ResidentialAddress3 struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type FriskyTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type Organisation struct {
	Organisation OrganisationOrganisation `json:"organisation"`
	Type         *OrganisationType        `json:"type,omitempty"`
}

type OrganisationOrganisation struct {
	ActiveAt             *string                          `json:"activeAt,omitempty"`
	Annotations          map[string]string                `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                        `json:"cddLevel,omitempty"`
	CreatedAt            *string                          `json:"createdAt,omitempty"`
	ID                   *string                          `json:"id,omitempty"`
	PreferredInstruments *MischievousPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *Profile3                        `json:"profile,omitempty"`
	RoleID               *string                          `json:"roleId,omitempty"`
	ShortID              *string                          `json:"shortId,omitempty"`
	StateID              *CustomerState                   `json:"stateId,omitempty"`
	Tags                 []string                         `json:"tags"`
	TargetCDDLevel       *CDDLevel                        `json:"targetCDDLevel,omitempty"`
	Terms                *MischievousTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                          `json:"updatedAt,omitempty"`
}

type MischievousPreferredInstruments struct {
	BankAccount BankAccount1 `json:"bankAccount"`
	Card        Card1        `json:"card"`
}

type BankAccount1 struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Card1 struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Profile3 struct {
	Email                    *string                   `json:"email,omitempty"`
	Name                     *string                   `json:"name,omitempty"`
	PhoneNumber              *string                   `json:"phoneNumber,omitempty"`
	RegisteredAddress        *CunningRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                   `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string         `json:"supplementaryInformation,omitempty"`
}

type CunningRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type MischievousTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type OrganisationProfile struct {
	Email                    *string                               `json:"email,omitempty"`
	Name                     *string                               `json:"name,omitempty"`
	PhoneNumber              *string                               `json:"phoneNumber,omitempty"`
	RegisteredAddress        *OrganisationProfileRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                               `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string                     `json:"supplementaryInformation,omitempty"`
}

type OrganisationProfileRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateOrganisationProfile struct {
	Email                    *string                                     `json:"email,omitempty"`
	Name                     *string                                     `json:"name,omitempty"`
	PhoneNumber              *string                                     `json:"phoneNumber,omitempty"`
	RegisteredAddress        *UpdateOrganisationProfileRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                                     `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string                           `json:"supplementaryInformation,omitempty"`
}

type UpdateOrganisationProfileRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type PurpleOrganisation struct {
	ActiveAt             *string                           `json:"activeAt,omitempty"`
	Annotations          map[string]string                 `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                         `json:"cddLevel,omitempty"`
	CreatedAt            *string                           `json:"createdAt,omitempty"`
	ID                   *string                           `json:"id,omitempty"`
	PreferredInstruments *OrganisationPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *OrganisationProfileClass         `json:"profile,omitempty"`
	RoleID               *string                           `json:"roleId,omitempty"`
	ShortID              *string                           `json:"shortId,omitempty"`
	StateID              *CustomerState                    `json:"stateId,omitempty"`
	Tags                 []string                          `json:"tags"`
	TargetCDDLevel       *CDDLevel                         `json:"targetCDDLevel,omitempty"`
	Terms                *OrganisationTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                           `json:"updatedAt,omitempty"`
}

type OrganisationPreferredInstruments struct {
	BankAccount BankAccount2 `json:"bankAccount"`
	Card        Card2        `json:"card"`
}

type BankAccount2 struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Card2 struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type OrganisationProfileClass struct {
	Email                    *string                   `json:"email,omitempty"`
	Name                     *string                   `json:"name,omitempty"`
	PhoneNumber              *string                   `json:"phoneNumber,omitempty"`
	RegisteredAddress        *MagentaRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                   `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string         `json:"supplementaryInformation,omitempty"`
}

type MagentaRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type OrganisationTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CreateOrganisation struct {
	Annotations map[string]string               `json:"annotations,omitempty"`
	Profile     CreateOrganisationProfile       `json:"profile"`
	Tags        []string                        `json:"tags"`
	Terms       CreateOrganisationTerms         `json:"terms"`
	Role        *CreateIncorporatedBusinessRole `json:"role,omitempty"`
	RoleID      *string                         `json:"roleId,omitempty"`
	RoleSlug    *string                         `json:"roleSlug,omitempty"`
}

type CreateOrganisationProfile struct {
	Email                    *string                  `json:"email,omitempty"`
	Name                     *string                  `json:"name,omitempty"`
	PhoneNumber              *string                  `json:"phoneNumber,omitempty"`
	RegisteredAddress        *FriskyRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                  `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string        `json:"supplementaryInformation,omitempty"`
}

type FriskyRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CreateOrganisationTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type UpdateOrganisation struct {
	ID       string                          `json:"id"`
	Profile  *UpdateOrganisationProfileClass `json:"profile,omitempty"`
	Terms    *UpdateOrganisationTerms        `json:"terms,omitempty"`
	Role     *CreateIncorporatedBusinessRole `json:"role,omitempty"`
	RoleID   *string                         `json:"roleId,omitempty"`
	RoleSlug *string                         `json:"roleSlug,omitempty"`
}

type UpdateOrganisationProfileClass struct {
	Email                    *string                       `json:"email,omitempty"`
	Name                     *string                       `json:"name,omitempty"`
	PhoneNumber              *string                       `json:"phoneNumber,omitempty"`
	RegisteredAddress        *MischievousRegisteredAddress `json:"registeredAddress,omitempty"`
	RegisteredNumber         *string                       `json:"registeredNumber,omitempty"`
	SupplementaryInformation map[string]string             `json:"supplementaryInformation,omitempty"`
}

type MischievousRegisteredAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateOrganisationTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type SoleTraderProfile struct {
	Dob                string                              `json:"dob"`
	Email              string                              `json:"email"`
	FirstName          string                              `json:"firstName"`
	Industry           *string                             `json:"industry,omitempty"`
	LastName           string                              `json:"lastName"`
	PhoneNumber        string                              `json:"phoneNumber"`
	ResidentialAddress SoleTraderProfileResidentialAddress `json:"residentialAddress"`
	TradingName        *string                             `json:"tradingName,omitempty"`
	URL                *string                             `json:"url,omitempty"`
}

type SoleTraderProfileResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateSoleTraderProfile struct {
	Dob                *string                                    `json:"dob,omitempty"`
	Email              *string                                    `json:"email,omitempty"`
	FirstName          *string                                    `json:"firstName,omitempty"`
	Industry           *string                                    `json:"industry,omitempty"`
	LastName           *string                                    `json:"lastName,omitempty"`
	PhoneNumber        *string                                    `json:"phoneNumber,omitempty"`
	ResidentialAddress *UpdateSoleTraderProfileResidentialAddress `json:"residentialAddress,omitempty"`
	TradingName        *string                                    `json:"tradingName,omitempty"`
	URL                *string                                    `json:"url,omitempty"`
}

type UpdateSoleTraderProfileResidentialAddress struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type PurpleSoleTrader struct {
	ActiveAt             *string                         `json:"activeAt,omitempty"`
	Annotations          map[string]string               `json:"annotations,omitempty"`
	CddLevel             *CDDLevel                       `json:"cddLevel,omitempty"`
	CreatedAt            *string                         `json:"createdAt,omitempty"`
	ID                   *string                         `json:"id,omitempty"`
	PreferredInstruments *SoleTraderPreferredInstruments `json:"preferredInstruments,omitempty"`
	Profile              *SoleTraderProfileClass         `json:"profile,omitempty"`
	RoleID               *string                         `json:"roleId,omitempty"`
	ShortID              *string                         `json:"shortId,omitempty"`
	StateID              *CustomerState                  `json:"stateId,omitempty"`
	Tags                 []string                        `json:"tags"`
	TargetCDDLevel       *CDDLevel                       `json:"targetCDDLevel,omitempty"`
	Terms                *SoleTraderTerms                `json:"terms,omitempty"`
	UpdatedAt            *string                         `json:"updatedAt,omitempty"`
}

type SoleTraderPreferredInstruments struct {
	BankAccount BankAccount3 `json:"bankAccount"`
	Card        Card3        `json:"card"`
}

type BankAccount3 struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Card3 struct {
	ID        *string `json:"id,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type SoleTraderProfileClass struct {
	Dob                string              `json:"dob"`
	Email              string              `json:"email"`
	FirstName          string              `json:"firstName"`
	Industry           *string             `json:"industry,omitempty"`
	LastName           string              `json:"lastName"`
	PhoneNumber        string              `json:"phoneNumber"`
	ResidentialAddress ResidentialAddress4 `json:"residentialAddress"`
	TradingName        *string             `json:"tradingName,omitempty"`
	URL                *string             `json:"url,omitempty"`
}

type ResidentialAddress4 struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type SoleTraderTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type CreateSoleTrader struct {
	Annotations map[string]string               `json:"annotations,omitempty"`
	Profile     CreateSoleTraderProfile         `json:"profile"`
	Tags        []string                        `json:"tags"`
	Terms       CreateSoleTraderTerms           `json:"terms"`
	Role        *CreateIncorporatedBusinessRole `json:"role,omitempty"`
	RoleID      *string                         `json:"roleId,omitempty"`
	RoleSlug    *string                         `json:"roleSlug,omitempty"`
}

type CreateSoleTraderProfile struct {
	Dob                string              `json:"dob"`
	Email              string              `json:"email"`
	FirstName          string              `json:"firstName"`
	Industry           *string             `json:"industry,omitempty"`
	LastName           string              `json:"lastName"`
	PhoneNumber        string              `json:"phoneNumber"`
	ResidentialAddress ResidentialAddress5 `json:"residentialAddress"`
	TradingName        *string             `json:"tradingName,omitempty"`
	URL                *string             `json:"url,omitempty"`
}

type ResidentialAddress5 struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type CreateSoleTraderTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
}

type UpdateSoleTrader struct {
	ID       string                          `json:"id"`
	Profile  *UpdateSoleTraderProfileClass   `json:"profile,omitempty"`
	Terms    *UpdateSoleTraderTerms          `json:"terms,omitempty"`
	Role     *CreateIncorporatedBusinessRole `json:"role,omitempty"`
	RoleID   *string                         `json:"roleId,omitempty"`
	RoleSlug *string                         `json:"roleSlug,omitempty"`
}

type UpdateSoleTraderProfileClass struct {
	Dob                *string              `json:"dob,omitempty"`
	Email              *string              `json:"email,omitempty"`
	FirstName          *string              `json:"firstName,omitempty"`
	Industry           *string              `json:"industry,omitempty"`
	LastName           *string              `json:"lastName,omitempty"`
	PhoneNumber        *string              `json:"phoneNumber,omitempty"`
	ResidentialAddress *ResidentialAddress6 `json:"residentialAddress,omitempty"`
	TradingName        *string              `json:"tradingName,omitempty"`
	URL                *string              `json:"url,omitempty"`
}

type ResidentialAddress6 struct {
	CountryISO      string  `json:"countryISO"`
	FlatNumber      *string `json:"flatNumber,omitempty"`
	HouseNameNumber *string `json:"houseNameNumber,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Region          *string `json:"region,omitempty"`
	Street          *string `json:"street,omitempty"`
	TownCity        *string `json:"townCity,omitempty"`
}

type UpdateSoleTraderTerms struct {
	AcceptedAt string `json:"acceptedAt"`
	Revision   string `json:"revision"`
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

type CDDLevel string

const (
	Five  CDDLevel = "FIVE"
	Four  CDDLevel = "FOUR"
	One   CDDLevel = "ONE"
	Three CDDLevel = "THREE"
	Two   CDDLevel = "TWO"
	Zero  CDDLevel = "ZERO"
)

type CustomerState string

const (
	Closing    CustomerState = "CLOSING"
	Disabled   CustomerState = "DISABLED"
	Enabled    CustomerState = "ENABLED"
	Locked     CustomerState = "LOCKED"
	Pending    CustomerState = "PENDING"
	Terminated CustomerState = "TERMINATED"
)

type CreateIncorporatedBusinessRole string

const (
	PurpleRoleID   CreateIncorporatedBusinessRole = "roleId"
	PurpleRoleSlug CreateIncorporatedBusinessRole = "roleSlug"
)

type RoleIDRole string

const (
	FluffyRoleID RoleIDRole = "roleId"
)

type RoleSlugRole string

const (
	FluffyRoleSlug RoleSlugRole = "roleSlug"
)

type CustomerType string

const (
	FluffyIncorporatedBusiness CustomerType = "incorporatedBusiness"
	FluffyIndividual           CustomerType = "individual"
	FluffyOrganisation         CustomerType = "organisation"
	FluffySoleTrader           CustomerType = "soleTrader"
)

type IncorporatedBusinessType string

const (
	TentacledIncorporatedBusiness IncorporatedBusinessType = "incorporatedBusiness"
)

type IndividualType string

const (
	TentacledIndividual IndividualType = "individual"
)

type SoleTraderType string

const (
	TentacledSoleTrader SoleTraderType = "soleTrader"
)

type OrganisationType string

const (
	TentacledOrganisation OrganisationType = "organisation"
)
