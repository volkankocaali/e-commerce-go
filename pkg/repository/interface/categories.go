package _interface

import "github.com/volkankocaali/e-commorce-go/pkg/models"

type CategoriesRepository interface {
	Create(categories models.Categories) (models.Categories, error)
	FindByProductIdCategories(productId []uint) ([]models.ProductCategories, error)
	ListCategories(page int, perPage int, categoriesId *string) ([]models.Categories, error)
}
