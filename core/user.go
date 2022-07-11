package core

import "setu/dao"

func GetUserConnection(userId string) []dao.User {
	userConnections := dao.GetUserConnections(userId)
	userConnectionDetails := []dao.User{}
	for _, v := range userConnections {
		if v.UserID1 != userId {
			user, _ := dao.GetUserByID(v.UserID1)
			userConnectionDetails = append(userConnectionDetails, *user)
		}
		if v.UserID2 != userId {
			user, _ := dao.GetUserByID(v.UserID2)
			userConnectionDetails = append(userConnectionDetails, *user)
		}
	}
	return userConnectionDetails
}
