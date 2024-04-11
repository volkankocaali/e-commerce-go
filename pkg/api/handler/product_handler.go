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

func (p *ProductHandler) GetProduct(ctx *fiber.Ctx) error {
	// get product id
	productId := ctx.Params("id")

	product, err := p.productUseCase.ListProduct(1, 1, &productId)

	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.NotList, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	if len(product) == 0 {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.NotFound, nil, nil, 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, product_constant.SuccessList, product[0], nil, len(product))
	return ctx.Status(success.StatusCode).JSON(success)
}

func (p *ProductHandler) ListProduct(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page", "1")
	page, err := strconv.Atoi(pageStr)

	perPageStr := ctx.Query("perPage", "10")
	perPage, err := strconv.Atoi(perPageStr)

	// page number control
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.PageNumberNotValid, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// list products user case
	products, err := p.productUseCase.ListProduct(page, perPage, nil)

	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, product_constant.NotList, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, product_constant.SuccessList, products, nil, len(products))
	return ctx.Status(success.StatusCode).JSON(success)
}
