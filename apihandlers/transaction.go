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
	transactionId, ok := vars["transactionId"]
	if !ok {
		// params pagenumber pagesize
		groupId, ok := vars["groupId"]
		if !ok {
			// get user non-group transaction
			transactions := dao.GetNonGroupTransactions(user.ID)
			return handler.Response200OK(transactions)
		}
		_, err := dao.GetGroupById(groupId)
		if err != nil {
			return handler.SimpleBadRequest("Invalid Group")
		}
		if !dao.IsGroupUser(groupId, user.ID) {
			return handler.SimpleBadRequest("User is not member of group")
		}
		transactions := dao.GetGroupTransactions(groupId)
		return handler.Response200OK(transactions)
	}
	transaction, err := dao.GetTransactionByID(transactionId)
	if err != nil {
		return handler.SimpleBadRequest(err.Error())
	}
	splits := dao.GetTransactionSplits(transactionId)

	return handler.Response200OK(responsedto.ConvertDaoTransactionsToResponse(transaction, splits))
}

func (u *TransactionHandler) Delete(r *http.Request) handler.ServiceResponse {
	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	vars := mux.Vars(r)
	transactionId, ok := vars["transactionId"]
	if !ok {
		return handler.SimpleBadRequest("transactionId not given")
	}
	transaction, err := dao.GetTransactionByID(transactionId)
	if err != nil {
		return handler.SimpleBadRequest(err.Error())
	}
	if transaction.CreatedByID != user.ID {
		return handler.SimpleBadRequest("Can not delete transaction created by other user")
	}
	dao.DeleteTransaction(*transaction)
	return handler.Response200OK("Deleted transaction successfully")
}
