package dao

import (
	"errors"
	"setu/utils"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                    string  `json:"id,omitempty"`
	GroupID               string  `json:"groupID,omitempty"`     // optional
	CreatedByID           string  `json:"createdByID,omitempty"` // user_id
	UpdatedByID           string  `json:"updatedBy,omitempty"`
	CreatedOn             int64   `json:"createdOn,omitempty"`
	UpdatedOn             int64   `json:"updatedOn,omitempty"`
	TransactionAmount     float64 `json:"transactionAmount,omitempty"`
	Remark                string  `json:"remark,omitempty"`
	SettleMentTransaction bool    `json:"settleMentTransaction,omitempty"`
}

type Split struct {
	ID                     string  `json:"id,omitempty"`
	TransactionID          string  `json:"transactionID,omitempty"`
	Amount                 float64 `json:"amount,omitempty"`
	User1                  string  `json:"user1,omitempty"`
	User2                  string  `json:"user2,omitempty"`
	SettledUp              bool    `json:"settledUp,omitempty"`
	SettledOn              int64   `json:"settledOn,omitempty"`
	SettleComment          string  `json:"settleComment,omitempty"`
	SettledUpTransactionID string  `json:"settledUpTransactionID,omitempty"`
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
	TransactionSplitMap[transaction.ID] = map[string]Split{}
	return &transaction, nil
}

func (t *Transaction) UpdateTransaction(userId string, amount float64, remark string) {
	t.UpdatedByID = userId
	t.TransactionAmount = amount
	t.Remark = remark
	t.UpdatedOn = utils.GetUTCTime()
	TransactionMap[t.ID] = *t
	TransactionSplitMap[t.ID] = map[string]Split{}
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

func DeleteTransaction(transaction Transaction) {
	if transaction.SettleMentTransaction {
		for transId, splitmap := range TransactionSplitMap {
			for splitId, split := range splitmap {
				if split.SettledUpTransactionID == transaction.ID {
					split.SettledUp = false
					split.SettleComment = ""
					split.SettledOn = 0
					split.SettledUpTransactionID = ""
					TransactionSplitMap[transId][splitId] = split
				}
			}
		}
	}
	delete(TransactionMap, transaction.ID)
	delete(TransactionSplitMap, transaction.ID)
}

func GetTransactionSplits(transactionId string) map[string]Split {
	return TransactionSplitMap[transactionId]
}

func GetGroupTransactions(groupId string) []Transaction {
	// Add Pagination
	transactions := []Transaction{}
	for _, t := range TransactionMap {
		if t.GroupID == groupId {
			transactions = append(transactions, t)
		}
	}
	return transactions
}

func GetNonGroupTransactions(userId string) []Transaction {
	// Add Pagination
	transactions := []Transaction{}
	for transactionId, splits := range TransactionSplitMap {
		IsUserTransaction := false
		for _, split := range splits {
			if split.User1 == userId || split.User2 == userId {
				IsUserTransaction = true
				break
			}
		}
		if IsUserTransaction {
			transaction, _ := GetTransactionByID(transactionId)
			if transaction.GroupID == "" {
				transactions = append(transactions, *transaction)
			}
		}
	}
	return transactions
}

func GetTransactionForUsers(user1, user2 string) []Transaction {
	transactions := []Transaction{}
	for _, t := range TransactionMap {
		if t.GroupID != "" {
			continue
		}
		if t.CreatedByID == user1 || t.CreatedByID == user2 {
			transactions = append(transactions, t)
		}
	}
	return transactions
}
