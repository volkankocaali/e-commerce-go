package schema

import (
	"github.com/volkankocaali/e-commorce-go/pkg/validation"
	"time"
)

type UserSchema struct {
	Id                   int    `json:"id" form:"id"`
	Name                 string `json:"name" form:"name" validate:"required,min=5,max=40"`
	Email                string `json:"email" form:"email" validate:"required,email"`
	Phone                string `json:"phone" form:"phone" validate:"required,min=10,max=10"`
	ReferralCode         string `json:"referral_code" form:"referral_code"`
	Password             string `json:"password" form:"password" validate:"required,min=8,max=20"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,min=8,max=20"`
}

type UserLoginSchema struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserSchemaResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserSignInResponse struct {
	Id        uint                    `json:"id"`
	Name      string                  `json:"name"`
	Email     string                  `json:"email"`
	Password  string                  `json:"password"`
	Phone     string                  `json:"phone"`
	IsAdmin   bool                    `json:"is_admin"`
	BirthDate string                  `json:"birth_date"`
	Address   []AddressSchemaResponse `json:"address"`
}

type UserLoginResponseSchema struct {
	Id        uint                    `json:"id"`
	Name      string                  `json:"name"`
	Email     string                  `json:"email"`
	Phone     string                  `json:"phone"`
	IsAdmin   bool                    `json:"is_admin"`
	BirthDate string                  `json:"birth_date"`
	Address   []AddressSchemaResponse `json:"address"`
}

type TokenUsers struct {
	Users UserSchemaResponse `json:"user"`
	Token string             `json:"token"`
}

type TokenUsersLogin struct {
	Users UserLoginResponseSchema `json:"user"`
	Token string                  `json:"token"`
}

type AddressSchemaResponse struct {
	Id           uint      `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Province     string    `json:"province" form:"province"`
	District     string    `json:"district" form:"district"`
	Neighborhood string    `json:"neighborhood" form:"neighborhood"`
	FullAddress  string    `json:"full_address" form:"full_address"`
	Phone        string    `json:"phone" form:"phone"`
	PostalCode   string    `json:"postal_code" form:"postal_code"`
	Country      string    `json:"country" form:"country"`
	City         string    `json:"city" form:"city"`
	Default      bool      `json:"default" form:"default"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

func (u *UserSchema) Validate() []validation.ValidationError {
	validator := validation.NewValidator()
	return validator.ValidateStruct(u)
}

func (u *UserLoginSchema) Validate() []validation.ValidationError {
	validator := validation.NewValidator()
	return validator.ValidateStruct(u)
}
