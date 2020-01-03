package auth

import (
	"time"
)

type Factor struct {
	ExpiredAt *string     `json:"expiredAt,omitempty"`
	Valid     *bool       `json:"valid,omitempty"`
	Knows     *Knows      `json:"knows,omitempty"`
	Type      *FactorType `json:"type,omitempty"`
	Has       *Has        `json:"has,omitempty"`
	Is        *Is         `json:"is,omitempty"`
}

type AuthenticationRequirements struct {
	Missing         *float64 `json:"missing,omitempty"`
	PossibleFactors []string `json:"possibleFactors"`
	ProvidedFactors []string `json:"providedFactors"`
	Undefined       []string `json:"undefined"`
}

type AuthenticationToken struct {
	AccessToken  *string       `json:"accessToken,omitempty"`
	ExpiresAt    time.Time     `json:"expiresAt"`
	Requirements *Requirements `json:"requirements,omitempty"`
	ShortID      *string       `json:"shortId,omitempty"`
	Success      bool          `json:"success"`
}

type Requirements struct {
	Missing         *float64 `json:"missing,omitempty"`
	PossibleFactors []string `json:"possibleFactors"`
	ProvidedFactors []string `json:"providedFactors"`
	Undefined       []string `json:"undefined"`
}

type Attestation struct {
	CapturedAt time.Time   `json:"capturedAt"`
	Knows      *Knows      `json:"knows,omitempty"`
	Type       *FactorType `json:"type,omitempty"`
	Has        *Has        `json:"has,omitempty"`
	Is         *Is         `json:"is,omitempty"`
}

type CreateToken struct {
	Attestations []AttestationElement `json:"attestations"`
	ID           string               `json:"id"`
}

type AttestationElement struct {
	CapturedAt time.Time   `json:"capturedAt"`
	Knows      *Knows      `json:"knows,omitempty"`
	Type       *FactorType `json:"type,omitempty"`
	Has        *Has        `json:"has,omitempty"`
	Is         *Is         `json:"is,omitempty"`
}

type PurpleKnows struct {
	Knows Knows      `json:"knows"`
	Type  *KnowsType `json:"type,omitempty"`
}

type PurpleHas struct {
	Has  Has      `json:"has"`
	Type *HasType `json:"type,omitempty"`
}

type PurpleIs struct {
	Is   Is      `json:"is"`
	Type *IsType `json:"type,omitempty"`
}

type Has string

const (
	Authenticator Has = "AUTHENTICATOR"
	Device        Has = "DEVICE"
	Email         Has = "EMAIL"
	PhoneNumber   Has = "PHONE_NUMBER"
)

type Is string

const (
	FacialRecognition Is = "FACIAL_RECOGNITION"
	FingerPrint       Is = "FINGER_PRINT"
)

type Knows string

const (
	Password Knows = "PASSWORD"
	Pin      Knows = "PIN"
)

type FactorType string

const (
	FluffyHas   FactorType = "has"
	FluffyIs    FactorType = "is"
	FluffyKnows FactorType = "knows"
)

type KnowsType string

const (
	TentacledKnows KnowsType = "knows"
)

type HasType string

const (
	TentacledHas HasType = "has"
)

type IsType string

const (
	TentacledIs IsType = "is"
)
