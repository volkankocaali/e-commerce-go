package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type ProductUseCase interface {
	ListProduct(page int, perPage int, userId int) ([]schema.ProductResponseSchema, error)
	GetSimilarCategoryProduct(preferences []models.Preferences) (map[string][]uint, error)
}
