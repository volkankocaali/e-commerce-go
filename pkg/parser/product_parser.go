package parser

import (
	"github.com/Rhymond/go-money"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
	"sort"
)

// ProductParser is the default implementation of ProductParser interface
type ProductParser struct{}

// NewProductParser returns a new instance of ProductParser
func NewProductParser() *ProductParser {
	return &ProductParser{}
}

// Parse implements the Parse method of the ProductParser interface
func (p *ProductParser) Parse(pc []models.ProductCategories) ([]schema.ProductResponseSchema, error) {
	tagParser := NewTagParser()
	productMap := make(map[uint]schema.ProductResponseSchema)

	//  Combine categories by browsing all products
	for _, v := range pc {
		existingProduct, found := productMap[v.Product.ID]
		if found {
			// Product already exists, update by adding the category
			existingProduct.Categories = append(existingProduct.Categories, schema.CategoriesResponseSchema{
				ID:          v.Categories.ID,
				Name:        v.Categories.Name,
				Description: v.Categories.Description,
				Icon:        v.Categories.Icon,
				ImagePath:   v.Categories.ImagePath,
				Active:      v.Categories.Active,
				CreatedAt:   v.Categories.CreatedAt,
				UpdatedAt:   v.Categories.UpdatedAt,
			})
			productMap[v.Product.ID] = existingProduct
		} else {
			regularPrice := money.NewFromFloat(v.Product.RegularPrice, v.Product.Currency)
			discountPrice := money.NewFromFloat(v.Product.DiscountPrice, v.Product.Currency)
			shippingPrice := money.NewFromFloat(v.Product.ShippingPrice, v.Product.Currency)
			tags := tagParser.Parse(v.Product.ProductTags)

			// Product not yet added, create new product
			productMap[v.Product.ID] = schema.ProductResponseSchema{
				ID:          v.Product.ID,
				SKU:         v.Product.SKU,
				ProductName: v.Product.ProductName,
				ImagePath:   v.Product.ImagePath,
				Size:        v.Product.Size,
				Stock:       v.Product.Stock,
				Currency:    v.Product.Currency,
				RegularPrice: schema.Price{
					Amount:  regularPrice,
					Display: regularPrice.Display(),
				},
				DiscountPrice: schema.Price{
					Amount:  discountPrice,
					Display: discountPrice.Display(),
				},
				ShippingPrice: schema.Price{
					Amount:  shippingPrice,
					Display: shippingPrice.Display(),
				},
				Quantity:           v.Product.Quantity,
				ProductNote:        v.Product.ProductNote,
				ShortDescription:   v.Product.ShortDescription,
				ProductDescription: v.Product.ProductDescription,
				CreatedAt:          v.Product.CreatedAt,
				UpdatedAt:          v.Product.UpdatedAt,
				Categories: []schema.CategoriesResponseSchema{
					{
						ID:          v.Categories.ID,
						Name:        v.Categories.Name,
						Description: v.Categories.Description,
						Icon:        v.Categories.Icon,
						ImagePath:   v.Categories.ImagePath,
						Active:      v.Categories.Active,
						CreatedAt:   v.Categories.CreatedAt,
						UpdatedAt:   v.Categories.UpdatedAt,
					},
				},
				Tags: tags,
			}
		}
	}

	// Convert merged products in a map into a slice
	var mergedProductList []schema.ProductResponseSchema
	for _, product := range productMap {
		mergedProductList = append(mergedProductList, product)
	}

	// Sorting by id and product name

	sort.SliceStable(mergedProductList, func(i, j int) bool {
		if mergedProductList[i].ID != mergedProductList[j].ID {
			return mergedProductList[i].ID < mergedProductList[j].ID
		}
		return mergedProductList[i].ProductName < mergedProductList[j].ProductName
	})

	return mergedProductList, nil
}
