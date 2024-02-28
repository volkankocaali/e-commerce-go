package models

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models/traits"
)

type Preferences struct {
	ID        uint     `json:"id" gorm:"unique;not null"`
	UserID    uint     `json:"user_id" gorm:"not null"`
	ProductID uint     `json:"product_id" gorm:"not null"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID"`
	Users     *Users   `json:"users" gorm:"foreignKey:UserID"`
	Rating    int      `json:"rating"`
	traits.Timestampable
}
