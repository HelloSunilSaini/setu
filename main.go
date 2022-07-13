package main

import (
	"net/http"
	"setu/apihandlers"
	"setu/utils"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {
	logger := utils.Logger.Sugar()

	pingHandler := &apihandlers.PingHandler{}
	userHandler := &apihandlers.UserHandler{}
	userValidationHandler := &apihandlers.UserValidationHandler{}
	userConnectionHandler := &apihandlers.UserConnectionHandler{}
	groupHandler := &apihandlers.GroupHandler{}
	transactionHandler := &apihandlers.TransactionHandler{}
	dashboardHandler := &apihandlers.DashboardHandler{}
	settleUpHandler := &apihandlers.SettleUpHandler{}

	logger.Info("Setting up resources.")
	h := mux.NewRouter()

	h.Handle("/setusplitwise/ping/", pingHandler)
	h.Handle("/setusplitwise/user/", userHandler)
	h.Handle("/setusplitwise/uservalidation/", userValidationHandler)
	h.Handle("/setusplitwise/userconnection/", userConnectionHandler)
	h.Handle("/setusplitwise/group/", groupHandler)
	h.Handle("/setusplitwise/group/{groupId}/", groupHandler)
	h.Handle("/setusplitwise/group/{groupId}/transaction/{transactionId}/", transactionHandler)
	h.Handle("/setusplitwise/transaction/{transactionId}/", transactionHandler)
	h.Handle("/setusplitwise/transaction/", transactionHandler)
	h.Handle("/setusplitwise/dashboard/", dashboardHandler)
	h.Handle("/setusplitwise/settleup/", settleUpHandler)

	logger.Info("Resource Setup Done.")

	addr := ":7776"
	s := &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	s.ListenAndServe()
}

// - Implement BE for above solution
// - Readme file AND upload it to gitlab or github
// - Dockerfile <Docker compose>
// - API - OpenAPI standards, if possible you swagger or postman
// - Test cases
// - Clean code with standard practises
// - Makefile
