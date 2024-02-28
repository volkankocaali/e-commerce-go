package repository

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"gorm.io/gorm"
)

type tagDatabase struct {
	DB *gorm.DB
}

func NewTagRepository(DB *gorm.DB) interfaces.TagRepository {
	return &tagDatabase{DB}
}

func (t *tagDatabase) Create(tag models.Tags) (models.Tags, error) {
	if err := t.DB.Create(&tag).Error; err != nil {
		return models.Tags{}, err
	}

	return tag, nil
}

func (tp *tagDatabase) CreateTagProduct(tps models.ProductTags) (models.ProductTags, error) {
	if err := tp.DB.Create(&tps).Error; err != nil {
		return models.ProductTags{}, err
	}

	return tps, nil
}
