package apihandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"setu/constants"
	"setu/dao"
	"setu/handler"
	"setu/requestdto"
	"setu/responsedto"
	"setu/utils"

	"github.com/gorilla/mux"
)

type GroupHandler struct {
	BaseHandler
}

func (u *GroupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(u, r, []string{})
	response.RenderResponse(w)
}

func (u *GroupHandler) Post(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var groupData *requestdto.GroupCreationRequest
	err = json.Unmarshal(body, &groupData)
	if err != nil {
		return handler.ProcessError(err)
	}

	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	// Create group
	group, err := dao.CreateGroup(groupData.Name, groupData.Description, user.ID)
	if err != nil {
		return handler.ReponseInternalError(err.Error())
	}
	for _, userId := range groupData.Users {
		dao.AddGroupUsers(group.ID, userId)
	}
	return handler.Response200OK(group)

}

func (u *GroupHandler) Patch(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var groupData *requestdto.GroupUsersRequest
	err = json.Unmarshal(body, &groupData)
	if err != nil {
		return handler.ProcessError(err)
	}
	_, err = dao.GetGroupById(groupData.GroupID)
	if err != nil {
		return handler.SimpleBadRequest(err.Error())
	}
	for _, userId := range groupData.Users {
		_, err = dao.GetUserByID(userId)
		if err != nil {
			return handler.SimpleBadRequest("Invalid UserIds Provided")
		}
	}
	switch groupData.Action {
	case "ADD":
		for _, userId := range groupData.Users {
			dao.AddGroupUsers(groupData.GroupID, userId)
		}
	case "REMOVE":
		for _, userId := range groupData.Users {
			dao.RemoveGroupUser(groupData.GroupID, userId)
		}
	}

	return handler.Simple200OK("Operation Sussess")
}

// Get method for UserHandler
func (u *GroupHandler) Get(r *http.Request) handler.ServiceResponse {
	logger := utils.Logger.Sugar()
	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	vars := mux.Vars(r)
	logger.Info(vars)
	groupId, ok := vars["groupId"]
	if !ok {
		groups := dao.GetUserGroups(user.ID)
		return handler.Response200OK(groups)
	}
	group, err := dao.GetGroupById(groupId)
	if err != nil {
		return handler.SimpleBadRequest(err.Error())
	}
	logger.Info(group)
	groupUsers := dao.GetGroupUsers(groupId)
	logger.Info(groupUsers)
	users := responsedto.ConvertDaoUsersToConnectionsResponse(groupUsers)
	resp := responsedto.SingleGroupResponse{
		GroupDetails: *group,
		Members:      users,
	}
	return handler.Response200OK(resp)
}
