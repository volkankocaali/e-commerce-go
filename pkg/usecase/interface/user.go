package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type UserUseCase interface {
	UserSignUp(user schema.UserSchema, ref string) (schema.TokenUsers, error)
	UserLogin(user schema.UserLoginSchema) (schema.TokenUsersLogin, error)
}
