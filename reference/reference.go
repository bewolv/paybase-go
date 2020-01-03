package reference

type Reference struct {
	AccountName   *string `json:"accountName,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	Iban          *string `json:"IBAN,omitempty"`
	Reference     *string `json:"reference,omitempty"`
	SortCode      *string `json:"sortCode,omitempty"`
}

type Get struct {
	FromID *string `json:"fromId,omitempty"`
	ToID   string  `json:"toId"`
}
