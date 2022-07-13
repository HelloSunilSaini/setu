package apihandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"setu/constants"
	"setu/core"
	"setu/dao"
	"setu/handler"
	"setu/requestdto"
)

type SettleUpHandler struct {
	BaseHandler
}

func (u *SettleUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(u, r, []string{})
	response.RenderResponse(w)
}

func (u *SettleUpHandler) Post(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var settleUpRequest *requestdto.SettleUpRequest
	err = json.Unmarshal(body, &settleUpRequest)
	if err != nil {
		return handler.ProcessError(err)
	}
	user1 := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	user2, err := dao.GetUserByID(settleUpRequest.UserID)
	if err != nil {
		return handler.SimpleBadRequest("Invalid UserId")
	}
	if user1.ID == user2.ID {
		return handler.SimpleBadRequest("userId provided of same user")
	}
	if settleUpRequest.GroupID != "" {
		_, err := dao.GetGroupById(settleUpRequest.GroupID)
		if err != nil {
			return handler.SimpleBadRequest("Invalid GroupId")
		}
		core.SettleUpGroupTransactions(settleUpRequest.GroupID, user1.ID, user2.ID, settleUpRequest.Comment)
	}
	core.SettleUpUsersTransaction(user1.ID, user2.ID, settleUpRequest.Comment)
	return handler.Response200OK("settleUp successfully")

}
