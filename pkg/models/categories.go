package models

import (
	"database/sql"
	"github.com/volkankocaali/e-commorce-go/pkg/models/traits"
)

type Categories struct {
	ID           uint          `json:"id" gorm:"unique;not null"`
	CategoriesID sql.NullInt64 `json:"categories_id"`
	Categories   *Categories   `json:"categories" gorm:"foreignkey:CategoriesID;constraint:OnDelete:CASCADE"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Icon         string        `json:"icon"`
	traits.Image
	Active bool `json:"active"`
	traits.Timestampable
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
}
