package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type UserRepository interface {
	CheckUserExist(email string) bool
	FindUserByReference(ref string) (uint, error)
	FindUserByEmail(email string) (schema.UserSignInResponse, error)
	SignUp(user schema.UserSchema) (schema.UserSchemaResponse, error)
	CheckUserIsBlocked(email string) (bool, error)
	Create(user models.Users) (models.Users, error)
}
