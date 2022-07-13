package core

import (
	"setu/dao"
)

type UserWiseBalance struct {
	Amount float64   `json:"amount,omitempty"`
	User   *dao.User `json:"user,omitempty"`
}

type DashboardTile struct {
	Owed             float64           `json:"owed,omitempty"`
	Owes             float64           `json:"owes,omitempty"`
	Balance          float64           `json:"balance,omitempty"`
	Group            *dao.Group        `json:"group,omitempty"`
	UserWiseBalances []UserWiseBalance `json:"userWiseBalances,omitempty"`
}

func GetUserNonGroupBalances(userId string) DashboardTile {
	transactions := dao.GetNonGroupTransactions(userId)
	return CreateDashboardTile(userId, transactions)
}

func GetGroupBalancesForUser(groupId, userId string) DashboardTile {
	transactions := dao.GetGroupTransactions(groupId)
	return CreateDashboardTile(userId, transactions)
}

func CreateDashboardTile(userId string, transactions []dao.Transaction) DashboardTile {
	nongroupdash := DashboardTile{}
	userWisebalances := map[string]UserWiseBalance{}
	for _, transaction := range transactions {
		splits := dao.GetTransactionSplits(transaction.ID)
		for _, split := range splits {
			if (split.User1 == userId || split.User2 == userId) && !split.SettledUp {
				if split.User1 == userId {
					nongroupdash.Owed += split.Amount
					nongroupdash.Balance += split.Amount
					_, ok := userWisebalances[split.User2]
					if !ok {
						user, _ := dao.GetUserByID(split.User2)
						userWisebalances[split.User2] = UserWiseBalance{Amount: -split.Amount, User: user}
					} else {
						userbalance := userWisebalances[split.User2]
						userbalance.Amount -= split.Amount
						userWisebalances[split.User2] = userbalance
					}
				} else {
					nongroupdash.Owes += split.Amount
					nongroupdash.Balance -= split.Amount
					_, ok := userWisebalances[split.User2]
					if !ok {
						user, _ := dao.GetUserByID(split.User1)
						userWisebalances[split.User1] = UserWiseBalance{Amount: split.Amount, User: user}
					} else {
						userbalance := userWisebalances[split.User1]
						userbalance.Amount += split.Amount
						userWisebalances[split.User1] = userbalance
					}
				}
			}
		}
	}
	userWisebalancesList := []UserWiseBalance{}
	for _, v := range userWisebalances {
		userWisebalancesList = append(userWisebalancesList, v)
	}
	nongroupdash.UserWiseBalances = userWisebalancesList
	return nongroupdash
}
