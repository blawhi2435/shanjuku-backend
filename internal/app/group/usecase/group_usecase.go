package usecase

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/jwt"
	"github.com/blawhi2435/shanjuku-backend/internal/util"
)

type groupUsecase struct {
	dbRepository    domain.DBRepository
	groupRepository domain.GroupPostgreRepository
	userRepository  domain.UserPostgreRepository
	userUsecase     domain.UserUsecase
}

func ProvideGroupUsecase(
	dbRepository domain.DBRepository,
	groupRepository domain.GroupPostgreRepository,
	userRepository domain.UserPostgreRepository,
	userUsecase domain.UserUsecase,
) domain.GroupUsecase {
	return &groupUsecase{
		dbRepository,
		groupRepository,
		userRepository,
		userUsecase,
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

func (g *groupUsecase) QueryByGroupID(ctx context.Context, groupID int64) (domain.Group, error) {

	if ok, err := g.IsGroupExist(ctx, groupID); err != nil {
		return domain.Group{}, err
	} else if !ok {
		return domain.Group{}, cerror.ErrGroupNotExist
	}

	group, err := g.groupRepository.QueryByID(ctx, groupID)
	if err != nil {
		return domain.Group{}, err
	}

	return group, nil
}

func (g *groupUsecase) QueryUserGroups (ctx context.Context, userID int64) ([]domain.Group, error) {


	if ok, err := g.userUsecase.IsUserExistByIDs(ctx, []int64{userID}); err != nil {
		return []domain.Group{}, err
	} else if !ok {
		return []domain.Group{}, cerror.ErrUserNotExist
	}

	groups, err := g.userRepository.QueryAssociationGroups(ctx, userID)
	if err != nil {
		return []domain.Group{}, err
	}

	return groups, nil
}

func (g *groupUsecase) QueryGroupUsers(ctx context.Context, groupID int64) ([]domain.User, error) {

	if ok, err := g.IsGroupExist(ctx, groupID); err != nil {
		return []domain.User{}, err
	} else if !ok {
		return []domain.User{}, cerror.ErrGroupNotExist
	}

	users, err := g.groupRepository.QueryAssociationUsers(ctx, groupID)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (g *groupUsecase) UpdateGroup(ctx context.Context, group *domain.Group) (domain.Group, error) {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return domain.Group{}, cerror.ErrTokenInvalid
	}

	if ok, err := g.IsGroupExist(ctx, group.ID); err != nil {
		return domain.Group{}, err
	} else if !ok {
		return domain.Group{}, cerror.ErrGroupNotExist
	}

	if ok, err := g.IsGroupHasUsers(ctx, group.ID, []int64{tokenPayload.UserID}); err != nil {
		return domain.Group{}, err
	} else if !ok {
		return domain.Group{}, cerror.ErrUserNotInGroup
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

	if ok, err := g.IsGroupExist(ctx, groupID); err != nil {
		return err
	} else if !ok {
		return cerror.ErrGroupNotExist
	}

	if ok, err := g.IsBelongToUser(ctx, groupID, tokenPayload.UserID); err != nil {
		return err
	} else if !ok {
		return cerror.ErrGroupNotBelongToUser
	}

	if err := g.dbRepository.Transaction(func(i interface{}) error {

		err := g.groupRepository.ClearUserInGroupWithTx(ctx, i, groupID)
		if err != nil {
			return err
		}

		err = g.groupRepository.DeleteGroupWithTx(ctx, i, groupID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (g *groupUsecase) InviteUser(ctx context.Context, groupID int64, userIDs []int64) ([]domain.User, error) {
	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return []domain.User{}, cerror.ErrTokenInvalid
	}

	if ok, err := g.IsGroupExist(ctx, groupID); err != nil {
		return []domain.User{}, err
	} else if !ok {
		return []domain.User{}, cerror.ErrGroupNotExist
	}

	if ok, err := g.IsGroupHasUsers(ctx, groupID, []int64{tokenPayload.UserID}); err != nil {
		return []domain.User{}, err
	} else if !ok {
		return []domain.User{}, cerror.ErrUserNotInGroup
	}

	if ok, err := g.userUsecase.IsUserExistByIDs(ctx, userIDs); err != nil {
		return []domain.User{}, err
	} else if !ok {
		return []domain.User{}, cerror.ErrUserNotExist
	}

	if ok, err := g.IsGroupHasUsers(ctx, groupID, userIDs); err != nil {
		return []domain.User{}, err
	} else if ok {
		return []domain.User{}, cerror.ErrUserAlreadyInGroup
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

	if userID == tokenPayload.UserID {
		return cerror.ErrCannotRemoveYourself
	}

	group, err := g.groupRepository.QueryByID(ctx, groupID)
	if err != nil {
		return err
	}

	if group.CreatorID == userID {
		return cerror.ErrCannotRemoveCreator
	}

	if ok, err := g.IsGroupHasUsers(ctx, groupID, []int64{userID}); err != nil {
		return err
	} else if !ok {
		return cerror.ErrUserNotInGroup
	}

	err = g.groupRepository.RemoveUserFromGroup(ctx, groupID, userID)
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

func (g *groupUsecase) IsGroupExist(ctx context.Context, groupID int64) (bool, error) {
	_, err := g.groupRepository.QueryByID(ctx, groupID)
	if err != nil {
		if err == cerror.ErrGroupNotExist {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (g *groupUsecase) IsGroupHasUsers(ctx context.Context, groupID int64, userID []int64) (bool, error) {
	users, err := g.groupRepository.QueryAssociationUsersWithSpecifyUserIDs(ctx, groupID, userID)
	if err != nil {
		return false, err
	}

	if len(users) != len(userID) {
		return false, nil
	}

	return true, nil
}
