package models

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models/traits"
)

type Product struct {
	ID          uint   `json:"id" gorm:"unique;not null"`
	SKU         string `json:"sku"`
	ProductName string `json:"product_name"`
	traits.Image
	Size               string  `json:"size" gorm:"size:5;default:'M';check:size IN ('S', 'M', 'L', 'XL', 'XXL')"`
	Stock              int     `json:"stock"`
	RegularPrice       float64 `json:"regular_price" gorm:"column:regular_price;default:0.00"`
	DiscountPrice      float64 `json:"discount_price" gorm:"column:discount_price;default:0.00"`
	ShippingPrice      float64 `json:"shipping_price" gorm:"column:shipping_price;default:0.00"`
	Currency           string  `json:"currency" gorm:"column:currency;default:'TRY'"`
	Quantity           int     `json:"quantity"`
	ProductNote        string  `json:"product_note"`
	ShortDescription   string  `json:"short_description"`
	ProductDescription string  `json:"product_description" sql:"type:text;"`
	traits.Timestampable
	Status      bool          `json:"status" gorm:"default:true"`
	Preferences []Preferences `json:"preferences" gorm:"constraint:OnDelete:CASCADE"`
	ProductTags []ProductTags `json:"product_tags" gorm:"constraint:OnDelete:CASCADE"`
	CreatedBy   int           `json:"created_by"`
	UpdatedBy   int           `json:"updated_by"`
}

type Galleries struct {
	ID        uint    `json:"id" gorm:"unique;not null"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"-" gorm:"foreignkey:ProductID;constraint:OnDelete:CASCADE"`
	traits.Image
	Thumbnail    bool `json:"thumbnail"`
	DisplayOrder int  `json:"display_order"`
	traits.Timestampable
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
}
