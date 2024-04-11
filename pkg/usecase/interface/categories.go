package _interface

import "github.com/volkankocaali/e-commorce-go/pkg/models"

type CategoriesUseCase interface {
	ListCategories(page int, perPage int, categoriesId *string) ([]models.Categories, error)
}
