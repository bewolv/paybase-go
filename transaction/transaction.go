package transaction

type Adjustment struct {
	AccountID          *string         `json:"accountId,omitempty"`
	Amount             *string         `json:"amount,omitempty"`
	CreatedAt          *string         `json:"createdAt,omitempty"`
	Description        *string         `json:"description,omitempty"`
	ID                 *string         `json:"id,omitempty"`
	LinkedAdjustmentID *string         `json:"linkedAdjustmentId,omitempty"`
	Rate               *string         `json:"rate,omitempty"`
	Type               *AdjustmentType `json:"type,omitempty"`
}

type Transaction struct {
	Adjustments         []TransactionAdjustment `json:"adjustments"`
	Amount              *string                 `json:"amount,omitempty"`
	Annotations         map[string]string       `json:"annotations,omitempty"`
	CreatedAt           *string                 `json:"createdAt,omitempty"`
	Description         *string                 `json:"description,omitempty"`
	Flow                *Flow                   `json:"flow,omitempty"`
	FromOwnerID         *string                 `json:"fromOwnerId,omitempty"`
	GroupID             *string                 `json:"groupId,omitempty"`
	ID                  *string                 `json:"id,omitempty"`
	Method              *string                 `json:"method,omitempty"`
	PaymentInstrumentID *string                 `json:"paymentInstrumentId,omitempty"`
	Purpose             *Purpose                `json:"purpose,omitempty"`
	Reference           *string                 `json:"reference,omitempty"`
	Requirements        map[string]string       `json:"requirements,omitempty"`
	StateID             *TransactionState       `json:"stateId,omitempty"`
	Tags                []string                `json:"tags"`
	ToOwnerID           *string                 `json:"toOwnerId,omitempty"`
	UpdatedAt           *string                 `json:"updatedAt,omitempty"`
}

type TransactionAdjustment struct {
	AccountID          *string         `json:"accountId,omitempty"`
	Amount             *string         `json:"amount,omitempty"`
	CreatedAt          *string         `json:"createdAt,omitempty"`
	Description        *string         `json:"description,omitempty"`
	ID                 *string         `json:"id,omitempty"`
	LinkedAdjustmentID *string         `json:"linkedAdjustmentId,omitempty"`
	Rate               *string         `json:"rate,omitempty"`
	Type               *AdjustmentType `json:"type,omitempty"`
}

type TransactionList struct {
	Results []Result `json:"results"`
}

type Result struct {
	Adjustments         []ResultAdjustment `json:"adjustments"`
	Amount              *string            `json:"amount,omitempty"`
	Annotations         map[string]string  `json:"annotations,omitempty"`
	CreatedAt           *string            `json:"createdAt,omitempty"`
	Description         *string            `json:"description,omitempty"`
	Flow                *Flow              `json:"flow,omitempty"`
	FromOwnerID         *string            `json:"fromOwnerId,omitempty"`
	GroupID             *string            `json:"groupId,omitempty"`
	ID                  *string            `json:"id,omitempty"`
	Method              *string            `json:"method,omitempty"`
	PaymentInstrumentID *string            `json:"paymentInstrumentId,omitempty"`
	Purpose             *Purpose           `json:"purpose,omitempty"`
	Reference           *string            `json:"reference,omitempty"`
	Requirements        map[string]string  `json:"requirements,omitempty"`
	StateID             *TransactionState  `json:"stateId,omitempty"`
	Tags                []string           `json:"tags"`
	ToOwnerID           *string            `json:"toOwnerId,omitempty"`
	UpdatedAt           *string            `json:"updatedAt,omitempty"`
}

type ResultAdjustment struct {
	AccountID          *string         `json:"accountId,omitempty"`
	Amount             *string         `json:"amount,omitempty"`
	CreatedAt          *string         `json:"createdAt,omitempty"`
	Description        *string         `json:"description,omitempty"`
	ID                 *string         `json:"id,omitempty"`
	LinkedAdjustmentID *string         `json:"linkedAdjustmentId,omitempty"`
	Rate               *string         `json:"rate,omitempty"`
	Type               *AdjustmentType `json:"type,omitempty"`
}

type AdjustedOptions struct {
	AdjustmentAmount *string `json:"adjustmentAmount,omitempty"`
	AdjustmentID     *string `json:"adjustmentId,omitempty"`
	Reason           *string `json:"reason,omitempty"`
}

type TransitionOptions struct {
	Adjustment *TransitionOptionsAdjustment `json:"adjustment,omitempty"`
	Type       *Type                        `json:"type,omitempty"`
}

type TransitionOptionsAdjustment struct {
	AdjustmentAmount *string `json:"adjustmentAmount,omitempty"`
	AdjustmentID     *string `json:"adjustmentId,omitempty"`
	Reason           *string `json:"reason,omitempty"`
}

type Transition struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	FromStateID *TransactionState `json:"fromStateId,omitempty"`
	ID          string            `json:"id"`
	Options     *Options          `json:"options,omitempty"`
	Tags        []string          `json:"tags"`
	ToStateID   TransactionState  `json:"toStateId"`
}

type Options struct {
	Adjustment *OptionsAdjustment `json:"adjustment,omitempty"`
	Type       *Type              `json:"type,omitempty"`
}

type OptionsAdjustment struct {
	AdjustmentAmount *string `json:"adjustmentAmount,omitempty"`
	AdjustmentID     *string `json:"adjustmentId,omitempty"`
	Reason           *string `json:"reason,omitempty"`
}

type PurpleAdjustment struct {
	Adjustment AdjustmentAdjustment `json:"adjustment"`
	Type       *Type                `json:"type,omitempty"`
}

type AdjustmentAdjustment struct {
	AdjustmentAmount *string `json:"adjustmentAmount,omitempty"`
	AdjustmentID     *string `json:"adjustmentId,omitempty"`
	Reason           *string `json:"reason,omitempty"`
}

type RefundOutcome string

const (
	EightyFivePercent  RefundOutcome = "EIGHTY_FIVE_PERCENT"
	EightyPercent      RefundOutcome = "EIGHTY_PERCENT"
	FifteenPercent     RefundOutcome = "FIFTEEN_PERCENT"
	FiftyFivePercent   RefundOutcome = "FIFTY_FIVE_PERCENT"
	FiftyPercent       RefundOutcome = "FIFTY_PERCENT"
	FivePercent        RefundOutcome = "FIVE_PERCENT"
	FortyFivePercent   RefundOutcome = "FORTY_FIVE_PERCENT"
	FortyPercent       RefundOutcome = "FORTY_PERCENT"
	NinetyFivePercent  RefundOutcome = "NINETY_FIVE_PERCENT"
	NinetyPercent      RefundOutcome = "NINETY_PERCENT"
	SeventyFivePercent RefundOutcome = "SEVENTY_FIVE_PERCENT"
	SeventyPercent     RefundOutcome = "SEVENTY_PERCENT"
	SixtyFivePercent   RefundOutcome = "SIXTY_FIVE_PERCENT"
	SixtyPercent       RefundOutcome = "SIXTY_PERCENT"
	TenPercent         RefundOutcome = "TEN_PERCENT"
	ThirtyFivePercent  RefundOutcome = "THIRTY_FIVE_PERCENT"
	ThirtyPercent      RefundOutcome = "THIRTY_PERCENT"
	TwentyFivePercent  RefundOutcome = "TWENTY_FIVE_PERCENT"
	TwentyPercent      RefundOutcome = "TWENTY_PERCENT"
)

type AdjustmentType string

const (
	Fee     AdjustmentType = "FEE"
	Intent  AdjustmentType = "INTENT"
	Loyalty AdjustmentType = "LOYALTY"
	Refund  AdjustmentType = "REFUND"
)

type Flow string

const (
	Inbound  Flow = "INBOUND"
	Internal Flow = "INTERNAL"
	Outbound Flow = "OUTBOUND"
)

type Purpose string

const (
	Deposit          Purpose = "DEPOSIT"
	ManualAdjustment Purpose = "MANUAL_ADJUSTMENT"
	Purchase         Purpose = "PURCHASE"
	Transfer         Purpose = "TRANSFER"
	Withdrawal       Purpose = "WITHDRAWAL"
)

type TransactionState string

const (
	Accepted   TransactionState = "ACCEPTED"
	Adjusted   TransactionState = "ADJUSTED"
	Cancelled  TransactionState = "CANCELLED"
	Disputed   TransactionState = "DISPUTED"
	Effected   TransactionState = "EFFECTED"
	Errored    TransactionState = "ERRORED"
	Failed     TransactionState = "FAILED"
	Held       TransactionState = "HELD"
	Pending    TransactionState = "PENDING"
	Processing TransactionState = "PROCESSING"
)

type Type string

const (
	TypeAdjustment Type = "adjustment"
)
