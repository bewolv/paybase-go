package txexternal

import (
	"github.com/bewolv/paybase-go/bankaccount"
	"github.com/bewolv/paybase-go/common"
	"github.com/bewolv/paybase-go/transaction"
)

type BankAccount struct {
	BankAccount *bankaccount.Create `json:"bankAccount,omitempty"`
	Type        *string             `json:"type,omitempty"`
}

type Card struct {
	Card   map[string]interface{} `json:"card"`
	Method *string                `json:"method,omitempty"`
}

type Options struct {
	Outcome *transaction.RefundOutcome `json:"outcome,omitempty"`
	Card
}

type PaymentInstrument struct {
	*BankAccount
	*Card
}

type CreateOutbound struct {
	AccountID           string                        `json:"accountId" validate:"required"`
	Amount              string                        `json:"amount"`
	Method              string                        `json:"method"`
	TargetStateID       *transaction.TransactionState `json:"targetStateId,omitempty"`
	Purpose             common.Purpose                `json:"purpose"`
	PaymentInstrumentID *string                       `json:"paymentInstrumentId,omitempty"`
	Description         string                        `json:"description"`
	Annotations         *map[string]string            `json:"annotations,omitempty"`
	Tags                []*string                     `json:"tags,omitempty"`
	GroupID             *string                       `json:"groupId,omitempty"`
	Reference           *string                       `json:"reference,omitempty"`
}

type CreateInbound struct {
	AccountID           string                        `json:"accountId" validate:"required"`
	Amount              string                        `json:"amount"`
	Method              string                        `json:"method"`
	TargetStateID       *transaction.TransactionState `json:"targetStateId,omitempty"`
	Options             *Options                      `json:"options,omitempty"`
	Purpose             common.Purpose                `json:"purpose"`
	PaymentInstrumentID *string                       `json:"paymentInstrumentId,omitempty"`
	Description         string                        `json:"description"`
	Annotations         *map[string]string            `json:"annotations,omitempty"`
	Tags                []*string                     `json:"tags,omitempty"`
	GroupID             *string                       `json:"groupId,omitempty"`
}
