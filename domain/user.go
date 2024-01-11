package domain

type User struct {
	ID       int64
	Account  string
	Password string
	Name     string
	Avatar   string
	Token    string
}

type UserPostgreRepository interface {
	InsertNewUser(user *User) error
	QueryByAccount(account string) (User, error)
}

type UserUsecase interface {
	IsUserExist(account string) (bool, error)
}

type AuthUsecase interface {
	Register(user *User) (User, error)
	Login(account, password string) (User, error)
	Logout(account string) error
}
