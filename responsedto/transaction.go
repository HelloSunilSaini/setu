package responsedto

import (
	"fmt"
	"setu/dao"
)

type TransactionDetailsResponse struct {
	TransactionId string
	GroupDetails  *dao.Group
	CreatedBy     *UserCreationResponse
	UpdatedBy     *UserCreationResponse
	CreatedOn     int64
	UpdatedOn     int64
	Amount        float64
	Remark        string
	Splits        []TransactionSplit
}

type TransactionSplit struct {
	Amount        float64
	User1         *UserCreationResponse
	User2         *UserCreationResponse
	SettledUp     bool
	SettledOn     int64
	SettleComment string
	Statement     string
}

func ConvertDaoTransactionsToResponse(transaction *dao.Transaction, splits map[string]dao.Split) TransactionDetailsResponse {
	group, _ := dao.GetGroupById(transaction.GroupID)
	createdBy, _ := dao.GetUserByID(transaction.CreatedByID)
	updatedBy, _ := dao.GetUserByID(transaction.UpdatedByID)
	transactionDetailsResponse := TransactionDetailsResponse{
		TransactionId: transaction.ID,
		GroupDetails:  group,
		CreatedBy:     ConvertUserDtoToCreateResponseDto(createdBy),
		UpdatedBy:     ConvertUserDtoToCreateResponseDto(updatedBy),
		CreatedOn:     transaction.CreatedOn,
		UpdatedOn:     transaction.UpdatedOn,
		Amount:        transaction.TransactionAmount,
		Remark:        transaction.Remark,
	}
	tSplits := []TransactionSplit{}
	for _, split := range splits {
		user1, _ := dao.GetUserByID(split.User1)
		user2, _ := dao.GetUserByID(split.User2)
		tsplit := TransactionSplit{
			Amount:        split.Amount,
			User1:         ConvertUserDtoToCreateResponseDto(user1),
			User2:         ConvertUserDtoToCreateResponseDto(user2),
			SettledUp:     split.SettledUp,
			SettledOn:     split.SettledOn,
			SettleComment: split.SettleComment,
			Statement:     fmt.Sprintf("%s owes %s %v", user2.Name, user1.Name, split.Amount),
		}
		tSplits = append(tSplits, tsplit)
	}
	transactionDetailsResponse.Splits = tSplits
	return transactionDetailsResponse
}
