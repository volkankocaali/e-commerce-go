package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
)

type PreferencesRepository interface {
	FindByUserIdPreferences(userId int) ([]models.Preferences, error)
}
