package domain

import (
	"context"
)

type Group struct {
	ID        int64
	CreatorID int64
	Name      string
}

type GroupPostgreRepository interface {
	InsertGroup(ctx context.Context, group *Group) error
	QueryByID(ctx context.Context, groupID int64) (Group, error)
	QueryAssociationUsers(ctx context.Context, groupID int64) ([]User, error)
	QueryAssociationUsersWithSpecifyUserIDs(ctx context.Context, groupID int64, userID []int64) ([]User, error)
	UpdateGroup(ctx context.Context, group *Group) (int64, error)
	DeleteGroupWithTx(ctx context.Context, tx any, groupID int64) error
	AddUserToGroup(ctx context.Context, groupID int64, userID []int64) error
	RemoveUserFromGroup(ctx context.Context, groupID, userID int64) error
	ClearUserInGroupWithTx(ctx context.Context, tx any, groupID int64) error
}

type GroupUsecase interface {
	CreateGroup(ctx context.Context, group *Group) (Group, error)
	QueryByGroupID(ctx context.Context, groupID int64) (Group, error)
	QueryUserGroups(ctx context.Context, userID int64) ([]Group, error)
	QueryGroupUsers(ctx context.Context, groupID int64) ([]User, error)
	UpdateGroup(ctx context.Context, group *Group) (Group, error)
	DeleteGroup(ctx context.Context, groupID int64) error
	InviteUser(ctx context.Context, groupID int64, userIDs []int64) ([]User, error)
	RemoveUser(ctx context.Context, groupID, userID int64) error
	IsBelongToUser(ctx context.Context, groupID, userID int64) (bool, error)
	IsGroupExist(ctx context.Context, groupID int64) (bool, error)
	IsGroupHasUsers(ctx context.Context, groupID int64, userID []int64) (bool, error)
}
