package repository

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"gorm.io/gorm"
)

type preferencesDatabase struct {
	DB *gorm.DB
}

func NewPreferencesRepository(DB *gorm.DB) interfaces.PreferencesRepository {
	return &preferencesDatabase{DB}
}

func (p preferencesDatabase) FindByUserIdPreferences(userId int) ([]models.Preferences, error) {
	var preferences []models.Preferences
	if err := p.DB.Where("user_id = ?", userId).Find(&preferences).Error; err != nil {
		return []models.Preferences{}, err
	}

	return preferences, nil
}
