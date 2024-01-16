package usecase

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/jwt"
	"github.com/blawhi2435/shanjuku-backend/internal/util"
)

type activityUsecase struct {
	dbRepository               domain.DBRepository
	activityPosygresRepository domain.ActivityPostgreRepository
	groupUsecase               domain.GroupUsecase
}

func ProvideActivityUsecase(
	dbRepository domain.DBRepository,
	activityPosygresRepository domain.ActivityPostgreRepository,
	groupUsecase domain.GroupUsecase,
) domain.ActivityUsecase {
	return &activityUsecase{
		dbRepository,
		activityPosygresRepository,
		groupUsecase,
	}
}

func (a *activityUsecase) CreateActivity(ctx context.Context, activity *domain.Activity) (domain.Activity, error) {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return domain.Activity{}, cerror.ErrTokenInvalid
	}
	activity.CreatorID = tokenPayload.UserID

	activityID, err := util.GenerateSnowflake()
	if err != nil {
		return domain.Activity{}, err
	}
	activity.ID = activityID

	err = a.activityPosygresRepository.InsertActivity(ctx, activity)
	if err != nil {
		return domain.Activity{}, err
	}

	return *activity, nil
}

func (a *activityUsecase) QueryByID(ctx context.Context, activityID int64) (domain.Activity, error) {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return domain.Activity{}, cerror.ErrTokenInvalid
	}

	activity, err := a.activityPosygresRepository.QueryByID(ctx, activityID)
	if err != nil {
		return domain.Activity{}, err
	}

	if ok, err := a.groupUsecase.IsGroupHasUsers(ctx, activity.GroupID,
		[]int64{tokenPayload.UserID}); err != nil {
		return domain.Activity{}, err
	} else if !ok {
		return domain.Activity{}, cerror.ErrUserNotInGroup
	}

	return activity, nil
}

func (a *activityUsecase) UpdateActivityName(ctx context.Context,
	activity *domain.Activity) (domain.Activity, error) {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return domain.Activity{}, cerror.ErrTokenInvalid
	}

	currentActivity, err := a.activityPosygresRepository.QueryByID(ctx, activity.ID)
	if err != nil {
		return domain.Activity{}, err
	}

	if ok, err := a.groupUsecase.IsGroupHasUsers(ctx, currentActivity.GroupID,
		[]int64{tokenPayload.UserID}); err != nil {
		return domain.Activity{}, err
	} else if !ok {
		return domain.Activity{}, cerror.ErrUserNotInGroup
	}

	_, err = a.activityPosygresRepository.UpdateActivityName(ctx, activity)
	if err != nil {
		return domain.Activity{}, err
	}

	currentActivity.ActivityName = activity.ActivityName

	return currentActivity, nil
}

func (a *activityUsecase) DeleteActivity(ctx context.Context, activityID int64) error {

	tokenPayload, ok := jwt.ValidateTokenAndGetPayload(ctx)
	if !ok {
		return cerror.ErrTokenInvalid
	}

	if ok, err := a.groupUsecase.IsBelongToUser(ctx, activityID, tokenPayload.UserID); err != nil {
		return err
	} else if !ok {
		return cerror.ErrActivityNotBelongToUser
	}

	if err := a.dbRepository.Transaction(func(i interface{}) error {

		err := a.activityPosygresRepository.DeleteActivityWithTx(ctx, i, activityID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (a *activityUsecase) IsBelongToUser(ctx context.Context, activityID, userID int64) (bool, error) {

	activity, err := a.activityPosygresRepository.QueryByID(ctx, activityID)
	if err != nil {
		return false, err
	}

	if activity.CreatorID != userID {
		return false, nil
	}

	return true, nil
}

func (a *activityUsecase) IsActivityExist(ctx context.Context, activityID int64) (bool, error) {
	_, err := a.activityPosygresRepository.QueryByID(ctx, activityID)
	if err != nil {
		if err == cerror.ErrGroupNotExist {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}