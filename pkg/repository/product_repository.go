package repository

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

func (p *productDatabase) ListProduct(page int, perPage int) ([]models.ProductCategories, error) {
	var productCategories []models.ProductCategories

	offset := (page - 1) * perPage

	if err := p.DB.Preload("Product").Preload("Categories").Preload("Product.ProductTags.Tags").
		Offset(offset).Limit(perPage).
		Find(&productCategories).Error; err != nil {
		return []models.ProductCategories{}, err
	}

	return productCategories, nil
}

func (p *productDatabase) Create(product models.Product) (models.Product, error) {
	if err := p.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *productDatabase) CreateProductCategories(pc models.ProductCategories) (models.ProductCategories, error) {
	if err := p.DB.Create(&pc).Error; err != nil {
		return models.ProductCategories{}, err
	}

	return pc, nil
}
