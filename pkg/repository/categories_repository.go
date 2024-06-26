package repository

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"gorm.io/gorm"
)

type categoriesDatabase struct {
	DB *gorm.DB
}

func NewCategoriesRepository(DB *gorm.DB) interfaces.CategoriesRepository {
	return &categoriesDatabase{DB}
}

func (c *categoriesDatabase) Create(categories models.Categories) (models.Categories, error) {
	if err := c.DB.Create(&categories).Error; err != nil {
		return models.Categories{}, err
	}

	return categories, nil
}

func (c *categoriesDatabase) FindByProductIdCategories(productId []uint) ([]models.ProductCategories, error) {
	var productCategories []models.ProductCategories
	if err := c.DB.Where("product_id IN (?)", productId).Preload("Categories").Find(&productCategories).Error; err != nil {
		return []models.ProductCategories{}, err
	}

	return productCategories, nil
}

func (c *categoriesDatabase) ListCategories(page int, perPage int, categoriesId *string) ([]models.Categories, error) {
	var categories []models.Categories
	q := c.DB

	q = q.Preload("Categories.Categories")

	offset := (page - 1) * perPage

	if categoriesId != nil {
		q = q.Where("categories_id = ?", *categoriesId)
	}

	if err := q.
		Offset(offset).Limit(perPage).
		Find(&categories).Error; err != nil {
		return []models.Categories{}, err
	}

	return categories, nil
}
