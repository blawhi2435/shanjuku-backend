package usecase

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"gorm.io/gorm"
)

type UserUsecase struct {
	userRepository domain.UserPostgreRepository
}

func ProvideUserUsecase(userRepository domain.UserPostgreRepository) domain.UserUsecase {
	return &UserUsecase{userRepository}
}

func (u *UserUsecase) IsUserExist(ctx context.Context, account string) (bool, error) {
	_, err := u.userRepository.QueryByAccount(ctx, account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	
	return true, nil
}