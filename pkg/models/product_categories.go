package models

type ProductCategories struct {
	CategoriesId uint        `json:"categories_id"`
	Categories   *Categories `json:"categories" gorm:"foreignkey:CategoriesId;constraint:OnDelete:CASCADE"`
	ProductId    uint        `json:"product_id"`
	Product      *Product    `json:"products" gorm:"foreignkey:ProductId;constraint:OnDelete:CASCADE"`
}
