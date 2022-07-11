package dao

var UsersMap = map[string]User{}

var UserConnections = map[string]UserConnection{}

var UserSessions = map[string]UserSession{}

var GroupMap = map[string]Group{}

var GroupUsersMap = map[string]GroupUsers{}

var TransactionMap = map[string]Transaction{}

var TransactionSplitMap = map[string](map[string]Split){}
