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
	"setu/responsedto"
)

type UserConnectionHandler struct {
	BaseHandler
}

func (u *UserConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(u, r, []string{})
	response.RenderResponse(w)
}

func (u *UserConnectionHandler) Post(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var userData *requestdto.UserConnectionRequest
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return handler.ProcessError(err)
	}

	// Validate User
	user2, err := dao.GetUserByEmail(userData.Email)
	if err != nil {
		return handler.SimpleBadRequest("User does not exist with Provided Email !!")
	}

	user1 := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	dao.CreateUserConnection(user1.ID, user2.ID)

	return handler.Response200OK("Connection Created")
}

func (u *UserConnectionHandler) Get(r *http.Request) handler.ServiceResponse {
	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	userConnections := core.GetUserConnection(user.ID)
	return handler.Response200OK(responsedto.ConvertDaoUsersToConnectionsResponse(userConnections))
}
