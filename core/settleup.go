package core

import (
	"fmt"
	"setu/dao"
	"setu/utils"
)

func SettleUpGroupTransactions(groupId, user1, user2, comment string) {
	transactions := dao.GetGroupTransactions(groupId)
	SettleUpTransactionsForUsers(groupId, user1, user2, comment, transactions)
}

func SettleUpUsersTransaction(user1, user2, comment string) {
	transactions := dao.GetTransactionForUsers(user1, user2)
	SettleUpTransactionsForUsers("", user1, user2, comment, transactions)
}

func SettleUpTransactionsForUsers(groupId, user1, user2, comment string, transactions []dao.Transaction) {
	currentTime := utils.GetUTCTime()
	settleUpRemark := fmt.Sprintf("Settle up for : %s", comment)
	var settleUpAmount float64
	settleUpTransaction, _ := dao.CreateTransaction(user1, groupId, settleUpAmount, settleUpRemark)
	var user1Amount, user2Amount float64
	for _, transaction := range transactions {
		splits := dao.GetTransactionSplits(transaction.ID)
		for _, split := range splits {
			if !split.SettledUp {
				if (split.User1 == user1 && split.User2 == user2) || (split.User1 == user2 && split.User2 == user1) {
					if split.User1 == user1 {
						user1Amount += split.Amount
					} else {
						user2Amount += split.Amount
					}
					split.SettledUp = true
					split.SettledOn = currentTime
					split.SettleComment = comment
					split.SettledUpTransactionID = settleUpTransaction.ID
					dao.TransactionSplitMap[split.TransactionID][split.ID] = split
				}
			}
		}
	}
	var splitUser1, splitUser2 string
	if user1Amount > user2Amount {
		settleUpAmount = user1Amount - user2Amount
		splitUser1 = user1
		splitUser2 = user2
	} else {
		settleUpAmount = user2Amount - user1Amount
		splitUser1 = user2
		splitUser2 = user1
	}
	settleUpTransaction.SettleMentTransaction = true
	settleUpTransaction.UpdateTransaction(user1, settleUpAmount, settleUpRemark)
	split := dao.Split{
		TransactionID: settleUpTransaction.ID,
		User1:         splitUser1,
		User2:         splitUser2,
		Amount:        user1Amount - user2Amount,
		SettledUp:     true,
		SettledOn:     currentTime,
		SettleComment: settleUpRemark,
	}
	split.AddOrUpdateTransactionSplit()
}
