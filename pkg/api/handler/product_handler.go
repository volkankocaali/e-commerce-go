package handler

import (
	"github.com/gofiber/fiber/v2"
	product_constant "github.com/volkankocaali/e-commorce-go/pkg/constant/product"
	"github.com/volkankocaali/e-commorce-go/pkg/schema/response"
	services "github.com/volkankocaali/e-commorce-go/pkg/usecase/interface"
	"strconv"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(usecase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: usecase,
	}
}

func (p *ProductHandler) ListProduct(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page", "1")
	page, err := strconv.Atoi(pageStr)

	perPageStr := ctx.Query("perPage", "10")
	perPage, err := strconv.Atoi(perPageStr)

	id := ctx.Locals("id")
	userId, ok := id.(int)

	if !ok {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.UserIdNotValid, nil, nil, 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// page number control
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.PageNumberNotValid, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// list products user case
	products, err := p.productUseCase.ListProduct(page, perPage, userId)

	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.NotList, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, product_constant.SuccessList, products, nil, len(products))
	return ctx.Status(success.StatusCode).JSON(success)
}
