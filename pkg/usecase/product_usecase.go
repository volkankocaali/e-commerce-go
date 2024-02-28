package usecase

import (
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	helper_interface "github.com/volkankocaali/e-commorce-go/pkg/helper/interface"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	parser_interface "github.com/volkankocaali/e-commorce-go/pkg/parser/interface"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type ProductUseCase struct {
	productRepository     interfaces.ProductRepository
	categoriesRepository  interfaces.CategoriesRepository
	preferencesRepository interfaces.PreferencesRepository
	config                config.Config
	helper                helper_interface.Helper
	productParser         parser_interface.ProductParser
}

func NewProductUseCase(
	productRepository interfaces.ProductRepository,
	categoriesRepository interfaces.CategoriesRepository,
	preferencesRepository interfaces.PreferencesRepository,
	cfg *config.Config,
	h helper_interface.Helper,
	productParser parser_interface.ProductParser,
) *ProductUseCase {
	return &ProductUseCase{
		productRepository:     productRepository,
		categoriesRepository:  categoriesRepository,
		preferencesRepository: preferencesRepository,
		config:                *cfg,
		helper:                h,
		productParser:         productParser,
	}
}

func (p *ProductUseCase) ListProduct(page int, perPage int, userId int) ([]schema.ProductResponseSchema, error) {
	/* @TODO: User-Based product fulfillment

	When user purchases the product, we will add a record to the preferences table.
	2- When the user adds a product to the wishlist, we will assign it to the preferences table.
	3- When the user looks at any type of product frequently, we will add a record about that product to the preferences table.

	So in this section, in the preferences table, we can actually find products that might be of interest to the user and list them accordingly.
	*/

	// User's preferred products
	preferences, err := p.preferencesRepository.FindByUserIdPreferences(userId)

	if err != nil {
		return nil, err
	}

	// first preferences repository fetch all product id
	products, err := p.productRepository.ListProduct(page, perPage)

	for _, preference := range preferences {
		for _, product := range products {
			if product.ProductId == preference.ProductID {
				product.Product.Preferences = preferences
			}
		}
	}
	// Parse product response
	parsedResponse, err := p.productParser.Parse(products)

	// Check error
	if err != nil {
		return nil, err
	}

	return parsedResponse, nil
}

func (p *ProductUseCase) GetSimilarCategoryProduct(preferences []models.Preferences) (map[string][]uint, error) {
	var productIds []uint
	var categoriesIds []uint
	ids := make(map[string][]uint)

	for _, preference := range preferences {
		productIds = append(productIds, preference.ProductID)
	}

	// get product category
	similarProductCategories, err := p.categoriesRepository.FindByProductIdCategories(productIds)

	if err != nil {
		return nil, err
	}

	for _, pc := range similarProductCategories {
		categoriesIds = append(categoriesIds, pc.CategoriesId)
	}

	ids["categoriesIds"] = categoriesIds
	ids["productIds"] = productIds

	return ids, nil
}
