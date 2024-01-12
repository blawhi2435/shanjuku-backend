package usecase

import (
	"context"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/environment"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/jwt"
	"github.com/blawhi2435/shanjuku-backend/internal/util"
	"gorm.io/gorm"
)

type authUsecase struct {
	userRepository domain.UserPostgreRepository
	userUsecase    domain.UserUsecase
}

func ProvideAuthUsecase(
	userRepository domain.UserPostgreRepository,
	userUsecase domain.UserUsecase,
) domain.AuthUsecase {
	return &authUsecase{
		userRepository,
		userUsecase,
	}
}

func (a *authUsecase) Register(ctx context.Context, user *domain.User) (domain.User, error) {

	isUserExist, err := a.userUsecase.IsUserExist(ctx, user.Account)
	if err != nil {
		return domain.User{}, err
	}
	if isUserExist {
		return domain.User{}, cerror.ErrAccountAlreadyExist
	}

	passwordWithPrefix := getPasswordWithPrefix(user.Password)
	passwordHashed := util.SHA256(passwordWithPrefix)

	id, err := util.GenerateSnowflake()
	if err != nil {
		return domain.User{}, err
	}

	user.ID = id
	user.Password = passwordHashed

	err = a.userRepository.InsertNewUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	token, err := jwt.GenToken(jwt.PayloadData{
		UserID:  user.ID,
		Account: user.Account,
	})

	user.Token = token

	return *user, nil
}

func (a *authUsecase) Login(ctx context.Context, account, password string) (domain.User, error) {

	user, err := a.userRepository.QueryByAccount(ctx, account)
	if err == gorm.ErrRecordNotFound {
		return domain.User{}, cerror.ErrAccountOrPasswordNotMatch
	}

	if err != nil {
		return domain.User{}, err
	}

	passwordWithPrefix := getPasswordWithPrefix(password)
	passwordHashed := util.SHA256(passwordWithPrefix)

	if user.Password != passwordHashed {
		return domain.User{}, cerror.ErrAccountOrPasswordNotMatch
	}

	return user, nil
}

func (a *authUsecase) Logout(ctx context.Context, account string) error {
	
	isUserExist, err := a.userUsecase.IsUserExist(ctx, account)
	if err != nil {
		return err
	}
	if !isUserExist {
		return cerror.ErrAccountOrPasswordNotMatch
	}

	return nil
}

func getPasswordWithPrefix(password string) string {
	return environment.Setting.Auth.PasswordPrefix + password
}
