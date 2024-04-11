package schema

import (
	"github.com/Rhymond/go-money"
	"time"
)

type CategoriesResponseSchema struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Icon        string    `json:"icon" form:"icon"`
	ImagePath   string    `json:"image_path" form:"image_path"`
	Active      bool      `json:"active" form:"active"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type ProductResponseSchema struct {
	ID                 uint                       `json:"id" form:"id"`
	SKU                string                     `json:"sku" form:"sku"`
	ProductName        string                     `json:"product_name" form:"product_name"`
	ImagePath          string                     `json:"image_path" form:"image_path"`
	Size               string                     `json:"size" form:"size"`
	Stock              int                        `json:"stock" form:"stock"`
	RegularPrice       Price                      `json:"regular_price" form:"regular_price"`
	DiscountPrice      Price                      `json:"discount_price" form:"discount_price"`
	ShippingPrice      Price                      `json:"shipping_price" form:"shipping_price"`
	Currency           string                     `json:"currency" form:"currency"`
	Quantity           int                        `json:"quantity" form:"quantity"`
	ProductNote        string                     `json:"product_note" form:"product_note"`
	ShortDescription   string                     `json:"short_description" form:"short_description"`
	ProductDescription string                     `json:"product_description" form:"product_description"`
	CreatedAt          time.Time                  `json:"created_at" form:"created_at"`
	UpdatedAt          time.Time                  `json:"updated_at" form:"updated_at"`
	Categories         []CategoriesResponseSchema `json:"categories" form:"categories"`
	Tags               []TagsResponseSchema       `json:"tags" form:"tags"`
}

type TagsResponseSchema struct {
	ID        uint      `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type Price struct {
	Amount  *money.Money `json:"amount"`
	Display string       `json:"display"`
}
