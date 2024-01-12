package cerror

import (
	"context"
	"fmt"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s\n", e.Code, e.Message)
}

var (
	ErrAccountOrPasswordNotMatch = Error{Code: "A0001", Message: "account or password not match"}
	ErrAccountAlreadyExist       = Error{Code: "A0002", Message: "account already exist"}
	ErrTokenExpired              = Error{Code: "A0003", Message: "token expired"}
	ErrInternalServerError       = Error{Code: "S0001", Message: "internal server error"}
)

func GetGQLError(ctx context.Context, e error) error {
	var gqlErr gqlerror.Error

	switch e {
	case
		ErrAccountOrPasswordNotMatch,
		ErrAccountAlreadyExist,
		ErrTokenExpired:

		
		gqlErr.Err = e
		gqlErr.Message = e.(Error).Message
		gqlErr.Extensions = map[string]interface{}{
			"code": e.(Error).Code,
		}
	default:
		gqlErr.Err = e
		gqlErr.Message = ErrInternalServerError.Message
		gqlErr.Extensions = map[string]interface{}{
			"code": ErrInternalServerError.Code,
		}
	}

	errorList := gqlerror.List{}
	errorList = append(errorList, &gqlErr)
	
	return errorList
}
