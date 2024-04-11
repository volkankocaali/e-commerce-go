package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
)

type ProductRepository interface {
	ListProduct(page int, perPage int, productId *string) ([]models.ProductCategories, error)
	Create(product models.Product) (models.Product, error)
	CreateProductCategories(pc models.ProductCategories) (models.ProductCategories, error)
}
