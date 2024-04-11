package usecase

import (
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	helper_interface "github.com/volkankocaali/e-commorce-go/pkg/helper/interface"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
)

type CategoriesUseCase struct {
	categoriesRepository interfaces.CategoriesRepository
	config               config.Config
	helper               helper_interface.Helper
}

func NewCategoriesUseCase(
	categoriesRepository interfaces.CategoriesRepository,
	cfg *config.Config,
	h helper_interface.Helper,
) *CategoriesUseCase {
	return &CategoriesUseCase{
		categoriesRepository: categoriesRepository,
		config:               *cfg,
		helper:               h,
	}
}

func (c *CategoriesUseCase) ListCategories(page int, perPage int, categoriesId *string) ([]models.Categories, error) {
	categories, err := c.categoriesRepository.ListCategories(page, perPage, categoriesId)

	if err != nil {
		return []models.Categories{}, err
	}

	return categories, nil
}
