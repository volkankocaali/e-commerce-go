package handler

import (
	"github.com/gofiber/fiber/v2"
	categories_constant "github.com/volkankocaali/e-commorce-go/pkg/constant/categories"
	"github.com/volkankocaali/e-commorce-go/pkg/schema/response"
	services "github.com/volkankocaali/e-commorce-go/pkg/usecase/interface"
	"strconv"
)

type CategoriesHandler struct {
	categoriesUseCase services.CategoriesUseCase
}

func NewCategoriesHandler(usecase services.CategoriesUseCase) *CategoriesHandler {
	return &CategoriesHandler{
		categoriesUseCase: usecase,
	}
}

func (c *CategoriesHandler) ListCategories(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page", "1")
	page, err := strconv.Atoi(pageStr)

	perPageStr := ctx.Query("perPage", "10")
	perPage, err := strconv.Atoi(perPageStr)

	// page number control
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, categories_constant.PageNumberNotValid, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// list categories user case
	categories, err := c.categoriesUseCase.ListCategories(page, perPage, nil)

	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, categories_constant.NotList, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, categories_constant.SuccessList, categories, nil, len(categories))
	return ctx.Status(success.StatusCode).JSON(success)
}
