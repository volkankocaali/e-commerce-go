package user

const (
	PasswordNotMatch                      = "user password and confirm password not match"
	AlreadyExist                          = "user already exist, please try another email or sign in"
	NotValidate                           = "user not validate"
	NotSignIn                             = "user found, please sign in"
	NotFound                              = "user not found"
	PasswordNotCorrect                    = "password not correct"
	BlockedUser                           = "user blocked"
	ReferenceCodeNotFound                 = "cannot find reference user"
	SuccessSignUp                         = "User successfully created"
	HashError                             = "cannot hash password"
	ReferralCodeError                     = "cannot generate referral code"
	SignUpError                           = "cannot create user"
	TokenError                            = "cannot generate jwt token"
	ReferenceUserAddedPointsToWalletError = "cannot add points to wallet"
	CreateNewWalletError                  = "cannot create new wallet"
)
