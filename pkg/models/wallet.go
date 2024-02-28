package models

type Wallet struct {
	ID      uint    `json:"id"  gorm:"unique;not null"`
	UserId  uint    `json:"user_id"`
	User    Users   `json:"-" gorm:"foreignkey:UserId"`
	Balance float64 `json:"balance" gorm:"default:0"`
}
