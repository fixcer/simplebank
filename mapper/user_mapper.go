package mapper

import (
	"github.com/fixcer/simplebank/api/model"
	db "github.com/fixcer/simplebank/db/sqlc"
)

func ToUserResponse(user db.User) model.UserResponse {
	return model.UserResponse{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: user.PasswordChangedAt,
		CreatedAt:        user.CreatedAt,
	}
}
