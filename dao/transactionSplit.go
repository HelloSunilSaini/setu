package dao

import (
	"errors"
	"setu/utils"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                string  `json:"id,omitempty"`
	GroupID           string  `json:"groupID,omitempty"`     // optional
	CreatedByID       string  `json:"createdByID,omitempty"` // user_id
	CreatedOn         int64   `json:"createdOn,omitempty"`
	UpdatedOn         int64   `json:"updatedOn,omitempty"`
	TransactionAmount float64 `json:"transactionAmount,omitempty"`
	Remark            string  `json:"remark,omitempty"`
}

type Split struct {
	ID            string  `json:"id,omitempty"`
	TransactionID string  `json:"transactionID,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	User1         string  `json:"user1,omitempty"`
	User2         string  `json:"user2,omitempty"`
	SettledUp     bool    `json:"settledUp,omitempty"`
	SettledOn     uint64  `json:"settledOn,omitempty"`
	SettleComment string  `json:"settleComment,omitempty"`
}

func CreateTransaction(userId, groupId string, amount float64, remark string) (*Transaction, error) {
	transaction := Transaction{
		ID:                uuid.New().String(),
		GroupID:           groupId,
		CreatedByID:       userId,
		CreatedOn:         utils.GetUTCTime(),
		TransactionAmount: amount,
		Remark:            remark,
	}
	TransactionMap[transaction.ID] = transaction
	return &transaction, nil
}

func (s *Split) AddOrUpdateTransactionSplit() {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	TransactionSplitMap[s.TransactionID][s.ID] = *s
}

func GetTransactionByID(transactionId string) (*Transaction, error) {
	transaction, ok := TransactionMap[transactionId]
	if !ok {
		return nil, errors.New("Transaction not found")
	}
	return &transaction, nil
}

// func
