package txinternal

import (
	"github.com/bewolv/paybase-go/common"
	"github.com/bewolv/paybase-go/transaction"
)

type InternalOptions map[string]interface{}

type Create struct {
	FromAccountID string                        `json:"fromAccountId"`
	ToAccountID   string                        `json:"toAccountId"`
	Amount        string                        `json:"amount"`
	Method        string                        `json:"method"`
	TargetStateID *transaction.TransactionState `json:"targetStateId,omitempty"`
	Options       *InternalOptions              `json:"options,omitempty"`
	Purpose       common.Purpose                `json:"purpose"`
	Description   *string                       `json:"description,omitempty"`
	Annotations   []*string                     `json:"annotations,omitempty"`
	Tags          []*string                     `json:"tags,omitempty"`
	GroupID       *string                       `json:"groupId,omitempty"`
}
