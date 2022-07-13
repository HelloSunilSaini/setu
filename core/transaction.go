package core

import (
	"errors"
	"setu/dao"
	"setu/requestdto"
)

func CreateTransaction(user dao.User, req requestdto.TransactionRequest) (bool, error) {
	// validation
	extraPayers, lessPayers, extraPayersUserIds, lessPayersUserIds, err := getExtraAndLessMaps(req)
	if err != nil {
		return false, err
	}
	transaction, err := dao.CreateTransaction(user.ID, req.GroupID, req.Amount, req.Comment)
	if err != nil {
		return false, err
	}
	addSplits(transaction.ID, extraPayers, lessPayers, extraPayersUserIds, lessPayersUserIds)

	return true, nil
}

func UpdateTransaction(user dao.User, req requestdto.TransactionRequest) (bool, error) {
	extraPayers, lessPayers, extraPayersUserIds, lessPayersUserIds, err := getExtraAndLessMaps(req)
	if err != nil {
		return false, err
	}
	transaction, err := dao.GetTransactionByID(req.ID)
	if err != nil {
		return false, err
	}
	transaction.UpdateTransaction(user.ID, req.Amount, req.Comment)
	addSplits(transaction.ID, extraPayers, lessPayers, extraPayersUserIds, lessPayersUserIds)

	return true, nil
}

func getExtraAndLessMaps(req requestdto.TransactionRequest) (extraPayers, lessPayers map[string]float64, extraPayersUserIds, lessPayersUserIds []string, err error) {
	extraPayers = map[string]float64{}
	lessPayers = map[string]float64{}
	extraPayersUserIds = []string{}
	lessPayersUserIds = []string{}
	var extrapaid, lesspaid, totalpaid, totalowes float64
	for _, v := range req.SplitDetails {
		_, err = dao.GetUserByID(v.UserID)
		if err != nil {
			return
		}
		if req.GroupID != "" && !dao.IsGroupUser(req.GroupID, v.UserID) {
			err = errors.New("User(s) not belongs to group")
			return
		}
		if v.PaidAmount > v.OwesAmount {
			extraPayers[v.UserID] = v.PaidAmount - v.OwesAmount
			extraPayersUserIds = append(extraPayersUserIds, v.UserID)
			extrapaid += v.PaidAmount - v.OwesAmount
		}
		if v.OwesAmount > v.PaidAmount {
			lessPayers[v.UserID] = v.OwesAmount - v.PaidAmount
			lessPayersUserIds = append(lessPayersUserIds, v.UserID)
			lesspaid += v.OwesAmount - v.PaidAmount
		}
		totalpaid += v.PaidAmount
		totalowes += v.OwesAmount
	}
	if extrapaid != lesspaid || totalpaid != req.Amount || totalowes != req.Amount {
		err = errors.New("Amount Mismatch")
		return
	}
	return
}

func addSplits(transactionID string, extraPayers, lessPayers map[string]float64, extraPayersUserIds, lessPayersUserIds []string) {
	extraPayerIndex, lessPayerIndex := 0, 0
	for extraPayerIndex < len(extraPayersUserIds) && lessPayerIndex < len(lessPayersUserIds) {
		split := dao.Split{
			TransactionID: transactionID,
			User1:         extraPayersUserIds[extraPayerIndex],
			User2:         lessPayersUserIds[lessPayerIndex],
		}
		if extraPayers[extraPayersUserIds[extraPayerIndex]] > lessPayers[lessPayersUserIds[lessPayerIndex]] {
			split.Amount = lessPayers[lessPayersUserIds[lessPayerIndex]]
			extraPayers[extraPayersUserIds[extraPayerIndex]] -= lessPayers[lessPayersUserIds[lessPayerIndex]]
			lessPayers[lessPayersUserIds[lessPayerIndex]] = 0
			lessPayerIndex += 1
		} else if extraPayers[extraPayersUserIds[extraPayerIndex]] < lessPayers[lessPayersUserIds[lessPayerIndex]] {
			split.Amount = extraPayers[extraPayersUserIds[extraPayerIndex]]
			lessPayers[lessPayersUserIds[lessPayerIndex]] -= extraPayers[extraPayersUserIds[extraPayerIndex]]
			extraPayers[extraPayersUserIds[extraPayerIndex]] = 0
			extraPayerIndex += 1
		} else {
			split.Amount = extraPayers[extraPayersUserIds[extraPayerIndex]]
			lessPayers[lessPayersUserIds[lessPayerIndex]] = 0
			extraPayers[extraPayersUserIds[extraPayerIndex]] = 0
			extraPayerIndex += 1
			lessPayerIndex += 1
		}
		split.AddOrUpdateTransactionSplit()
	}
}
