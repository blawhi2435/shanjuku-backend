package domain

import (
	"context"
)

type Group struct {
	ID          int64
	CreatorID   int64
	Name        string
}

type GroupPostgreRepository interface {
	InsertGroup(ctx context.Context, group *Group) error
	QueryByID(ctx context.Context, groupID int64) (Group, error)
	QueryAssociationUsers(ctx context.Context, groupID int64) ([]User, error)
	UpdateGroup(ctx context.Context, group *Group) (int64, error)
	DeleteGroup(ctx context.Context, groupID int64) error
	AddUserToGroup(ctx context.Context, groupID int64, userID []int64) error
	RemoveUserFromGroup(ctx context.Context, groupID, userID int64) error
	ClearUserInGroup(ctx context.Context, groupID int64) error
}

type GroupUsecase interface {
	CreateGroup(ctx context.Context, group *Group) (Group, error)
	UpdateGroup(ctx context.Context, group *Group) (Group, error)
	DeleteGroup(ctx context.Context, groupID int64) error
	InviteUser(ctx context.Context, groupID int64, userIDs []int64) ([]User, error)
	RemoveUser(ctx context.Context, groupID, userID int64) error
	IsBelongToUser(ctx context.Context, groupID, userID int64) (bool, error)
}
