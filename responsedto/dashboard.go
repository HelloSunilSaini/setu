package responsedto

import (
	"setu/core"
	"setu/dao"
)

type UserWiseBalance struct {
	Amount float64              `json:"amount,omitempty"`
	User   UserCreationResponse `json:"user,omitempty"`
}
type DashboardTile struct {
	Owed             float64           `json:"owed,omitempty"`
	Owes             float64           `json:"owes,omitempty"`
	Balance          float64           `json:"balance,omitempty"`
	Group            *dao.Group        `json:"group,omitempty"`
	UserWiseBalances []UserWiseBalance `json:"userWiseBalances,omitempty"`
}

type DashboardResponse struct {
	NonGroupBalance DashboardTile   `json:"nonGroupBalance,omitempty"`
	GroupBalances   []DashboardTile `json:"groupBalances,omitempty"`
}

func ConvertCoreDtoToResponse(nongroupBalanse core.DashboardTile, groupBalances []core.DashboardTile) DashboardResponse {
	userWiseBalances := []UserWiseBalance{}
	for _, v := range nongroupBalanse.UserWiseBalances {
		userWiseBalances = append(userWiseBalances, UserWiseBalance{
			Amount: v.Amount,
			User:   *ConvertUserDtoToCreateResponseDto(v.User),
		})
	}
	nongroupBalanseResp := DashboardTile{
		Owed:             nongroupBalanse.Owed,
		Owes:             nongroupBalanse.Owes,
		Balance:          nongroupBalanse.Balance,
		UserWiseBalances: userWiseBalances,
	}
	groupBalancesResp := []DashboardTile{}
	for _, dashTile := range groupBalances {
		userWiseBalances := []UserWiseBalance{}
		for _, v := range dashTile.UserWiseBalances {
			userWiseBalances = append(userWiseBalances, UserWiseBalance{
				Amount: v.Amount,
				User:   *ConvertUserDtoToCreateResponseDto(v.User),
			})
		}
		groupBalanseResp := DashboardTile{
			Owed:             dashTile.Owed,
			Owes:             dashTile.Owes,
			Balance:          dashTile.Balance,
			Group:            dashTile.Group,
			UserWiseBalances: userWiseBalances,
		}
		groupBalancesResp = append(groupBalancesResp, groupBalanseResp)
	}
	return DashboardResponse{
		NonGroupBalance: nongroupBalanseResp,
		GroupBalances:   groupBalancesResp,
	}
}
