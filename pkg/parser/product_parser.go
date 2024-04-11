package parser

import (
	"github.com/Rhymond/go-money"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
	"sort"
)

type ProductParser struct{}

func NewProductParser() *ProductParser {
	return &ProductParser{}
}

func (p *ProductParser) Parse(pc []models.ProductCategories) ([]schema.ProductResponseSchema, error) {
	tagParser := NewTagParser()
	productMap := make(map[uint]schema.ProductResponseSchema)

	for _, v := range pc {
		existingProduct, found := productMap[v.Product.ID]
		if found {
			existingProduct.Categories = append(existingProduct.Categories, p.createCategoriesResponseSchema(v))
			productMap[v.Product.ID] = existingProduct
		} else {
			productMap[v.Product.ID] = p.createProductResponseSchema(v, tagParser)
		}
	}

	return p.sortProducts(productMap), nil
}

func (p *ProductParser) createProductResponseSchema(v models.ProductCategories, tagParser *TagParser) schema.ProductResponseSchema {
	return schema.ProductResponseSchema{
		ID:                 v.Product.ID,
		SKU:                v.Product.SKU,
		ProductName:        v.Product.ProductName,
		ImagePath:          v.Product.ImagePath,
		Size:               v.Product.Size,
		Stock:              v.Product.Stock,
		Currency:           v.Product.Currency,
		RegularPrice:       p.createPriceSchema(v.Product.RegularPrice, v.Product.Currency),
		DiscountPrice:      p.createPriceSchema(v.Product.DiscountPrice, v.Product.Currency),
		ShippingPrice:      p.createPriceSchema(v.Product.ShippingPrice, v.Product.Currency),
		Quantity:           v.Product.Quantity,
		ProductNote:        v.Product.ProductNote,
		ShortDescription:   v.Product.ShortDescription,
		ProductDescription: v.Product.ProductDescription,
		CreatedAt:          v.Product.CreatedAt,
		UpdatedAt:          v.Product.UpdatedAt,
		Categories:         []schema.CategoriesResponseSchema{p.createCategoriesResponseSchema(v)},
		Tags:               tagParser.Parse(v.Product.ProductTags),
	}
}

func (p *ProductParser) createCategoriesResponseSchema(v models.ProductCategories) schema.CategoriesResponseSchema {
	return schema.CategoriesResponseSchema{
		ID:          v.Categories.ID,
		Name:        v.Categories.Name,
		Description: v.Categories.Description,
		Icon:        v.Categories.Icon,
		ImagePath:   v.Categories.ImagePath,
		Active:      v.Categories.Active,
		CreatedAt:   v.Categories.CreatedAt,
		UpdatedAt:   v.Categories.UpdatedAt,
	}
}

func (p *ProductParser) createPriceSchema(price float64, currency string) schema.Price {
	moneyPrice := money.NewFromFloat(price, currency)
	return schema.Price{
		Amount:  moneyPrice,
		Display: moneyPrice.Display(),
	}
}

func (p *ProductParser) sortProducts(productMap map[uint]schema.ProductResponseSchema) []schema.ProductResponseSchema {
	var mergedProductList []schema.ProductResponseSchema
	for _, product := range productMap {
		mergedProductList = append(mergedProductList, product)
	}

	sort.SliceStable(mergedProductList, func(i, j int) bool {
		if mergedProductList[i].ID != mergedProductList[j].ID {
			return mergedProductList[i].ID < mergedProductList[j].ID
		}
		return mergedProductList[i].ProductName < mergedProductList[j].ProductName
	})

	return mergedProductList
}
