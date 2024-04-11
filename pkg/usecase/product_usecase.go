package usecase

import (
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	helper_interface "github.com/volkankocaali/e-commorce-go/pkg/helper/interface"
	parser_interface "github.com/volkankocaali/e-commorce-go/pkg/parser/interface"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type ProductUseCase struct {
	productRepository    interfaces.ProductRepository
	categoriesRepository interfaces.CategoriesRepository
	config               config.Config
	helper               helper_interface.Helper
	productParser        parser_interface.ProductParser
}

func NewProductUseCase(
	productRepository interfaces.ProductRepository,
	categoriesRepository interfaces.CategoriesRepository,
	cfg *config.Config,
	h helper_interface.Helper,
	productParser parser_interface.ProductParser,
) *ProductUseCase {
	return &ProductUseCase{
		productRepository:    productRepository,
		categoriesRepository: categoriesRepository,
		config:               *cfg,
		helper:               h,
		productParser:        productParser,
	}
}

func (p *ProductUseCase) ListProduct(page int, perPage int, productId *string) ([]schema.ProductResponseSchema, error) {
	// fetch all product id
	products, err := p.productRepository.ListProduct(page, perPage, productId)

	// Parse product response
	parsedResponse, err := p.productParser.Parse(products)

	// Check error
	if err != nil {
		return nil, err
	}

	return parsedResponse, nil
}
