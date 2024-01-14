package domain

import "context"

type User struct {
	ID       int64
	Account  string
	Password string
	Name     string
	Avatar   string
	Token    string
}

type UserPostgreRepository interface {
	InsertNewUser(ctx context.Context, user *User) error
	QueryByAccount(ctx context.Context, account string) (User, error)
	QueryByIDs(ctx context.Context, ids []int64) ([]User, error)
}

type UserUsecase interface {
	IsUserExistByIDs(ctx context.Context, id []int64) (bool, error)
	IsUserExistByAccount(ctx context.Context, account string) (bool, error)
}

type AuthUsecase interface {
	Register(ctx context.Context, user *User) (User, error)
	Login(ctx context.Context, account, password string) (User, error)
	Logout(ctx context.Context, account string) error
}
