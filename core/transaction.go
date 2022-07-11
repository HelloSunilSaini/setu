package core

import (
	"setu/dao"
	"setu/requestdto"
)

func CreateTransaction(user dao.User, req requestdto.TransactionRequest) (bool, error) {
	// validation
	_, err := dao.CreateTransaction(user.ID, req.GroupID, req.Amount, req.Comment)
	if err != nil {
		return false, err
	}
	extraPayers := map[string]float64{}
	lessPayers := map[string]float64{}
	for _, v := range req.SplitDetails {
		if v.PaidAmount > v.OwesAmount {
			extraPayers[v.UserID] = v.PaidAmount - v.OwesAmount
		}
		if v.OwesAmount > v.PaidAmount {
			lessPayers[v.UserID] = v.OwesAmount - v.PaidAmount
		}
	}
	// extraPayerIndex, lessPayerIndex := 0, 0
	// for extraPayerIndex < len(extraPayers) && lessPayerIndex < len(lessPayerIndex) {

	// 	split := dao.Split{
	// 		TransactionID: transaction.ID,
	// 		Amount: extraPayers[]

	// 	}
	// }
	return true, nil
}

// Amount        float64 `json:"amount,omitempty"`
// User1         string  `json:"user1,omitempty"`
// User2         string  `json:"user2,omitempty"`
// SettledUp     bool    `json:"settledUp,omitempty"`
// SettledOn     uint64  `json:"settledOn,omitempty"`
// SettleComment string  `json:"settleComment,omitempty"`
