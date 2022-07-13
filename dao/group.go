package dao

import (
	"errors"
	"setu/utils"

	"github.com/google/uuid"
)

type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedOn   int64  `json:"createdOn,omitempty"`
	CreatedByID string `json:"createdByID,omitempty"`
	UpdatedOn   int64  `json:"updatedOn,omitempty"`
	UpdatedByID string `json:"updatedByID,omitempty"`
}

type GroupUsers struct {
	UserID  string `json:"userID,omitempty"`
	GroupID string `json:"groupID,omitempty"`
}

func GetGroupById(groupId string) (*Group, error) {
	group, ok := GroupMap[groupId]
	if !ok {
		return nil, errors.New("group not found")
	}
	return &group, nil
}

func CreateGroup(name, description, createdByID string) (*Group, error) {
	id := uuid.New().String()
	group := Group{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedOn:   utils.GetUTCTime(),
		CreatedByID: createdByID,
	}
	// db entry
	GroupMap[id] = group
	return &group, nil
}

func AddGroupUsers(groupId, userId string) error {
	_, ok := GroupMap[groupId]
	if !ok {
		return errors.New("Group Not Exist")
	}
	GroupUsersMap[groupId+"_"+userId] = GroupUsers{
		GroupID: groupId,
		UserID:  userId,
	}
	return nil
}

func RemoveGroupUser(groupId, userId string) {
	delete(GroupUsersMap, groupId+"_"+userId)
}

func IsGroupUser(groupId, userId string) bool {
	_, ok := GroupUsersMap[groupId+"_"+userId]
	if !ok {
		return false
	}
	return true
}

func GetGroupUsers(groupId string) []User {
	groupUsers := []User{}
	for _, groupuser := range GroupUsersMap {
		if groupuser.GroupID == groupId {
			user, err := GetUserByID(groupuser.UserID)
			if err != nil {
				groupUsers = append(groupUsers, *user)
			}
		}
	}
	return groupUsers
}

func GetUserGroups(userId string) []Group {
	groupIds := map[string]bool{}
	for _, groupUser := range GroupUsersMap {
		if groupUser.UserID == userId {
			groupIds[groupUser.GroupID] = true
		}
	}
	groups := []Group{}
	for groupId, _ := range groupIds {
		group, _ := GetGroupById(groupId)
		groups = append(groups, *group)
	}
	return groups
}
