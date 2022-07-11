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

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	BaseHandler
}

func (u *TransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(u, r, []string{})
	response.RenderResponse(w)
}

func (u *TransactionHandler) Post(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var transactionData *requestdto.TransactionRequest
	err = json.Unmarshal(body, &transactionData)
	if err != nil {
		return handler.ProcessError(err)
	}

	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)

	if transactionData.GroupID != "" {
		_, err := dao.GetGroupById(transactionData.GroupID)
		if err != nil {
			return handler.SimpleBadRequest(err.Error())
		}
		if !dao.IsGroupUser(transactionData.GroupID, user.ID) {
			return handler.SimpleBadRequest("User Does not belongs to group")
		}
	}

	_, err = core.CreateTransaction(user, *transactionData)
	if err != nil {
		return handler.PreconditionFailed(err.Error())
	}

	return handler.Response200OK("Transaction created successfully")

}

func (u *TransactionHandler) Put(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var transactionData *requestdto.TransactionRequest
	err = json.Unmarshal(body, &transactionData)
	if err != nil {
		return handler.ProcessError(err)
	}

	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	transaction, err := dao.GetTransactionByID(transactionData.ID)
	if err != nil {
		return handler.Simple404Response(err.Error())
	}
	if transaction.GroupID != transactionData.GroupID {
		return handler.SimpleBadRequest("Transaction Group mismatch")
	}

	if transactionData.GroupID != "" {
		_, err := dao.GetGroupById(transactionData.GroupID)
		if err != nil {
			return handler.SimpleBadRequest(err.Error())
		}
		if !dao.IsGroupUser(transactionData.GroupID, user.ID) {
			return handler.SimpleBadRequest("User Does not belongs to group")
		}
	}

	_, err = core.UpdateTransaction(user, *transactionData)
	if err != nil {
		return handler.PreconditionFailed(err.Error())
	}

	return handler.Response200OK("Transaction created successfully")
}

// Get method for UserHandler
func (u *TransactionHandler) Get(r *http.Request) handler.ServiceResponse {

	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	vars := mux.Vars(r)
	groupId, ok := vars["groupId"]
	if !ok {
		groups := dao.GetUserGroups(user.ID)
		return handler.Response200OK(groups)
	}
	group, err := dao.GetGroupById(groupId)
	if err != nil {
		return handler.SimpleBadRequest(err.Error())
	}
	groupUsers := dao.GetGroupUsers(groupId)
	users := responsedto.ConvertDaoUsersToConnectionsResponse(groupUsers)
	resp := responsedto.SingleGroupResponse{
		GroupDetails: *group,
		Members:      users,
	}
	return handler.Response200OK(resp)
}
