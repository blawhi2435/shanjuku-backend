package usecase

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/jwt"
	"github.com/blawhi2435/shanjuku-backend/internal/util"
)

type groupUsecase struct {
	groupRepository domain.GroupPostgreRepository
	userRepository  domain.UserPostgreRepository
}

func ProvideGroupUsecase(
	groupRepository domain.GroupPostgreRepository,
	userRepository domain.UserPostgreRepository,
	) domain.GroupUsecase {
	return &groupUsecase{
		groupRepository,
		userRepository,
	}
}

func (g *groupUsecase) CreateGroup(ctx context.Context, group *domain.Group) (domain.Group, error) {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return domain.Group{}, cerror.ErrTokenInvalid
	}
	group.CreatorID = tokenPayload.UserID

	groupID, err := util.GenerateSnowflake()
	if err != nil {
		return domain.Group{}, err
	}

	group.ID = groupID
	err = g.groupRepository.InsertGroup(ctx, group)
	if err != nil {
		return domain.Group{}, err
	}

	return *group, nil
}

func (g *groupUsecase) UpdateGroup(ctx context.Context, group *domain.Group) (domain.Group, error) {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return domain.Group{}, cerror.ErrTokenInvalid
	}

	if ok, err := g.IsBelongToUser(ctx, group.ID, tokenPayload.UserID); err != nil {
		return domain.Group{}, err
	} else if !ok {
		return domain.Group{}, cerror.ErrGroupNotBelongToUser
	}

	_, err := g.groupRepository.UpdateGroup(ctx, group)
	if err != nil {
		return domain.Group{}, err
	}

	group.CreatorID = tokenPayload.UserID
	return *group, nil
}

func (g *groupUsecase) DeleteGroup(ctx context.Context, groupID int64) error {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return cerror.ErrTokenInvalid
	}

	if ok, err := g.IsBelongToUser(ctx, groupID, tokenPayload.UserID); err != nil {
		return err
	} else if !ok {
		return cerror.ErrGroupNotBelongToUser
	}

	err := g.groupRepository.DeleteGroup(ctx, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (g *groupUsecase) InviteUser(ctx context.Context, groupID int64, userIDs []int64) ([]domain.User, error) {
	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return []domain.User{}, cerror.ErrTokenInvalid
	}

	if ok, err := g.IsBelongToUser(ctx, groupID, tokenPayload.UserID); err != nil {
		return []domain.User{}, err
	} else if !ok {
		return []domain.User{}, cerror.ErrGroupNotBelongToUser
	}

	err := g.groupRepository.AddUserToGroup(ctx, groupID, userIDs)
	if err != nil {
		return []domain.User{}, err
	}

	users, err := g.groupRepository.QueryAssociationUsers(ctx, groupID)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (g *groupUsecase) RemoveUser(ctx context.Context, groupID, userID int64) error {
	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return cerror.ErrTokenInvalid
	}

	if ok, err := g.IsBelongToUser(ctx, groupID, tokenPayload.UserID); err != nil {
		return err
	} else if !ok {
		return cerror.ErrGroupNotBelongToUser
	}

	err := g.groupRepository.RemoveUserFromGroup(ctx, groupID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (g *groupUsecase) IsBelongToUser(ctx context.Context, groupID, userID int64) (bool, error) {

	group, err := g.groupRepository.QueryByID(ctx, groupID)
	if err != nil {
		return false, err
	}

	if group.CreatorID != userID {
		return false, nil
	}

	return true, nil
}
