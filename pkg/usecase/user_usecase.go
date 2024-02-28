package usecase

import (
	"errors"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	user_constant "github.com/volkankocaali/e-commorce-go/pkg/constant/user"
	helper_interface "github.com/volkankocaali/e-commorce-go/pkg/helper/interface"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type UserUseCase struct {
	userRepository   interfaces.UserRepository
	walletRepository interfaces.WalletRepository
	config           config.Config
	helper           helper_interface.Helper
}

func NewUserUseCase(
	userRepository interfaces.UserRepository,
	walletRepository interfaces.WalletRepository,
	cfg *config.Config, h helper_interface.Helper) *UserUseCase {
	return &UserUseCase{
		userRepository:   userRepository,
		walletRepository: walletRepository,
		config:           *cfg,
		helper:           h,
	}
}

func (u *UserUseCase) UserSignUp(user schema.UserSchema, ref string) (schema.TokenUsers, error) {
	// first check user already exist
	userCheck := u.userRepository.CheckUserExist(user.Email)

	if userCheck {
		return schema.TokenUsers{}, errors.New(user_constant.AlreadyExist)
	}

	// check user password and confirm password
	if user.Password != user.PasswordConfirmation {
		return schema.TokenUsers{}, errors.New(user_constant.PasswordNotMatch)
	}

	// find reference user
	referenceUser, err := u.userRepository.FindUserByReference(ref)

	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.ReferenceCodeNotFound)
	}

	// hashed password
	hashedPassword, err := u.helper.GeneratePasswordHash(user.Password)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.HashError)
	}

	user.Password = hashedPassword

	// generate referral code
	referralCode, err := u.helper.GenerateReferralCode()
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.ReferralCodeError)
	}

	user.ReferralCode = referralCode

	// create user and save database
	userCreate, err := u.userRepository.SignUp(user)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.SignUpError)
	}

	// create jwt token
	tokenClients, err := u.helper.GenerateTokenClients(userCreate)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.TokenError)
	}

	// ref user add wallet bonus credit
	if err := u.walletRepository.ReferenceUserAddedPointsToWallet(referenceUser); err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.ReferenceUserAddedPointsToWalletError)
	}

	// create new user wallet
	if _, err = u.walletRepository.CreateNewWallet(userCreate.Id); err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.CreateNewWalletError)
	}

	return schema.TokenUsers{
		Users: userCreate,
		Token: tokenClients,
	}, nil
}

func (u *UserUseCase) UserLogin(user schema.UserLoginSchema) (schema.TokenUsersLogin, error) {
	// first check user already exist
	userCheck := u.userRepository.CheckUserExist(user.Email)
	if !userCheck {
		return schema.TokenUsersLogin{}, errors.New(user_constant.NotFound)
	}

	// blocked user check
	_, err := u.userRepository.CheckUserIsBlocked(user.Email)
	if err != nil {
		return schema.TokenUsersLogin{}, errors.New(user_constant.BlockedUser)
	}

	// check user email
	userDetails, err := u.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return schema.TokenUsersLogin{}, errors.New(user_constant.NotFound)
	}

	// check find db password and request password compare
	err = u.helper.CompareHashAndPassword(userDetails.Password, user.Password)
	if err != nil {
		return schema.TokenUsersLogin{}, errors.New(user_constant.PasswordNotCorrect)
	}
	// generate token clients
	tokenString, err := u.helper.GenerateTokenClients(schema.UserSchemaResponse{
		Id:    userDetails.Id,
		Name:  userDetails.Name,
		Email: userDetails.Email,
		Phone: userDetails.Phone,
	})

	return schema.TokenUsersLogin{
		Users: u.SyncUserData(userDetails),
		Token: tokenString,
	}, nil
}

func (u *UserUseCase) SyncUserData(response schema.UserSignInResponse) schema.UserLoginResponseSchema {
	return schema.UserLoginResponseSchema{
		Id:        response.Id,
		Name:      response.Name,
		Email:     response.Email,
		Phone:     response.Phone,
		IsAdmin:   response.IsAdmin,
		BirthDate: response.BirthDate,
		Address:   response.Address,
	}
}
