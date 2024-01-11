package cerror

import "fmt"

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s\n", e.Code, e.Message)
}

var (
	ErrAccountOrPasswordNotMatch = Error{Code: "A0001", Message: "account or password not match"}
	ErrAccountAlreadyExist			 = Error{Code: "A0002", Message: "account already exist"}
)
