package apihandlers

import (
	"net/http"
	"setu/constants"
	"setu/core"
	"setu/dao"
	"setu/handler"
	"setu/responsedto"
)

type DashboardHandler struct {
	BaseHandler
}

func (u *DashboardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(u, r, []string{})
	response.RenderResponse(w)
}

// Get method for UserHandler
func (u *DashboardHandler) Get(r *http.Request) handler.ServiceResponse {
	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	groups := dao.GetUserGroups(user.ID)

	nonGroupbalances := core.GetUserNonGroupBalances(user.ID)
	groupbalances := []core.DashboardTile{}
	for _, group := range groups {
		tile := core.GetGroupBalancesForUser(group.ID, user.ID)
		tile.Group = &group
		groupbalances = append(groupbalances, tile)
	}
	return handler.Response200OK(responsedto.ConvertCoreDtoToResponse(nonGroupbalances, groupbalances))
}
