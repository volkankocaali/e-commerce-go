package handler

import (
	"github.com/gofiber/fiber/v2"
	request_constant "github.com/volkankocaali/e-commorce-go/pkg/constant/request"
	user_constant "github.com/volkankocaali/e-commorce-go/pkg/constant/user"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
	"github.com/volkankocaali/e-commorce-go/pkg/schema/response"
	services "github.com/volkankocaali/e-commorce-go/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (u *UserHandler) SignUp(ctx *fiber.Ctx) error {
	user := new(schema.UserSchema)

	if err := ctx.BodyParser(user); err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// validate user
	validationError := user.Validate()

	if len(validationError) > 0 {
		res := response.ClientResponse(fiber.StatusBadRequest, user_constant.NotValidate, nil, validationError, 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// check reference code

	ref := ctx.Query("ref")

	userCreate, err := u.userUseCase.UserSignUp(*user, ref)

	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, user_constant.NotSignIn, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusCreated, user_constant.SuccessSignUp, userCreate, nil, 0)
	return ctx.Status(success.StatusCode).JSON(success)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var user schema.UserLoginSchema

	if err := ctx.BodyParser(&user); err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// validate user
	validationError := user.Validate()

	if len(validationError) > 0 {
		res := response.ClientResponse(fiber.StatusBadRequest, user_constant.NotValidate, nil, validationError, 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	login, err := u.userUseCase.UserLogin(user)
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, user_constant.NotFound, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	return ctx.JSON(login)
}
