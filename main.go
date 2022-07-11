package main

import (
	"net/http"
	"setu/apihandlers"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/

	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {
	// godotenv.Load()
	pingHandler := &apihandlers.PingHandler{}

	// logger.WithContext(context.TODO()).Infof("Setting up resources.")
	h := mux.NewRouter()

	h.Handle("/splitwise/ping/", pingHandler)

	// logger.WithContext(context.TODO()).Infof("Resource Setup Done.")

	addr := ":7776"
	s := &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	s.ListenAndServe()
}

//  user SignUp Post (email, name, password)
// password -> MD5(userId, password)[:25] -> passwordHash

// user Signin Post Email, password

// Add Connection (UserID, ConnectionEmail)

// Get Connections (UserID)

// Create Group Post (groupName, description, []users)

// - Add User to Group (groupID, UserID)

// Get Group, Get Groups

// Add Transaction Non-Group (Amount, Description,  SplitDetails[{user, PaidAmount, SplitAmount}]}) POST
// Personal Dashboard (userID)
// {
// 	TotalOwed : 2000,
// 	Name : "sunil"
// 	GroupDetails: [
// 		{
// 			GroupID:
// 			Name:
// 			Amount: -2000
// 		},
// 	]
// }

// Sunil owed praveen 1000
// Sunil owed praveen1 2000
// sunil owe praveen2 1000

// // Non- Group
// {
// 	BalanceList : [
// 		{
// 			User2Name : "praveen"
// 			User2ID : "27e128et8"
// 			Amount : 1000
// 		},{
// 			User2Name : "praveen1"
// 			Amount : 2000
// 		},{
// 			User2Name : "praveen2"
// 			Amount : -1000
// 		}
// 	],
// 	"Transaction History" : {
// 		page: 1,
// 		Count: 10
// 		transactions: [

// 		],
// }

// // UserWise Transaction Histoty GET
// {
// 	Balance : +/-
// 	GroupSumrries [
// 		{
// 			GroupID :
// 			GroupName:
// 			Amount : -/+
// 		},
// 		{

// 		}
// 	]
// 	TotalIndependantTransactionsAmount : +/-
// 	user transactions :[

// 	]
// }

// // Group Dashboard Get (UserID, GroupID)
// {
// 	GroupDetails: name, Description,Id
// 	GroupBalances: [
// 		{
// 			UserID:
// 			Name:
// 			Amount: +/-
// 		},

// 	]
// 	GroupTransaction: [
// 		{
// 			TransactionID :
// 			User2: id,Name
// 			Amount:
// 			Comment:
// 			AddedOn:
// 		}
// 	]
// }

// // Delete Transaction

// //

// // Group A
// sunil -> praveen 500

// // Independent
// praveen -> sunil 200

// // Group A
// praveen -> sunil 300

// Personal Dashboard
// sunil -> praveen 300

// Group A Dashboard

// case 1

// case 2  // Group Shubham
// // Group A
// Shubham -> sunil 5000 // for flight ticket
// // Group A
// sunil -> praveen 500  X // for food

// // Independent
// praveen -> sunil 200
// // Group A
// praveen -> sunil 300

// ----------------------------
// Sunil

// Non-Group expences
// you are Owed $2,849.59

// AgrostarPucsd
// you owe $1,313.92

// User1 User2
// 200    300
// 150    350

// User1 User2
// 200    300
// 50%    50%

// - Implement BE for above solution
// - Readme file AND upload it to gitlab or github
// - Dockerfile <Docker compose>
// - API - OpenAPI standards, if possible you swagger or postman
// - Test cases
// - Clean code with standard practises
// - Makefile
