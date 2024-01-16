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
	ErrTokenInvalid              = Error{Code: "A0003", Message: "token invalid"}
	ErrActivityNotBelongToUser   = Error{Code: "AC0001", Message: "activity not belong to user"}
	ErrActivityNotExist          = Error{Code: "AC0002", Message: "activity not exist"}
	ErrGroupNotBelongToUser      = Error{Code: "G0001", Message: "group not belong to user"}
	ErrGroupNotExist             = Error{Code: "G0002", Message: "group not exist"}
	ErrUserAlreadyInGroup        = Error{Code: "G0003", Message: "user already in group"}
	ErrUserNotInGroup            = Error{Code: "G0004", Message: "user not in group"}
	ErrCannotRemoveYourself      = Error{Code: "G0005", Message: "cannot remove yourself"}
	ErrCannotRemoveCreator       = Error{Code: "G0006", Message: "cannot remove creator"}
	ErrUserNotExist              = Error{Code: "U0001", Message: "user not exist"}
	ErrInternalServerError       = Error{Code: "S0001", Message: "internal server error"}
	ErrGetContextFailed          = Error{Code: "S0002", Message: "get context failed"}
)

func GetGQLError(ctx context.Context, e error) error {
	var gqlErr gqlerror.Error

	switch e {
	case
		ErrAccountOrPasswordNotMatch,
		ErrAccountAlreadyExist,
		ErrTokenInvalid,
		ErrActivityNotBelongToUser,
		ErrActivityNotExist,
		ErrGroupNotBelongToUser,
		ErrGroupNotExist,
		ErrUserAlreadyInGroup,
		ErrUserNotInGroup,
		ErrCannotRemoveYourself,
		ErrCannotRemoveCreator,
		ErrUserNotExist,
		ErrGetContextFailed:

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
