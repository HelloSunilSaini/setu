package dao

import (
	"errors"
	"setu/utils"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"passwordHash,omitempty"`
	CreatedOn    int64  `json:"createdOn,omitempty"`
	UpdatedOn    int64  `json:"updatedOn,omitempty"`
	LastLogIn    int64  `json:"lastLogIn,omitempty"`
}

type UserSession struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"userID,omitempty"`
	CreatedOn int64  `json:"createdOn,omitempty"`
	Expiry    int64  `json:"expiry,omitempty"`
}

type UserConnection struct {
	UserID1   string `json:"userID1,omitempty"`
	UserID2   string `json:"userID2,omitempty"`
	CreatedOn int64  `json:"createdOn,omitempty"`
}

func CreateUser(name, email, password string) (*User, error) {
	id := uuid.New().String()
	user := User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: utils.GetPasswordHash(password, id),
		CreatedOn:    utils.GetUTCTime(),
	}
	// db entry
	UsersMap[id] = user
	return &user, nil
}

func (u *User) UpdateUser() {
	u.UpdatedOn = utils.GetUTCTime()
	UsersMap[u.ID] = *u
}

func (u *User) CreateUserSession(password string) (*UserSession, error) {
	// check password
	passwordHash := utils.GetPasswordHash(password, u.ID)
	if passwordHash != u.PasswordHash {
		return nil, errors.New("Invalid Password")
	}
	userSession := UserSession{
		ID:        uuid.New().String(),
		UserID:    u.ID,
		CreatedOn: utils.GetUTCTime(),
		Expiry:    utils.GetMillisByTime(time.Now().Add(time.Hour * 24 * 15)), // 15 days expiry
	}
	u.LastLogIn = userSession.CreatedOn
	u.UpdateUser()
	return &userSession, nil
}

func GetUserSession(sessionId string) (*UserSession, error) {
	userSession, ok := UserSessions[sessionId]
	if !ok {
		return nil, errors.New("Session Not Exist")
	}
	if time.UnixMilli(userSession.Expiry).Before(time.Now()) {
		return nil, errors.New("Session has expired, try login again")
	}
	return &userSession, nil
}

func GetUserByEmail(email string) (*User, error) {
	for _, v := range UsersMap {
		if v.Email == email {
			return &v, nil
		}
	}
	return nil, errors.New("User Not Found")
}

func GetUserByID(id string) (*User, error) {
	user, ok := UsersMap[id]
	if !ok {
		return nil, errors.New("User Not Found")
	}
	return &user, nil
}

func CreateUserConnection(user1Id, user2Id string) {
	uc := UserConnection{
		UserID1:   user1Id,
		UserID2:   user2Id,
		CreatedOn: utils.GetUTCTime(),
	}
	UserConnections[user1Id+"_"+user2Id] = uc
	UserConnections[user2Id+"_"+user1Id] = uc
	return
}

func GetUserConnections(userId string) []UserConnection {
	return nil
}
