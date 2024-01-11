package usecase

import (
	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/environment"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/jwt"
	"github.com/blawhi2435/shanjuku-backend/util"
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

func (a *authUsecase) Register(user *domain.User) (domain.User, error) {

	isUserExist, err := a.userUsecase.IsUserExist(user.Account)
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

	err = a.userRepository.InsertNewUser(user)
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

func (a *authUsecase) Login(account, password string) (domain.User, error) {

	user, err := a.userRepository.QueryByAccount(account)
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

func (a *authUsecase) Logout(account string) error {
	
	isUserExist, err := a.userUsecase.IsUserExist(account)
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
