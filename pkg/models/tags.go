package models

import "github.com/volkankocaali/e-commorce-go/pkg/models/traits"

type Tags struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	TagName string `json:"tag_name"`
	Icon    string `json:"icon"`
	traits.Timestampable
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
}

type ProductTags struct {
	TagID     uint     `json:"tag_id"`
	Tags      *Tags    `json:"tags" gorm:"foreignkey:TagID;constraint:OnDelete:CASCADE"`
	ProductID uint     `json:"product_id"`
	Product   *Product `json:"product" gorm:"foreignkey:ProductID;constraint:OnDelete:CASCADE"`
}
