package usecase

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
)

type UserUsecase struct {
	userRepository domain.UserPostgreRepository
}

func ProvideUserUsecase(userRepository domain.UserPostgreRepository) domain.UserUsecase {
	return &UserUsecase{userRepository}
}

func (u *UserUsecase) IsUserExistByAccount(ctx context.Context, account string) (bool, error) {
	_, err := u.userRepository.QueryByAccount(ctx, account)
	if err != nil {
		if err == cerror.ErrUserNotExist {
			return false, nil
		} else {
			return false, err
		}
	}
	
	return true, nil
}

func (u *UserUsecase) IsUserExistByIDs(ctx context.Context, ids []int64) (bool, error) {
	users, err := u.userRepository.QueryByIDs(ctx, ids)
	if err != nil {
		return false, err
	}

	if len(users) != len(ids) {
		return false, nil
	}
	
	return true, nil
}